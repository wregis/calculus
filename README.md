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

This will then make the following packages available to you:

* github.com/wregis/calculus
* github.com/wregis/calculus/csv
* github.com/wregis/calculus/ods (in development)
* github.com/wregis/calculus/xls (in development)
* github.com/wregis/calculus/xlsx (in development)

Format|Read|Write
-|-|-
CSV|✔️|✔️
GNumeric|❌|❌
ODS|❌|❌
XLS|❌|❌
XLSX|❌|❌

You can import and/or export data using the format package and can handle data using the main package:

```go
package foo

import (
  "github.com/wregis/calculus"
  "github.com/wregis/calculus/csv"
)

func DoSomeCsv() {
  workbook := calculus.New()
  sheet := workbook.ActiveSheet()
  sheet.SetValue(0, 0, "Hello")
  sheet.SetValue(0, 1, "World")
  sheet.SetValueByRef("A2", "Foo")
  sheet.SetValueByRef("B2", "Bar")
  sheet.SetValueByRef("C2", "Baz")

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
