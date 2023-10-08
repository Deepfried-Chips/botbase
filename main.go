package main

import (
	"botbase/constants"
	"botbase/models"
	"botbase/modules"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	client             *discordgo.Session
	commands           []models.Command
	registeredCommands []*discordgo.ApplicationCommand
	err                error
)

func main() {
	flag.Parse()
	commands = modules.LoadModules(nil)

	client, err = discordgo.New("Bot " + *constants.Token)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err = client.Open()
	if err != nil {
		log.Fatalf("Cannot open session: %v", err)
	}

	client.AddHandler(HandleSlashCommands)

	RegisterCommands(client)

	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Send interrupt signal (Ctrl+C) to end session")
	<-stop
}
