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
	// Get config file path
	configPath, err := config.GetConfigPath()
	if err != nil {
		log.Fatalf("Failed to get config file path: %v", err)
	}

	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		// Create default config file if it doesn't exist
		if os.IsNotExist(err) {
			cfg = &config.Config{
				DefaultTopic:   "default-topic",
				DefaultMessage: "default-message",
			}
			err = config.SaveConfig(configPath, cfg)
			if err != nil {
				log.Fatalf("Failed to create default config file: %v", err)
			}
			fmt.Printf("Default config file created at %s\n", configPath)
		} else {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

	// Define command line flags
	flags := struct {
		in       string
		at       string // New at field
		message  string
		topic    string
		title    string
		priority string
		tags     string
	}{}

	// Add long and short format parameters
	flag.StringVar(&flags.in, "in", "", "Set the message schedule time")
	flag.StringVar(&flags.in, "i", "", "Set the message schedule time (shorthand)")
	flag.StringVar(&flags.at, "at", "", "Set the message schedule time (same as --in)")
	flag.StringVar(&flags.at, "a", "", "Set the message schedule time (same as -i)")
	flag.StringVar(&flags.message, "msg", cfg.DefaultMessage, "Message content to send")
	flag.StringVar(&flags.message, "m", cfg.DefaultMessage, "Message content to send (shorthand)")
	flag.StringVar(&flags.topic, "topic", cfg.DefaultTopic, "Topic to publish to")
	flag.StringVar(&flags.title, "title", "", "Message title")
	flag.StringVar(&flags.title, "t", "", "Message title (shorthand)")
	flag.StringVar(&flags.priority, "priority", "", "Message priority (min|low|default|high|urgent)")
	flag.StringVar(&flags.priority, "p", "", "Message priority (shorthand)")
	flag.StringVar(&flags.tags, "tags", "", "Message tags (comma-separated)")

	flag.Parse()

	// Create ntfy client with command line parameters
	client := ntfy.NewClient(flags.topic, flags.message)

	// Send default message if no parameters provided
	if len(os.Args) == 1 {
		if err := client.SendMessage(nil); err != nil {
			log.Fatalf("Failed to send default message: %v", err)
		}
		fmt.Println("Default message sent")
		return
	}

	// Collect all non-empty options
	options := make(map[string]string)
	if flags.in != "" {
		options["in"] = flags.in
	} else if flags.at != "" {
		options["at"] = flags.at
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

	// Decide sending method based on options
	if len(options) > 0 {
		err = client.SendWithOptions(options)
	} else {
		err = client.SendMessage(flag.Args())
	}

	if err != nil {
		log.Fatal(err)
	}
}
