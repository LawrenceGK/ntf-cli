# ntf-cli

[中文](README_zh.md) | English

A simple command-line tool that wraps the ntfy client CLI tool to simplify the message sending process. This tool supports both Windows and Linux platforms.

## Features

- Configure default topic and message
- Send messages via simple command-line arguments
- Rich message options support:
  - Delayed sending (--in)
  - Message title (--title)
  - Message priority (--priority)
  - Message tags (--tags)
- Custom topic and message content support

## Installation

1. Ensure Go environment and ntfy client are installed
2. Clone this project:
   ```bash
   git clone https://github.com/yourusername/ntf-cli.git
   ```
3. Enter project directory and install dependencies:
   ```bash
   cd ntf-cli
   go mod tidy
   ```

## Usage Examples

Basic usage:

- Send default message to default topic:

  ```
  ntf
  ```

- Send message with time parameter:

  ```
  ntf --in 30min
  ```

- Send custom message to default topic:

  ```
  ntf "custom message"
  ```

## Configuration

Configuration file will be automatically created at the following locations:

- Windows: `%LOCALAPPDATA%\ntf-cli\config.json`
- Linux/macOS: `~/.config/ntf-cli/config.json`

Configuration file format (JSON):
```
{
  "default_topic": "your_default_topic",
  "default_message": "your_default_message"
}
```

Default topic and message can be set in the `internal/config/config.go` file. Please modify this file as needed.

## Contributing

Issues and contributions are welcome! Please follow the contribution guidelines.

## License

This project is licensed under the MIT License. See the LICENSE file for details.