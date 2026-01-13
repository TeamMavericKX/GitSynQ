# Environment Variables

GitSynq supports several environment variables that can be used to override configuration settings or control behavior.

## Configuration Overrides

You can override any value in `.gitsync.yaml` using the prefix `GITSYNC_` followed by the path to the setting in uppercase, with underscores separating the components.

- `GITSYNC_SERVER_HOST`: Overrides `server.host`.
- `GITSYNC_SERVER_USER`: Overrides `server.user`.
- `GITSYNC_PROJECT_BRANCH`: Overrides `project.branch`.

### Example

```bash
GITSYNC_SERVER_HOST=10.0.0.5 gitsync status
```

## SSH Authentication

- `SSH_AUTH_SOCK`: Path to the SSH agent socket. If set, GitSynq will attempt to use the agent for authentication.

## Debugging

- `GITSYNC_VERBOSE`: If set to `true`, enables verbose logging (equivalent to `-v` or `--verbose`).

## How it works

GitSynq uses the [Viper](https://github.com/spf13/cobra) library for configuration management. Viper automatically maps environment variables to configuration keys if `viper.AutomaticEnv()` is called.
