package main

import (
	"botbase/constants"
	"botbase/modules"
	"flag"

	"github.com/bwmarrin/discordgo"
)

func main() {
	flag.Parse()
	var commands = modules.LoadModules(nil)

	client, err := discordgo.New("Bot " + *constants.Token)
}
