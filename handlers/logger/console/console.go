package console

import (
	"bufio"
	"github.com/arcoz0308/arcoz0308.tech/handlers/logger"
	commands2 "github.com/arcoz0308/arcoz0308.tech/handlers/logger/console/commands"
	"os"
	"strings"
)

var Commands = make(map[string]commands2.Command, 2)

func LoadConsole() {

	Commands["ping"] = &commands2.Ping{}
	Commands["uptime"] = &commands2.Uptime{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		args := strings.Split(text, " ")

		command := args[0]
		if cmd, ok := Commands[command]; ok {
			cmd.Run(args)
		} else {
			logger.AppErrorf("command:unknown", "unknown command \"%s\"", command)
		}
	}
}
