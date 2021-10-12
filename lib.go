package toolkit

type Command interface {
	Name() string
	Desc() string
}

type command struct {
	subcmd *[]command
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
func (c *command) Desc() string { return c.desc }
