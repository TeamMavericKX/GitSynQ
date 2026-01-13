# Code Style Guidelines

We follow the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) with a few project-specific additions.

## 1. Error Handling

- **Wrap Errors:** Always wrap errors with context using `%w`.
  ```go
  if err != nil {
      return fmt.Errorf("failed to open bundle: %w", err)
  }
  ```
- **Specific Errors:** Create custom error types or variables for common error conditions.
- **Avoid Panics:** Never use `panic()` in the main code. Return errors instead.

## 2. Naming Conventions

- **Variable Names:** Use short, descriptive names. Avoid `temp` or `data` if something more specific like `bundlePath` is available.
- **Interfaces:** Name interfaces after what they do (e.g., `Syncer`, `Transporter`).
- **Exporting:** Only export what is absolutely necessary for other packages.

## 3. Function Design

- **Small Functions:** Keep functions small and focused on a single task.
- **Arguments:** Avoid passing too many arguments to a function. Use a struct if necessary.
- **Return Values:** Prefer returning a value and an error.

## 4. Documentation

- **GoDoc:** All exported functions, structs, and constants MUST have GoDoc comments.
- **Internal Comments:** Use comments to explain the *why* of complex logic, not the *what*.

## 5. Tooling

- Use `go fmt` for formatting.
- Use `go vet` to catch common mistakes.
- Use `staticcheck` if possible.

## Example

```go
// CreateBundle creates a Git bundle for the given branch.
// It returns the path to the created bundle or an error if it fails.
func CreateBundle(branch string) (string, error) {
    // Logic here...
}
```
