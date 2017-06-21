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
	cmd "github.com/ziemerz/gogobot/commands"
)

func init(){
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string

func main() {
	// Check if token was provided
	if token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

	// New discordgo
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

	if strings.HasPrefix(mc.Content, "!gogo") {
		channelID := mc.ChannelID
		if strings.Contains(mc.Content, "cat"){
			rnd := cmd.NewRandom()
			session.ChannelMessageSend(channelID, rnd.Subcommands["cat"]())
		} else {
			session.ChannelMessageSend(channelID, "You called me sir?")
		}
		return
	}


}