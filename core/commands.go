package core

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Interpret incoming Discord message as a command.
func read_command(discord *discordgo.Session, message *discordgo.MessageCreate, pref string) bool {
	msg := strings.TrimPrefix(message.Content, pref) // trim prefix
	if len(msg) < 1 {
		return false
	}

	msplit := strings.Split(msg, " ")
	cmd := msplit[0]
	// param := msplit[1:]

	switch cmd {
	case "help":
		discord.ChannelMessageSend(message.ChannelID, "Help command executed.")
		return true
	case "ping":
		msg = fmt.Sprintf("### Pong! 🏓\n-# %sms heartbeat latency", fmt.Sprint(discord.HeartbeatLatency().Milliseconds()))
		discord.ChannelMessageSend(message.ChannelID, msg)
		return true
	}
	return false
}
