## go-cli-ui

Originally this library was written for [BOSH CLI v2](http://bosh.io/docs/cli-v2.html). Given its generic nature, it's been extracted as a separate, easily importable library.

- `errors` package: helps a bit with formatting of errors
- `ui` package: helps with printing CLI content
  - `table` package: helps format CLI content as a table

Examples:

- [Table](examples/table/main.go)
