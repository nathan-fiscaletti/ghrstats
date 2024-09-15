# ghrstats

ghrstats is a simple command line tool to retrieve download statistics from GitHub releases.

## Installation

```bash
$ go install github.com/rodrigobrito/ghrstats/cmd/ghrstats@latest
```

## Usage

```bash
$ ghrstats [-patterns *.ext1,*.ext2,...] <repo>
```

## Example

```bash
$ ghrstats nathan-fiscaletti/framecast
2024/09/15 00:41:10 Total downloads: 544
```

## License

MIT