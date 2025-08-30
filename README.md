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
```

## Configuration

The application uses default Temporal configuration. To customize:

- **Temporal Server**: Set `TEMPORAL_HOST` environment variable
- **Task Queue**: Modify `user-task-queue` in worker and client
- **Timeouts**: Adjust activity timeouts in workflow code


## Troubleshooting

### Common Issues

1. **Worker not starting**: Ensure Temporal server is running
2. **Signals not received**: Check signal names match between client and workflow
3. **Activities failing**: Verify activity registration in worker



## License

MIT License
