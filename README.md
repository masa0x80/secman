# secman

[![MIT LICENSE](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

## Description

`secman` is a secrets file manager.
You can manage secret files such as included `.gitignore` by using this tool.

## Usage

### Save files

```bash
$ secman save [<fileName>...]
```

### Restore files

```bash
$ secman restore [<fileName>...]
```

### List files

```bash
$ secman list [<dirName>]
```

### Sync files

```bash
$ secman sync [<dirName>]
```

### Help

```bash
$ secman help
```

## Installation

```bash
$ go get github.com/masa0x80/secman/...
```

## Author

[Kimura Masayuki](https://github.com/masa0x80)
