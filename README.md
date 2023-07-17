# Archive Differ [![Go](https://github.com/4cecoder/archive-differ/actions/workflows/go.yml/badge.svg)](https://github.com/4cecoder/archive-differ/actions/workflows/go.yml)

Archive Differ is a CLI application developed in Go that helps in handling operations with tar archives. This tool allows you to open tar archives, search for `mets.xml` files in the archive, and compare the differences in `mets.xml` files between different archives.

## Getting Started

### Prerequisites

- Go programming language

### Installation

Clone this repository into your local machine:

```bash
git clone https://github.com/Diogenesoftoronto/archive-differ.git
```

Navigate into the directory:
`cd archive-differ`

Build the application:
`go build`


## Usage

The CLI application can be used with the following commands:

### Open an Archive

This command opens a tar archive:

```bash
./myapp open [path]
```

Replace `[path]` with the path of your tar archive.

### Search for mets.xml in an Archive

This command searches for a mets.xml file in a tar archive:

`./myapp search [archive path]`

Replace `[archive path]` with the path of your tar archive.


# Compare two mets.xml Files

This command compares two mets.xml files:
```bash
./myapp compare [path1] [path2]
```


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## License

This project is licensed under the terms of the GNU General Public License (GPL).
