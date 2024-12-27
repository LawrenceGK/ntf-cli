package ntfy

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Client wraps the ntfy client CLI calls
type Client struct {
	DefaultTopic   string
	DefaultMessage string
}

// NewClient creates a new ntfy client instance
func NewClient(defaultTopic, defaultMessage string) *Client {
	return &Client{
		DefaultTopic:   defaultTopic,
		DefaultMessage: defaultMessage,
	}
}

// SendMessage sends a message to the default topic
func (c *Client) SendMessage(args []string) error {
	message := c.DefaultMessage
	if len(args) > 0 {
		message = strings.Join(args, " ")
	}

	cmd := exec.Command("ntfy", "pub", c.DefaultTopic, message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ntfy command: %w", err)
	}
	return nil
}

// SendWithOptions sends a message with additional options
func (c *Client) SendWithOptions(options map[string]string) error {
	// Preallocate slice with appropriate capacity
	cmdArgs := make([]string, 0, len(options)+3) // pub + options + topic + message
	cmdArgs = append(cmdArgs, "pub")

	// Sort options by key for consistent order (optional)
	for _, key := range []string{"in", "priority", "tags", "title"} {
		if value, ok := options[key]; ok {
			cmdArgs = append(cmdArgs, fmt.Sprintf("--%s=%s", key, value))
		}
	}

	cmdArgs = append(cmdArgs, c.DefaultTopic, c.DefaultMessage)

	cmd := exec.Command("ntfy", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ntfy command: %w", err)
	}
	return nil
}
