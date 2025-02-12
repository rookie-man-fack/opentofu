package command

import (
	"github.com/mitchellh/cli"
)

// AliasCommand is a Command implementation that wraps another Command for the purpose of aliasing.
type AliasCommand struct {
	cli.Command
}

func (c *AliasCommand) Run(args []string) int {
	return c.Command.Run(args)
}

func (c *AliasCommand) Help() string {
	return c.Command.Help()
}

func (c *AliasCommand) Synopsis() string {
	return c.Command.Synopsis()
}
