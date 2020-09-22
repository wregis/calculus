package xlsx

import (
	"archive/zip"
	"compress/flate"
	"encoding/xml"
	"io"

	"github.com/wregis/calculus"
)

// Write takes a workbook object and writes it to a XLSX file.
func Write(workbook calculus.Workbook, out io.Writer) error {
	sheets := workbook.Sheets()
	if sheets == nil || len(sheets) == 0 {
		return calculus.NewError(nil, "No sheet to write XLSX")
	}
	writer := zip.NewWriter(out)
	writer.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(w, flate.BestSpeed)
	})
	if err := writeContentTypes(writer, workbook); err != nil {
		return err
	}
	if err := writePropertiesApp(writer, workbook); err != nil {
		return err
	}
	if err := writeRelationships(writer, workbook); err != nil {
		return err
	}
	if err := writeStyles(writer, workbook); err != nil {
		return err
	}
	if err := writeWorkbook(writer, workbook); err != nil {
		return err
	}
	if err := writeWorksheets(writer, workbook); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return calculus.NewError(err, "Unable to close ZIP archive")
	}
	return nil
}

const xmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n"

func writeXMLToFile(writer *zip.Writer, name string, data interface{}) error {
	file, err := writer.Create(name)
	if err != nil {
		return calculus.NewError(err, "Could not create file")
	}
	if _, err := file.Write([]byte(xmlHeader)); err != nil {
		return calculus.NewError(err, "Could not write XML header")
	}
	var buf []byte
	if buf, err = xml.Marshal(data); err != nil {
		return calculus.NewError(err, "Failed to serialize XML")
	}
	if _, err := file.Write(buf); err != nil {
		return calculus.NewError(err, "Could not write file data")
	}
	return nil
}
