# cutter

A command-line tool for extracting fields from delimited text input,
inspired by Unix `cut`.

## Usage

```text
cutter -f <fields> [-d <delimiter>]
```

`cutter` reads lines from standard input and prints the selected fields
for each line.

## Flags

| Flag | Description                    | Default |
|------|--------------------------------|---------|
| `-f` | Fields to extract (required)   | —       |
| `-d` | Input/output delimiter         | `\t`    |

## Field Syntax

Fields are 1-indexed, matching the behavior of Unix `cut`.

| Syntax      | Meaning                    |
|-------------|----------------------------|
| `1`         | Field 1                    |
| `1,3`       | Fields 1 and 3             |
| `1-3`       | Fields 1 through 3         |
| `1,3-5,7`   | Fields 1, 3, 4, 5, and 7  |

## Examples

```sh
go run cmd/cutter/main.go -d ',' -f 3,1-3 ./tmp/data.txt
```
## License

MIT
