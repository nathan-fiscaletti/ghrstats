# ghrstats

ghrstats is a simple command line tool to retrieve download statistics from GitHub releases based on release assets.

## Installation

```bash
$ go install github.com/rodrigobrito/ghrstats/cmd/ghrstats@latest
```

## Usage

```bash
$ ghrstats [-patterns *.ext1,*.ext2,...] <repo>
```

## Examples

```bash
$ ghrstats nathan-fiscaletti/framecast
```

```bash
$ ghrstats -patterns "*.exe,*.dmg" nathan-fiscaletti/framecast
```

## License

MIT