package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + "ENTER TOKEN HERE")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "?" reply with "It is possible!"
	if m.Content == "?" {
		nakResponse := "It is possible.\n"
		switch random := rand.Intn(6); random {
		case 0:
			nakResponse += "warm regards"
		case 1:
			nakResponse += "take care"
		case 2:
			nakResponse += "best"
		case 3:
			nakResponse += "sincerely"
		case 4:
			nakResponse += "thanks"
		case 5:
			nakResponse += "thank you"
		default:
			nakResponse += "warm regards"
		}
		nakResponse += ",\nNaka"
		s.ChannelMessageSend(m.ChannelID, nakResponse)
	}
}
