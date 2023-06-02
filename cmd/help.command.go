package cmd

import (
	"bytes"
	"fmt"
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/variables"
)

func HelpCommand(ctx framework.Context) {
	cmds := ctx.CmdHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands: \n")
	for cmdName, cmdStruct := range cmds {
		if len(cmdName) == 1 {
			continue
		}
		// fmt.Println(cmdName, cmdStruct)
		msg := fmt.Sprintf("\t %s%s - %s\n", variables.PREFIX, cmdName, cmdStruct.GetHelp())
		buffer.WriteString(msg)
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}
