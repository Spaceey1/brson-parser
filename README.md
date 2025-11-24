# brson-parser
A command line tool to parse and convert brson files produced by [Resonite](https://resonite.com/)

## Usage
- CLI

```
brsoncli [-e input.json | -d input.brson] -o output
  -e <file>    Encode JSON → BRSon.
  -d <file>    Decode BRSon → JSON.
  -o <file>    Output file path.
Exactly one of -e or -d must be specified.
```


- Go package

You can use the package in `./bronsparser` to use in your own go projects.

#### `ReadBrson(data []byte)` → `map[string]any, error`
Takes BRSON-encoded bytes and returns the decoded document.

#### `WriteBrson(doc map[string]any)` → `[]byte, error`
Takes a document and returns its BRSON-encoded byte representation.

#### `WriteBrsonToFile(doc map[string]any, path string)` → `error`
Takes a document and writes it as a BRSON file to the given path.

#### `ReadBrsonFromFile(path string)` → `map[string]any, error`
Takes a file path to a BRSON file and returns the decoded document.

## Building
Run `make build`
