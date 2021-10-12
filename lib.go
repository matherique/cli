package toolkit

type Command interface {
	Name() string
	Desc() string
	AddSub(sub Command)
}

type command struct {
	subcmd []*command
	name   string
	desc   string
	exec   func(args []string) error
}

func New(name, desc string) *command {
	cmd := new(command)
	cmd.name = name
	cmd.desc = desc

	return cmd
}

func (c *command) Name() string { return c.name }

// Desc return the description of the command
func (c *command) Desc() string { return c.desc }

func (c *command) AddSub(sub *command) {
	c.subcmd = append(c.subcmd, sub)
}

// Subs return all subcommands added in command
func (c *command) Subs() []*command { return c.subcmd }
