package main

import (
	"fmt"
	"trial/config"
	"trial/bot"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()
	<-make(chan struct{})
	return
}
