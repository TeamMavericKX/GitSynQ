# Command Reference

## `gitsync init`

Initializes GitSynq for the current repository.

- **Purpose:** Creates the `.gitsync.yaml` configuration file.
- **Interactive:** Prompts for server details, project name, etc.

## `gitsync push`

Synchronizes local changes to the remote server.

- **Options:**
  - `-f, --full`: Force a full repository push (useful for first-time setup).
  - `-a, --all`: Include all branches in the bundle.
- **Behavior:** Creates an incremental bundle by default.

## `gitsync pull`

Fetches and merges changes from the remote server.

- **Options:**
  - `-p, --push`: Automatically push to the origin remote (e.g., GitHub) after a successful pull and merge.
- **Behavior:** Creates a bundle on the server, downloads it, and merges it locally.

## `gitsync status`

Displays the current synchronization status.

- **Checks:**
  - Local branch and last commit.
  - Remote branch and last commit.
  - Presence of uncommitted changes on both sides.
  - Connection to the remote server.

## `gitsync config`

Manages the GitSynq configuration.

- **Options:**
  - `-s, --show`: Displays the current configuration (default).
  - `-e, --edit`: Re-runs the interactive initialization to edit settings.

## Global Flags

- `-c, --config string`: Path to a specific config file (default: `.gitsync.yaml`).
- `-v, --verbose`: Enable verbose output for debugging.
- `-h, --help`: Display help for a command.
- `--version`: Display the version of GitSynq.
