package ntfy

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Client 封装了ntfy客户端的调用
type Client struct {
	DefaultTopic   string
	DefaultMessage string
}

// NewClient 创建一个新的ntfy客户端
func NewClient(defaultTopic, defaultMessage string) *Client {
	return &Client{
		DefaultTopic:   defaultTopic,
		DefaultMessage: defaultMessage,
	}
}

// SendMessage 发送消息到默认主题
func (c *Client) SendMessage(args []string) error {
	message := c.DefaultMessage
	if len(args) > 0 {
		message = strings.Join(args, " ")
	}

	cmd := exec.Command("ntfy", "pub", c.DefaultTopic, message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("执行 ntfy 命令失败: %w", err)
	}
	return nil
}

// SendWithOptions 发送带有选项的消息
func (c *Client) SendWithOptions(options map[string]string) error {
	// 预分配合适的切片容量
	cmdArgs := make([]string, 0, len(options)+3) // pub + options + topic + message
	cmdArgs = append(cmdArgs, "pub")

	// 为了确保参数顺序一致，可以按字母顺序排序（可选）
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
		return fmt.Errorf("执行 ntfy 命令失败: %w", err)
	}
	return nil
}
