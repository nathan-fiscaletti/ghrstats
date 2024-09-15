# ghrstats

ghrstats is a simple command line tool to retrieve download statistics from GitHub releases based on release assets.

It will aggregate the release counts based on asset name, and then filter
the results based on the provided patterns (if any) before displaying the total aggregate count.

## Installation

```bash
$ go install github.com/nathan-fiscaletti/ghrstats/cmd/ghrstats@latest
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
$ ghrstats -patterns "*.exe,*.dmg,*.deb" nathan-fiscaletti/framecast
```

## License

MIT