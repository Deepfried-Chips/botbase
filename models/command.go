package models

import "github.com/bwmarrin/discordgo"

type Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)

type Command struct {
	InternalName       string
	ApplicationCommand discordgo.ApplicationCommand
	Handler            Handler
}
