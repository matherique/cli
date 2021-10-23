package cmd_test

import (
	"fmt"

	cmd "github.com/matherique/cmd"
)

func ExampleNew() {
	c := cmd.New("foo")

	fmt.Println(c.Name())

	// Output:
	// foo
}

func ExampleAddSub() {
	c := cmd.New("foo")
	sc := cmd.New("bar")

	c.AddSub(sc)
	fmt.Println(len(c.Sub()))

	// Output:
	// 1
}

func ExampleSetDesc() {
	c := cmd.New("foo")

	c.SetDesc("foo foo")

	fmt.Println(c.Desc())
	// Output:
	// foo foo
}

func ExampleSetLongDesc() {
	c := cmd.New("foo")
	c.SetLongDesc("long foo")

	fmt.Println(c.LongDesc())

	// Output:
	// long foo
}

func ExampleSetHandler() {
	c := cmd.New("foo")
	handler := func(args []string) error {
		fmt.Println("foo handler")
		return nil
	}

	c.SetHandler(handler)

	h := c.Handler()
	arg := make([]string, 0)
	h(arg)

	// Output:
	// foo handler
}

func ExampleHasSub() {
	c := cmd.New("foo")
	sc := cmd.New("bar")

	c.AddSub(sc)

	s, err := c.HasSub("bar")

	if err != nil {
		fmt.Printf("error %v", err)
	}

	fmt.Println(s.Name())

	// Output:
	// bar
}

func ExampleRun() {
	c := cmd.New("foo")
	c.SetLongDesc("foo long desc")
	c.SetHandler(func(a []string) error {
		fmt.Println("foo func")
		return nil
	})

	sc := cmd.New("bar")
	sc.SetHandler(func(a []string) error {
		fmt.Println("bar func")
		return nil
	})

	c.AddSub(sc)

	var a []string
	c.Run(a)

	a = []string{"bar"}
	c.Run(a)

	a = []string{"wrong"}
	err := c.Run(a)

	if err != nil {
		fmt.Println(err)
	}

	a = []string{"help"}
	c.Run(a)

	// Output:
	// foo func
	// bar func
	// no subcommand found
	// foo long desc
}
