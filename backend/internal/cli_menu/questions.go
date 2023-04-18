package cli_menu

import "github.com/AlecAivazis/survey/v2"

var qs = []*survey.Question{
	{
		Name: "command",
		Prompt: &survey.Select{
			Message: "Select Command",
			Options: []string{"Generate new wallet", "Import wallet", "okx", "exit"},
			Default: nil,
		},
	},
}

//var qs = []*survey.Question{
//	{
//		Name:      "name",
//		Prompt:    &survey.Input{Message: "What is your name?"},
//		Validate:  survey.Required,
//		Transform: survey.Title,
//	},
//	{
//		Name: "color",
//		Prompt: &survey.Select{
//			Message: "Choose a color:",
//			Options: []string{"red", "blue", "green"},
//			Default: "red",
//		},
//	},
//	{
//		Name:   "age",
//		Prompt: &survey.Input{Message: "How old are you?"},
//	},
//}
