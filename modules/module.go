package modules

import "botbase/models"

var ModuleList = []string{
	"gen",
}

func LoadModule(module string) []models.Command {
	switch module {
	case "gen":
		return NewGen()
	default:
		return nil
	}
}

func LoadModules(modules []string) *[]models.Command {
	var list []models.Command
	if modules != nil {
		for _, module := range modules {
			LoadModule(module)
			list = append(list, LoadModule(module)...)
		}
	} else {
		for _, module := range ModuleList {
			list = append(list, LoadModule(module)...)
		}
	}
	return &list
}
