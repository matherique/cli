package toolkit

type Command interface {
	Name() string
	Desc() string
	AddSub(sub Command)
	AddDesc(text string)
	Sub() []*command
}

type command struct {
	subcmd []*command
	name   string
	desc   string
	exec   func(args []string) error
}

func New(name string) *command {
	cmd := new(command)
	cmd.name = name

	return cmd
}

// Name return the name of the command
func (c *command) Name() string { return c.name }

// Desc return the description of the command
func (c *command) Desc() string { return c.desc }

// AddDesc add description to the command
func (c *command) AddDesc(text string) { c.desc = text }

// AddSub add a new subcommand to the command
func (c *command) AddSub(sub *command) {
	c.subcmd = append(c.subcmd, sub)
}

// Subs return all subcommands added in command
func (c *command) Sub() []*command { return c.subcmd }
