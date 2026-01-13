# Real-World Example: Research HPC Cluster

## The Scenario

Dr. Aris is a researcher working on high-performance computing (HPC) simulations.
- **The Cluster:** A massive HPC cluster with thousands of nodes.
- **The Access:** Only accessible via a secure login node.
- **The Constraint:** Compute nodes and the login node have NO direct internet access for security and to prevent data exfiltration.
- **The Workflow:** Aris develops complex simulation code on her workstation and needs to run it on the cluster.

## The Setup

1. **On Workstation:** Aris has her research code on a private GitLab instance.
2. **Configuration:**
   ```bash
   gitsync init
   # Host: hpc-login.university.edu
   # Remote Path: /home/aris/simulations
   ```

## The Workflow

1. **Pushing Code:**
   Whenever Aris updates her simulation logic, she runs:
   ```bash
   gitsync push
   ```
   GitSynq packages the new physics modules and sends them to her home directory on the cluster.

2. **Running Jobs:**
   Aris SSHs into the cluster and submits her jobs using the updated code.

3. **Retrieving Results:**
   Sometimes her simulations generate updated configuration files or small data summaries that she wants to keep in version control. She commits these on the cluster.
   ```bash
   # On laptop
   gitsync pull
   ```

## Why this is a "Good Fit"

- **Security:** No need to open firewall ports for Git/HTTPS.
- **Reproducibility:** Every simulation run on the cluster is tied to a specific Git commit hash.
- **Speed:** Only the changed source code is transferred, not large datasets.

## Why this is a "Bad Fit"

- **Large Datasets:** Aris should NOT use GitSynq for the terabytes of simulation *output*. She uses specialized tools like `rsync` or Globus for that.
- **Binary Blobs:** If the simulation requires large compiled binaries, they should be managed via a package manager or shared filesystem, not Git bundles.

## Real-World Pro-Tip

Aris uses `gitsync push` in her local "before-job" script to ensure the cluster always has the exact code she is about to run.
