# How It Works

GitSynq bridges the gap between internet-connected machines and air-gapped servers by treating the air-gapped server as a remote that is synchronized via file transfers.

## The Core Technology: Git Bundles

Git has a built-in feature called `git bundle`. It allows you to package objects and references into a single binary file. This file can be treated exactly like a remote repository.

### The Push Process

1. **Calculate Changes:** GitSynq identifies which commits are on your local branch but not on the remote tracking branch (e.g., `origin/main`).
2. **Create Bundle:** It runs `git bundle create` to package these specific commits into a `.bundle` file.
3. **Transfer:** The bundle is uploaded to the remote server via SFTP (SSH).
4. **Remote Update:** GitSynq executes a series of SSH commands on the server to:
   - Initialize a new repo from the bundle (if it doesn't exist).
   - Or, fetch from the bundle and merge it into the existing repo.

### The Pull Process

1. **Remote Bundle:** GitSynq connects to the server and runs `git bundle create` on the remote repository to package all its references.
2. **Download:** The resulting bundle is downloaded to your local machine.
3. **Local Merge:** GitSynq adds the local bundle file as a temporary remote and fetches/merges the changes into your local branch.

## Why this is better than SCPing files?

1. **Integrity:** Git ensures that the history is preserved exactly. If a transfer is corrupted, Git will detect it.
2. **Efficiency:** Incremental pushes only transfer the *new* commits, not the entire project.
3. **Conflict Resolution:** Since it uses standard Git merges, you can resolve conflicts using your favorite tools if changes happened on both sides.
4. **Visibility:** You can see exactly what changed using `git log` and `git diff` after a sync.
