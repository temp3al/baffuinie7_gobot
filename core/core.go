package core

import (
	"fmt"
	"log"
	"math/rand/v2"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BOT_TOKEN      string
	CMD_PREF       []string
	CHANNEL_GIMBLO string = "1355141642939535424"

	discord *discordgo.Session
)

func init() {
	CMD_PREF = []string{"baff, ", "b!"}
}

// Read incoming messages, respond to command prompts.
func On_Message_Received(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}
	log.Printf("message read @ %s from %s: \"%s\"", message.ChannelID, message.Author.Username, message.Content)

	// command handler
	for _, pref := range CMD_PREF {
		switch {
		case strings.HasPrefix(message.Content, pref) && len(message.Content) > 1:
			sc := read_command(discord, message, pref)
			if sc {
				return
			}
		}
	}

	// baffuinie reactions
	rchance := rand.IntN(1000)
	if rchance == 1 || slices.ContainsFunc(
		[]string{"baff", "uinie", "bafuinie"},
		func(word string) bool { return strings.Contains(strings.ToLower(message.Content), word) }) {
		emj_pool := []string{
			":baffled:1365660564269039636", ":baffuievil7:1365660649538981939", ":baffuinie7:1365660581083746346",
			":baffuinie7_angy:1365660594551918675", ":baffuinie7_baby:1365660715230167162", ":baffuinie7_blush:1365660758503067700",
			":baffuinie7_chop:1365660775334543441", ":baffuinie7_confuse:1365660812273909760", ":baffuinie7_dog:1365660826102403123",
			":baffuinie7_explosion:1365660840115699813", ":baffuinie7_face:1365660868804739112", ":baffuinie7_golf:1365660884076331038",
			":baffuinie7_horse:1365660899049865226", ":baffuinie7_naked:1365660913708830802", ":baffuinie7_pray:1365660927654891630",
			":baffuinie7_ramjam:1365660946017812540", ":baffuinie7_sad:1365660964875145216", ":baffuinie7_scared:1365660978183798835",
			":baffuinie7_sheep:1365660992947753062", ":baffuinie7_standing:1365661228906840095", ":baffuinie7_think:1365661251526459413",
			":baffuinie7_think2:1365660628345163786", ":baffuinie7_victory:1365661280559562833", "a:baffuinie7_wiggle:1365660612436299856",
			":muddrip:1370488496464199760"}
		emj_choice := rand.IntN(len(emj_pool))
		discord.MessageReactionAdd(message.ChannelID, message.ID, emj_pool[emj_choice])
		if rchance == 1 {
			fmt.Printf("Reacted to msg %s with \"%s\"", message.ID, emj_pool[emj_choice])
		}
	}
	// gimblo channel reactions
	if message.ChannelID == CHANNEL_GIMBLO {
		emj_pool := []string{
			":gimblo:1355141758681612397",
			":gimbgood:1365630764028268565",
			":gimblo_cowboy:1367795132526628885",
			":gimbevil:1365630793254309929"}
		emj_choice := rand.IntN(len(emj_pool))
		discord.MessageReactionAdd(message.ChannelID, message.ID, emj_pool[emj_choice])
		return
	}
}

// Called when bot joins a server.
func On_Server_Join(discord *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}
	// fmt.Println("joined!")
	// discord.ChannelMessageSend(event.Guild.SystemChannelID, "i farted")
}

func Err_stack(text string) {
	// prints traceback
	fmt.Printf("%s\n", debug.Stack())
	log.Println("^^^ " + text)
}
