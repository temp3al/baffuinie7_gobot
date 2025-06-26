package main

import (
	"baff-discordbot/control"
	"baff-discordbot/core"
	"bufio"
	"fmt"

	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	BOT_TOKEN       string
	CMD_PREF        string = "+"
	DISCORD_SESSION *discordgo.Session
)

func main() {
	// enviroment, token
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	BOT_TOKEN, e := os.LookupEnv("BOT_TOKEN")
	if !e {
		log.Fatal("\"BOT_TOKEN\" variable does not exist.")
	}

	// connect to main database
	// db_main, err := sql.Open("sqlite3", "./database/main.db")
	// if err != nil {
	// log.Panicln(err)
	// }
	// fmt.Println(db_main)
	// defer db_main.Close()

	// start up the bot
	println("Running bot with token:", BOT_TOKEN)
	// bot launch
	DISCORD_SESSION, err := discordgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	//
	control.DISCORD_SESSION = DISCORD_SESSION

	// discord event handlers
	DISCORD_SESSION.AddHandler(core.On_Message_Received)
	DISCORD_SESSION.AddHandler(core.On_Server_Join)
	DISCORD_SESSION.Open()
	defer DISCORD_SESSION.Close()
	// set custom display status
	// TODO: move this to "core.go"
	err = DISCORD_SESSION.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online", // "online", "idle", "dnd", or "invisible"
		Activities: []*discordgo.Activity{
			{
				Name: "Baff7 Party",
				Type: discordgo.ActivityTypeGame,
				// State: "We be codin'",
			},
		},
	})
	if err != nil {
		println("Failed to set custom status:", err.Error())
	}

	fmt.Println("bot running! (Exit with CTRL + C)")

	// initiate infinite loop to keep bot running
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Quitting...")
			break
		}
		control.C_interpret(input)

	}
}
