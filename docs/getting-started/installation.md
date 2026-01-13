# Installation Guide

GitSynq can be installed in several ways.

## 1. Using Go (Recommended)

If you have Go installed on your machine, you can install the latest version directly:

```bash
go install github.com/princetheprogrammerbtw/gitsynq@latest
```

This will install the `gitsync` binary to your `$GOPATH/bin` directory. Make sure this directory is in your system's `PATH`.

## 2. Pre-built Binaries

You can download the pre-built binaries for your operating system from the [Releases](https://github.com/princetheprogrammerbtw/gitsynq/releases) page on GitHub.

1. Download the archive for your platform.
2. Extract the archive.
3. Move the `gitsync` binary to a directory in your `PATH` (e.g., `/usr/local/bin` or `~/bin`).

## 3. Building from Source

If you want to build the project yourself:

```bash
git clone https://github.com/princetheprogrammerbtw/gitsynq
cd gitsynq
make build
```

The binary will be located in the `build/` directory. You can also run `make install` to copy it to your `~/bin` directory.

## Requirements

- **Local:** Git 2.0+
- **Remote:** Git 2.0+, SSH access
- **OS:** Linux, macOS, or Windows
