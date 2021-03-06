package main

import (
	"github.com/bwmarrin/discordgo"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"
)

func init(){
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string
var handler = NewCommandHandler()

func main() {
	// Check if token was provided
	if token == "" {
		fmt.Println("No token provided. Please run: gogobot -t <bot token>")
		return
	}

	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("Error creating discordgo session ", err)
	}

	// Add a handler for when we receive the ready event from Discord
	discord.AddHandler(ready)

	// Add handler for handling messages
	discord.AddHandler(messageCreate)

	// Open a session with discord.
	err = discord.Open()

	if err != nil {
		log.Fatal("Error opening session to discord ", err)
	}

	// Run until terminated
	fmt.Println("Gogobot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc

	// Clean up and close Discord connection
	discord.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "!gogo")
}

func messageCreate(session *discordgo.Session, mc *discordgo.MessageCreate){

	//Ignore own messages
	if mc.Author.ID == session.State.User.ID { return }
	message := strings.Split(mc.Content, " ")
	if message[0] == "!gogo" {
		channelID := mc.ChannelID
		if len(message) > 1 {
			session.ChannelMessageSend(channelID, handler.HandleCommand(message[1:]))
			return
		}  else {
			session.ChannelMessageSend(channelID, "You called me sir? Type in `!gogo help` if you need any help")
			return
		}
	}
}