package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Bot token")

	flag.Parse()
	if token == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Bot starting...")
	// fmt.Println(token)
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating discord session")
		return
	}

	// Register handlers here
	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Penis")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ping-Pong stuff
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
		return
	}
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
		return
	}

	// Screwing around stuff :)
	s.ChannelMessageSend(m.ChannelID, "Oof")

}
