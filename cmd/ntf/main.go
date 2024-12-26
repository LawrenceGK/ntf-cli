package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ntf-cli/internal/config"
	"ntf-cli/internal/ntfy"
)

func main() {
	// 获取配置文件路径
	configPath, err := config.GetConfigPath()
	if err != nil {
		log.Fatalf("无法获取配置文件路径: %v", err)
	}

	// 加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		// 如果配置文件不存在，创建一个默认配置文件
		if os.IsNotExist(err) {
			cfg = &config.Config{
				DefaultTopic:   "default-topic",
				DefaultMessage: "default-message",
			}
			err = config.SaveConfig(configPath, cfg)
			if err != nil {
				log.Fatalf("无法创建默认配置文件: %v", err)
			}
			fmt.Printf("已在 %s 创建默认配置文件\n", configPath)
		} else {
			log.Fatalf("无法加载配置: %v", err)
		}
	}

	// 定义命令行参数
	flags := struct {
		in       string
		message  string
		topic    string
		title    string
		priority string
		tags     string
	}{}

	flag.StringVar(&flags.in, "in", "", "设置消息的时间参数")
	flag.StringVar(&flags.message, "msg", cfg.DefaultMessage, "要发送的消息")
	flag.StringVar(&flags.topic, "topic", cfg.DefaultTopic, "要发送的主题")
	flag.StringVar(&flags.title, "title", "", "消息标题")
	flag.StringVar(&flags.priority, "priority", "", "消息优先级 (min|low|default|high|urgent)")
	flag.StringVar(&flags.tags, "tags", "", "消息标签 (逗号分隔)")

	flag.Parse()

	// 创建ntfy客户端，使用命令行参数覆盖默认值
	client := ntfy.NewClient(flags.topic, flags.message)

	// 如果没有提供任何参数，发送默认消息
	if len(os.Args) == 1 {
		if err := client.SendMessage(nil); err != nil {
			log.Fatalf("发送默认消息失败: %v", err)
		}
		fmt.Println("已发送默认消息")
		return
	}

	// 收集所有非空选项
	options := make(map[string]string)
	if flags.in != "" {
		options["in"] = flags.in
	}
	if flags.title != "" {
		options["title"] = flags.title
	}
	if flags.priority != "" {
		options["priority"] = flags.priority
	}
	if flags.tags != "" {
		options["tags"] = flags.tags
	}

	// 根据是否有选项决定发送方式
	if len(options) > 0 {
		err = client.SendWithOptions(options)
	} else {
		err = client.SendMessage(flag.Args())
	}

	if err != nil {
		log.Fatal(err)
	}
}
