package cmd

import (
	"bytes"
	"fmt"
	"sandhu-sahil/bot/variables"
)

func HelpCommandIntractions() string {
	buffer := bytes.NewBufferString("Commands: \n")
	for _, cmd := range variables.CreatedCommands {
		// return all commands in text format
		msg := fmt.Sprintf("\t %s - %s\n", cmd.Name, cmd.Description)
		buffer.WriteString(msg)
	}
	str := buffer.String()
	return str
}
