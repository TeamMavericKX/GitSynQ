# Quick Start Guide

This guide will get you up and running with GitSynq in less than 2 minutes.

## 1. Installation

Assuming you have Go installed:

```bash
go install github.com/princetheprogrammerbtw/gitsynq@latest
```

Verify the installation:

```bash
gitsync --version
```

## 2. Initialize your project

Navigate to your local Git repository and run:

```bash
gitsync init
```

Follow the prompts to configure your remote server details.

## 3. Perform your first push

```bash
gitsync push --full
```

The `--full` flag is required for the first push to initialize the repository on the remote server.

## 4. Pull changes back

After working on the server and committing your changes there:

```bash
gitsync pull
```

Or, to automatically push to GitHub after pulling:

```bash
gitsync pull --push
```
