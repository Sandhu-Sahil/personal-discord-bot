package cmd

import (
	"bytes"
	"fmt"
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/variables"
)

func HelpCommandIntractions() string {
	buffer := bytes.NewBufferString("```")
	buffer.WriteString("Commands->> \n")

	spacing := 30

	for _, cmd := range variables.CreatedCommands {
		// return all commands in text format
		msg := fmt.Sprintf("\t %s : %s %s\n", cmd.Name, framework.SpacingIndentation(spacing, len(cmd.Name)), cmd.Description)
		buffer.WriteString(msg)
	}
	buffer.WriteString("```")
	str := buffer.String()
	return str
}
