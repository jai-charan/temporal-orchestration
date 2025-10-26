# Temporal Orchestration Example

This project is a simple demonstration of a user onboarding workflow using Temporal.io and Go.

It shows how to use a Temporal Workflow to coordinate multiple steps, like sending an email and subscribing a user, by waiting for signals from external services.

---

## Running the Project

You'll need two terminals open.

### 1. Run the Temporal Worker

First, start the worker. The worker listens to the task queue for work (like running workflows or activities).

```bash
go run main.go
```

### 2. Run the Starter

Next, in a second terminal, run the starter. This will execute the workflow, which then waits for signals. The starter also mocks those signals to complete the process.
Bash

```bash
go run ./starter/starter.go
```