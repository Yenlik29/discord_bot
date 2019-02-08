package config

import (
	"os"
	"fmt"
)

var (
	Token string
	BotPrefix string
)

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file := os.Getenv("TOKEN")
	if file == "" {
		return fmt.Errorf("NO token")
	}
	prefix := "!"

	Token = file
	BotPrefix = prefix

	return nil
}
