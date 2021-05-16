Cheesy Arena Lite [![Build Status](https://github.com/Team254/cheesy-arena-lite/actions/workflows/test.yml/badge.svg)](https://github.com/Team254/cheesy-arena-lite/actions)
============
A game-agnostic field management system that just works.

For the game-specific version, see [Cheesy Arena](https://github.com/Team254/cheesy-arena).

## License
Teams may use Cheesy Arena Lite freely for practice, scrimmages, and off-season events. See [LICENSE](LICENSE) for more details.

## Installing
**From a pre-built release**

Download the [latest release](https://github.com/Team254/cheesy-arena-lite/releases). Pre-built packages are available for Linux, macOS (x64 and M1), and Windows.

On recent versions of macOS, you may be prevented from running an app from an unidentified developer; see [these instructions](https://support.apple.com/guide/mac-help/open-a-mac-app-from-an-unidentified-developer-mh40616/mac) on how to bypass the warning.

**From source**

1. Download [Go](https://golang.org/dl/) (version 1.16 or later recommended)
1. Clone this GitHub repository to a location of your choice
1. Navigate to the repository's directory in the terminal
1. Compile the code with `go build`
1. Run the `cheesy-arena-lite` or `cheesy-arena-lite.exe` binary
1. Navigate to http://localhost:8080 in your browser (Google Chrome recommended)

**IP address configuration**

When running Cheesy Arena Lite on a playing field with robots, set the IP address of the computer running Cheesy Arena Lite to 10.0.100.5. By a convention baked into the FRC Driver Station software, driver stations will broadcast their presence on the network to this hardcoded address so that the FMS does not need to discover them by some other method.

When running Cheesy Arena Lite without robots for testing or development, any IP address can be used.

## Further reading
Please see the game-specific [Cheesy Arena](https://github.com/Team254/cheesy-arena) README for technical details and acknowledgements.
