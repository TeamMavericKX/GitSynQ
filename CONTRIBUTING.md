# Contributing to GitSynq

First off, thanks for taking the time to contribute! 🎉

The following is a set of guidelines for contributing to GitSynq. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## How Can I Contribute?

### Reporting Bugs

- Use a clear and descriptive title for the issue to identify the problem.
- Describe the exact steps which reproduce the problem in as many details as possible.
- Include the output of `gitsync status --verbose`.

### Suggesting Enhancements

- Explain why this enhancement would be useful to most GitSynq users.
- Provide a step-by-step description of the suggested enhancement in as many details as possible.

### Pull Requests

- Follow the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).
- Ensure that the build and tests pass.
- Update the documentation if you're adding a new feature.
- Include a clear description of the change in your PR.

## Development Setup

1. Clone the repository.
2. Install dependencies: `make deps`.
3. Run tests: `make test`.
4. Build the binary: `make build`.

## Commit Message Format

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification:

- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `chore:` for maintenance tasks
- `refactor:` for code refactoring

Example: `feat: add support for encrypted bundles`

---

Built with ❤️ by PrinceTheProgrammer
