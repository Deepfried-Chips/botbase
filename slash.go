package main

import (
	"botbase/models"
	"github.com/bwmarrin/discordgo"
	"log"
)

func HandleSlashCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

func RegisterCommands(s *discordgo.Session) {
	for i, v := range commands {
		log.Printf("Index %v: %v", i, v)
		/*cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v.ApplicationCommand)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.InternalName, err)
		}
		registeredCommands[i] = cmd*/
	}
}

func DeregisterCommands(s *discordgo.Session, commands *[]models.Command) {

}
