# ntf-cli

[中文](README_zh.md) | English

A simple command-line tool that works with the ntfy client CLI tool to simplify the message sending process.

## Prerequisites
- [ntfy client](https://docs.ntfy.sh/install/) must be installed and accessible from command line

## Installation

### Windows
1. Download `ntf.exe` from [Releases](https://github.com/LawrenceGK/ntf-cli/releases)
2. Add the binary location to your system PATH

### Linux
Option 1: Download from releases (Recommended)
1. Download the latest `ntf` binary from [Releases](https://github.com/LawrenceGK/ntf-cli/releases)
2. Move the binary to `/usr/local/bin` and rename it to `ntf`:
   ```bash
   sudo mv ntf-cli /usr/local/bin/ntf
   ```
3. Make sure the binary is executable:
   ```bash
   sudo chmod +x /usr/local/bin/ntf
   ```

Option 2: Build from source
1. Ensure Go environment and ntfy client are installed
2. Clone this project:
   ```bash
   git clone https://github.com/lawrenceGK/ntf-cli.git
   ```
3. Enter project directory and install dependencies:
   ```bash
   cd ntf-cli
   go mod tidy
   ```
4. Build the binary:
   ```bash
   go build -o ntf cmd/ntf/main.go
   ```
5. Move to a directory in your PATH (optional):
   ```bash
   sudo mv ntf /usr/local/bin/
   ```

## Usage Examples

Basic usage:

- Send default message to default topic:

  ```
  ntf
  ```

- Send message with time parameter (supports both --in/-i and --at/-a):

  ```
  ntf -i 30min
  # or
  ntf -a 30min
  ```

- Send custom message to default topic:

  ```
  ntf "custom message"
  ```

- Send message with title:
  ```
  ntf --title "Title" "message"
  # or
  ntf -t "Title" "message"
  ```

- Send message with priority:
  ```
  ntf --priority high "message"
  # or
  ntf -p high "message"
  ```

- Send message with tags:
  ```
  ntf --tags tag1,tag2 "message"
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