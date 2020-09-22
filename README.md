# Calculus

[![Build Status](https://github.com/wregis/calculus/workflows/Go/badge.svg)](https://github.com/wregis/calculus/actions)

Calculus is a spreadsheet generator for Golang. It is intended to be used for generating and reading row and column
oriented files.

## Status

This library is under development and is not ready for production yet.

## Installation

To install Calculus, use `go get`:

```
go get github.com/wregis/calculus
```

This will then make the following formats and packages available to you:

Format|Read|Write|Package
-:|:-:|:-:|-
CSV|✔️|✔️|github.com/wregis/calculus/csv
GNumeric|➖ (partial)|➖ (partial)|github.com/wregis/calculus/gnumeric
ODS (Open Document Sheet) |❌|❌|github.com/wregis/calculus/ods (planned)
MS-XLS (Microsoft Excel Binary File) |❌|❌|github.com/wregis/calculus/xls (planned)
XLSX (Office Open XML, Spreadsheet ML)|➖ (partial)|➖ (partial)|github.com/wregis/calculus/xlsx

You can import and/or export data using the format package and can handle data using the main package:

```go
package main

import (
  "bytes"
  "fmt"
  "github.com/wregis/calculus"
  "github.com/wregis/calculus/csv"
)

func main() {
  workbook := calculus.New()

  sheet, _ := workbook.AddSheet("Sheet1")
  sheet.SetValue(0, 0, "Hello")
  sheet.SetValue(0, 1, "World")
  sheet.SetValueByRef("A2", "Foo")
  sheet.SetValueByRef("B2", "Bar")
  sheet.SetValueByRef("C2", "Baz")
  sheet.SetValue(2, 0, 42)

  buf := new(bytes.Buffer)
  err := csv.Write(workbook, buf)
  if err == nil {
    fmt.Println(buf.String())
  }
}
```

## Supported go versions

We support the two latest major Go versions, wich are 1.14 and 1.15 at the moment.

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.

## License

This project is licensed under the terms of the MIT license.
