# ntf-cli

ntf-cli 是一个简单的命令行工具，封装了 ntfy 客户端命令行工具，旨在简化消息发送的过程。该工具支持在 Windows 和 Linux 平台上运行。

## 特性

- 配置默认主题和默认消息
- 通过简单的命令行参数发送消息
- 支持丰富的消息选项：
  - 延迟发送 (--in)
  - 消息标题 (--title)
  - 消息优先级 (--priority)
  - 消息标签 (--tags)
- 支持自定义主题和消息内容

## 安装

1. 确保已安装 Go 语言环境和 ntfy 客户端
2. 克隆该项目：
   ```
   git clone https://github.com/yourusername/ntf-cli.git
   ```
3. 进入项目目录并安装依赖：
   ```
   cd ntf-cli
   go mod tidy
   ```

## 使用示例

基本用法：

- 发送默认消息到默认主题：

  ```
  ntf
  ```

- 发送带有时间参数的消息：

  ```
  ntf --in 30min
  ```

- 发送自定义消息到默认主题：

  ```
  ntf "自定义消息"
  ```

## 配置

配置文件会自动创建在以下位置：

- Windows: `%LOCALAPPDATA%\ntf-cli\config.json`
- Linux/macOS: `~/.config/ntf-cli/config.json`

配置文件格式(JSON)：
```
{
  "default_topic": "your_default_topic",
  "default_message": "your_default_message"
}
```

默认主题和消息可以在 `internal/config/config.go` 文件中进行设置。请根据需要修改该文件。

## 贡献

欢迎提交问题和贡献代码！请遵循贡献指南。

## 许可证

该项目采用 MIT 许可证，详细信息请查看 LICENSE 文件。