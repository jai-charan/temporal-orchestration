# Temporal Orchestration

A Go-based Temporal workflow application that demonstrates durable event-driven orchestration for user onboarding processes.

## Overview

This project implements a user onboarding workflow using Temporal that:
- Sends a welcome email to new users
- Waits for external service confirmations via signals
- Sends a completion email when onboarding is finished

## Architecture

```
├── types/           # Shared data types
├── activity/        # Temporal activities (business logic)
├── workflow/        # Temporal workflows (orchestration)
├── worker/          # Worker process to execute workflows/activities
└── client/          # Client to start workflows and send signals
```

## Prerequisites

- Go 1.24.2 or higher
- Temporal server running locally (or cloud instance)

## Quick Start

1. **Start Temporal server** (if running locally):
   ```bash
   temporal server start-dev
   ```

2. **Start the worker** (in one terminal):
   ```bash
   make worker
   ```

3. **Run the client** (in another terminal):
   ```bash
   make client
   ```

## Project Structure

### Types (`types/`)
- `user.go`: Shared data structures used across the application

### Activities (`activity/`)
- `email.go`: Email-related activities
  - `SendWelcomeEmail`: Sends initial welcome email
  - `SendOnboardingCompleteEmail`: Sends completion email

### Workflows (`workflow/`)
- `workflow.go`: Main user onboarding workflow
  - `UserOnboardingWorkflow`: Orchestrates the entire onboarding process

### Worker (`worker/`)
- `worker.go`: Temporal worker that processes workflows and activities

### Client (`client/`)
- `main.go`: Client application that starts workflows and sends signals

## Workflow Flow

1. **Workflow starts** with user data
2. **SendWelcomeEmail activity** executes
3. **Wait for external signals**:
   - `email_sent`: Confirmation from email service
   - `subscribed`: Confirmation from subscription service
4. **SendOnboardingCompleteEmail activity** executes
5. **Workflow completes**

## Available Commands

```bash
# Run both worker and client
make run

# Run in development mode
make dev

# Build all components
make build

# Clean build artifacts
make clean
```

## Configuration

The application uses default Temporal configuration. To customize:

- **Temporal Server**: Set `TEMPORAL_HOST` environment variable
- **Task Queue**: Modify `user-task-queue` in worker and client
- **Timeouts**: Adjust activity timeouts in workflow code

## Development

### Adding New Activities

1. Create activity in `activity/` directory
2. Register in `worker/worker.go`
3. Call from workflow in `workflow/workflow.go`

### Adding New Workflows

1. Create workflow in `workflow/` directory
2. Register in `worker/worker.go`
3. Execute from client in `client/main.go`

## Troubleshooting

### Common Issues

1. **Worker not starting**: Ensure Temporal server is running
2. **Signals not received**: Check signal names match between client and workflow
3. **Activities failing**: Verify activity registration in worker

### Logs

- Worker logs show activity and workflow execution
- Client logs show workflow start and signal sending
- Temporal UI shows workflow history and state

## License

MIT License
