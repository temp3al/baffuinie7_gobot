package control

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	DISCORD_SESSION *discordgo.Session
)

// Schedule a bot restart.
func Schedule_Reboot() {
}

// Restart the bot.
func Reboot() {
}

// Shutdown the bot and end all processes.
//
// Only ever use this when going under heavy mantainance, else
// use "Schedule_Reboot" after any updates for minimal interruptions.
func Shutdown() {
}

// Interpet terminal messages.
func C_interpret(message string) {
	msplice := strings.Split(message, " ")
	cmd := strings.ToLower(msplice[0])
	param := msplice[1:]

	if cmd == "speak" && len(param) >= 2 {
		speak(param[0], strings.Join(param[1:], " "))
	}
}

// Send a message.
func speak(channel string, message string) {
	_, err := DISCORD_SESSION.ChannelMessageSend(channel, message)
	if err != nil {
		log.Println(err)
	}
}
