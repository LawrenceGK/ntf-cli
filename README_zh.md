# ntf-cli

中文 | [English](README.md)

一个简单的命令行工具，封装了 ntfy 客户端以简化消息发送流程。

## 前置要求
- 安装 [ntfy 客户端](https://docs.ntfy.sh/install/) 并确保可在命令行中使用

## 安装

### Windows
1. 从 [Releases](https://github.com/LawrenceGK/ntf-cli/releases) 下载 `ntf.exe`
2. 将程序所在目录添加到系统 PATH

### Linux
方式一：从发布版本安装（推荐）
1. 从 [Releases](https://github.com/LawrenceGK/ntf-cli/releases) 页面下载最新的 `ntf` 二进制文件
2. 将二进制文件移动到 `/usr/local/bin` 目录下，并重命名为 `ntf`：
   ```bash
   sudo mv ntf-cli /usr/local/bin/ntf
   ```
3. 确保二进制文件具有执行权限：
   ```bash
   sudo chmod +x /usr/local/bin/ntf
   ```

方式二：从源码构建
1. 确保已安装 Go 语言环境和 ntfy 客户端
2. 克隆该项目：
   ```bash
   git clone https://github.com/lawgk/ntf-cli.git
   ```
3. 进入项目目录并安装依赖：
   ```bash
   cd ntf-cli
   go mod tidy
   ```
4. 构建二进制文件：
   ```bash
   go build -o ntf cmd/ntf/main.go
   ```
5. 移动到系统PATH目录（可选）：
   ```bash
   sudo mv ntf /usr/local/bin/
   ```

## 使用示例

基本用法：

- 发送默认消息到默认主题：

  ```bash
  ntf
  ```

- 发送带有时间参数的消息：

  ```bash
  ntf --in 30min
  ```

- 发送自定义消息到默认主题：

  ```bash
  ntf "自定义消息"
  ```

- 发送带有标题的消息：

  ```bash
  ntf --title "标题" "消息内容"
  ```

- 发送带有优先级的消息：

  ```bash
  ntf --priority high "消息内容"
  ```

- 发送带有标签的消息：

  ```bash
  ntf --tags tag1,tag2 "消息内容"
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