# Telegram Bot Project Structure

This project is structured for a Go-based Telegram bot integrated with Rust smart contracts. Below is an overview of the directory structure and its contents.

## Directory Structure

### `cmd/`

This directory contains the entry point for the bot application.

- **`bot/`**  
  Entry point for the bot.
  - `main.go`: Main application logic.

### `pkg/`

Contains the core logic for the bot, contract interactions, and configuration.

- **`bot/`**  
  Bot-specific code, handlers, and interaction logic.
  - `handler.go`: Handlers for different bot commands and interactions.
  - `inline_buttons.go`: Inline button generation and callback handling.
  - `updater.go`: Logic for updating and processing incoming messages.
  - `utils.go`: Utility functions for bot operations.

- **`contract/`**  
  Logic for interacting with Rust smart contracts.
  - `contract.go`: Functions for calling Rust smart contract APIs.
  - `types.go`: Data structures related to contract interactions.

- **`config/`**  
  Handles configuration and environment settings.
  - `config.go`: Configuration loading and environment variables.
  - `constants.go`: Constant values used throughout the project.

### `scripts/`

Scripts for database setup, data generation, and other auxiliary tasks.

### `tests/`

Contains unit and integration tests for different project components.

- `bot_test.go`: Tests for the bot logic.
- `contract_test.go`: Tests for the contract interaction logic.

### `go.mod`

Go module definition file.

### `go.sum`

Go module checksum file.

### `README.md`

Project documentation.

---

This structure ensures a clean, modular approach to building and maintaining the Telegram bot with integrations for smart contract interactions and configurations.
