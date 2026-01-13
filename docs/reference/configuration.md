# Configuration Reference

GitSynq uses a YAML file named `.gitsync.yaml` located at the root of your project.

## Schema

### `project`

- `name` (string): The name of your project. This is used for the directory name on the server.
- `branch` (string): The primary branch to synchronize (e.g., `main` or `master`).

### `server`

- `host` (string): The hostname or IP address of the remote server.
- `user` (string): The SSH username.
- `port` (int): The SSH port (default: `22`).
- `remote_path` (string): The base directory on the server where projects are stored (e.g., `~/projects`).
- `ssh_key_path` (string, optional): Path to a specific SSH private key. If omitted, GitSynq will try default locations (`~/.ssh/id_rsa`, etc.).

### `bundle`

- `directory` (string): The local directory where temporary bundles are stored (default: `.gitsync-bundles`).
- `compress` (bool): Whether to compress bundles (currently reserved for future use).
- `max_history` (int): Number of old bundles to keep locally (default: `10`).

## Example File

```yaml
project:
  name: gitsynq
  branch: main
server:
  host: 192.168.12.4
  user: prince
  port: 22
  remote_path: ~/lab-work
  ssh_key_path: ~/.ssh/id_ed25519
bundle:
  directory: .gitsync-bundles
  compress: true
  max_history: 10
```
