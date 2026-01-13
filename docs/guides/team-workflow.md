# Team Workflow Guide

GitSynq is not just for solo developers. It can be used by teams working in air-gapped environments.

## The Scenario

A team of 3 developers (Alice, Bob, and Charlie) are working on a secure project in a lab.
- They have a central server in the lab.
- They each have their own laptops with internet access.
- They use GitHub as their primary source of truth.

## The Workflow

### 1. Centralized Server Setup

One person (e.g., Alice) initializes the repository on the lab server:

```bash
gitsync push --full
```

### 2. Team Member Setup

Bob and Charlie also run `gitsync init` on their laptops using the same server details.

### 3. Synchronized Development

When Bob wants to work on Alice's latest changes:
1. Alice pushes her changes from her laptop to the lab server using `gitsync push`.
2. Bob pulls those changes to his laptop using `gitsync pull`.

### 4. Handling Conflicts

Conflicts are handled exactly like they are in a normal Git workflow. If Alice and Bob both change the same file:
1. One of them will push successfully.
2. The second person's `gitsync pull` will result in a merge conflict on their laptop.
3. They resolve the conflict locally on their laptop (using their preferred tools) and then `gitsync push` the resolution back to the server.

### 5. Pushing to GitHub

Since all developers have internet access on their laptops, they can push to GitHub whenever they want.

**Best Practice:** Before pushing to GitHub, always run `gitsync pull` to ensure you have the latest "lab-certified" code.

## Pro-Tip: Branching

Teams should use feature branches.
- Alice works on `feat/ui`.
- Bob works on `feat/api`.
- They both push their branches to the lab server.
- They can merge them on the server or on their laptops.

## Security Note

In a team environment, ensure the SSH user on the server has proper permissions for all team members, or use individual SSH accounts.
