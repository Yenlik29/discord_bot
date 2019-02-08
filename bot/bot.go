package bot

import (
	"trial/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("BOT is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!help" {
			_,_ = s.ChannelMessageSend(m.ChannelID, "Hi! Here are some commands for use:\n[!help]\nand etc")
		} else {
			_,_ = s.ChannelMessageSend(m.ChannelID, "[BOT]: commnad not found\nType [!help] for more info")
		}
	}

	fmt.Println(m.Content)
}

