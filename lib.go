package cmd

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Name() string
	Desc() string
	SetDesc(text string)
	Sub() []Command
	AddSub(sub Command)
	LongDesc() string
	SetLongDesc(text string)
	Handler() handler
	SetHandler(fn handler)
	HasSub(name string) Command
	Run(a []string) error
	HasAlias(name string) bool
}

type handler = func(args []string) error

type command struct {
	subcmd   []Command
	name     string
	desc     string
	longDesc string
	handler  handler
	aliases  []string
}

func New(name string, alias ...string) *command {
	cmd := new(command)
	cmd.name = name
	cmd.aliases = alias

	return cmd
}

// Name return the name of the command
func (c *command) Name() string { return c.name }

// Desc return the description of the command
func (c *command) Desc() string { return c.desc }

// SetDesc add description to the command
func (c *command) SetDesc(text string) { c.desc = text }

// AddSub add a new subcommand to the command
func (c *command) AddSub(sub Command) {
	c.subcmd = append(c.subcmd, sub)
}

// Subs return all subcommands added in command
func (c *command) Sub() []Command { return c.subcmd }

// LongDesc return the long description of the command
// this description will be show when user use '<command> help'
func (c *command) LongDesc() string { return c.longDesc }

// AddLongDesc return the long description
func (c *command) SetLongDesc(text string) { c.longDesc = text }

// Handler return the handler function that executes when the command
// subcommand is called
func (c *command) Handler() handler { return c.handler }

// SetHandler set the handler to execute when calls the command
func (c *command) SetHandler(fn handler) { c.handler = fn }

// HasAlias search and return if the command has the alias
func (c *command) HasAlias(name string) bool {
	for _, a := range c.aliases {
		if a == name {
			return true
		}
	}

	return false
}

// HasSub search and return the founded command or an error if not found
func (c *command) HasSub(name string) Command {
	for _, c := range c.Sub() {
		if c.Name() == name || c.HasAlias(name) {
			return c
		}
	}

	return nil
}

// Run run the command or subcommand based in the os args.
func (c *command) Run(args []string) error {
	h := c.Handler()

	if len(args) == 0 {
		return h(args)
	}

	if len(args) > 0 {
		if args[0] == "help" {
			fmt.Fprintln(os.Stdout, strings.Trim(c.LongDesc(), "\n"))
			return nil
		}

		s := c.HasSub(args[0])

		if s != nil {
			return s.Run(args[1:])
		}
	}

	return h(args)
}
