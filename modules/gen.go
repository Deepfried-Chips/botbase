package modules

import (
	"botbase/models"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func NewGen() []models.Command {
	return Gen
}

var Gen = []models.Command{
	{
		InternalName: "ping",
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "Check latency",
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Pong! Latency: %dms", s.HeartbeatLatency().Milliseconds()),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		},
	},
}
