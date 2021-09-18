package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type AddCommand struct{}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	fmt.Println("あざます")
	return 0
}

func main() {
	c := cli.NewCLI("todo", "0.1.0")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{}, nil
		},
	}

	exitCode, err := c.Run()
	if err != nil {
		fmt.Printf("Failed to execute: %s¥n", err.Error())
	}

	os.Exit(exitCode)
}
