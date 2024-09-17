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
$ ghrstats -r <repo> [-a <total|itemized>] [-f *.ext1,*.ext2,...]
```

## Examples

- Count all downloads for all release assets and return the value
    ```bash
    $ ghrstats -r nathan-fiscaletti/framecast
    ```
    ```json
    {"aggregate_downloads":544}
    ```

- Count all downloads for all release assets filtering based by the provided file patterns
    ```bash
    $ ghrstats -r nathan-fiscaletti/framecast -f "*.exe,*.dmg,*.deb" 
    ```
    ```json
    {"aggregate_downloads":490}
    ```

- Itemize aggregate downloads for each release asset matching the provided filter
    ```bash
    $ ghrstats -r nathan-fiscaletti/framecast -f "*.exe" -a itemized
    ```
    ```json
    {"Advanced.Screen.Streamer.Setup.1.0.0.exe":5,"Advanced.Screen.Streamer.Setup.1.0.1.exe":6,"Advanced.Screen.Streamer.Setup.win32.1.0.2.exe":9,"Advanced.Screen.Streamer.Setup.win32.1.0.4.exe":5,"Advanced.Screen.Streamer.Setup.win32.1.0.5.exe":6,"Advanced.Screen.Streamer.Setup.win32.1.0.6.exe":7,"Advanced.Screen.Streamer.Setup.win32.1.0.7.exe":13,"FrameCast.Setup.win32.exe":73,"FrameCast.Setup.windows_amd64.exe":65}
    ```

## License

MIT
