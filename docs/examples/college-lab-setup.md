# Real-World Example: College Lab Setup

## The Scenario

Prince is a computer science student.
- **The Lab:** A room full of workstations in a private LAN.
- **The Server:** A central compute node `lab-server.local` where all code must be run for grading.
- **The Constraint:** The lab network is air-gapped; no external internet.
- **Prince's Gear:** A laptop with his code and full internet access at home.

## The Setup

1. **At home:** Prince initializes his project and pushes it to GitHub.
   ```bash
   mkdir my-lab-project && cd my-lab-project
   git init
   # ... add files ...
   git commit -m "Initial commit"
   git remote add origin https://github.com/princetheprogrammerbtw/my-lab-project.git
   git push -u origin main
   ```

2. **In the lab:** Prince connects his laptop to the lab LAN.
   ```bash
   gitsync init
   # Host: lab-server.local
   # User: s12345
   # Path: ~/assignments
   ```

3. **Initialize the server:**
   ```bash
   gitsync push --full
   ```

4. **Work & Sync:**
   - Prince codes on his laptop (better editor, internet for docs).
   - `gitsync push` to test on the server.
   - If he makes quick fixes directly on the server (e.g., during a demo), he runs `gitsync pull` on his laptop.

5. **Final Submission:**
   - Prince pulls the final version from the server.
   - He goes home and `git push origin main` to GitHub.

## Why this is a "Good Fit"

- **Preserves Grading History:** The professor can see the commit history on the server.
- **Zero Server Setup:** Prince doesn't need to install anything on the lab server.
- **Safe:** No risk of losing work if a lab machine is wiped.

## Why this is a "Bad Fit"

- If the server has internet, just use `git push/pull` directly.
- If the project has multi-gigabyte binary assets, bundles might become slow.
