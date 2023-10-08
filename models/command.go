package models

import "github.com/bwmarrin/discordgo"

type Command struct {
	InternalName       string
	ApplicationCommand *discordgo.ApplicationCommand
	Handler            interface{}
}
