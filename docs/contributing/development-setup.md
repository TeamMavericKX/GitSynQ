# Development Setup

This guide will help you set up your local environment for contributing to GitSynq.

## Prerequisites

- **Go:** 1.21 or later.
- **Git:** 2.0 or later.
- **Make:** Standard make utility.
- **SSH Server:** A local or remote SSH server to test synchronization.

## 1. Fork and Clone

```bash
git clone https://github.com/princetheprogrammerbtw/gitsynq
cd gitsynq
```

## 2. Install Dependencies

```bash
make deps
```

## 3. Running the Project

You can run the project directly from source:

```bash
go run main.go status
```

Or build and run:

```bash
make build
./build/gitsync status
```

## 4. Running Tests

We use standard Go testing.

```bash
make test
```

## 5. Coding Standards

- Follow the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).
- Run `go fmt ./...` before committing.
- Ensure all exported functions have GoDoc comments.

## 6. Testing with a Local Server

For integration testing, you can use a local SSH server (e.g., `sshd` running on localhost).

1. Ensure your public key is in `~/.ssh/authorized_keys`.
2. Run `gitsync init` and use `localhost` as the host.
3. Verify that `push` and `pull` work as expected.
