package toolkit_test

import (
	"fmt"

	toolkit "github.com/matherique/cli-toolkit"
)

func ExampleNew() {
	cmd := toolkit.New("foo")

	fmt.Println(cmd.Name())

	// Output:
	// foo
}

func ExampleAddSub() {
	cmd := toolkit.New("foo")
	sub := toolkit.New("bar")

	cmd.AddSub(sub)
	fmt.Println(len(cmd.Sub()))

	// Output:
	// 1
}

func ExampleSetDesc() {
	cmd := toolkit.New("foo")

	cmd.SetDesc("foo foo")

	fmt.Println(cmd.Desc())
	// Output:
	// foo foo
}

func ExampleSetLongDesc() {
	cmd := toolkit.New("foo")
	cmd.SetLongDesc("long foo")

	fmt.Println(cmd.LongDesc())

	// Output:
	// long foo
}

func ExampleSetHandler() {
	cmd := toolkit.New("foo")
	handler := func(args []string) error {
		fmt.Println("foo handler")
		return nil
	}

	cmd.SetHandler(handler)

	h := cmd.Handler()
	arg := make([]string, 0)
	h(arg)

	// Output:
	// foo handler
}

func ExampleHasSub() {
	cmd := toolkit.New("foo")
	sub := toolkit.New("bar")

	cmd.AddSub(sub)

	s, err := cmd.HasSub("bar")

	if err != nil {
		fmt.Printf("error %v", err)
	}

	fmt.Println(s.Name())

	// Output:
	// bar
}

func ExampleRun() {
	cmd := toolkit.New("foo")
	cmd.SetLongDesc("foo long desc")
	cmd.SetHandler(func(a []string) error {
		fmt.Println("foo func")
		return nil
	})

	sub := toolkit.New("bar")
	sub.SetHandler(func(a []string) error {
		fmt.Println("bar func")
		return nil
	})

	cmd.AddSub(sub)

	var a []string
	cmd.Run(a)

	a = []string{"bar"}
	cmd.Run(a)

	a = []string{"wrong"}
	err := cmd.Run(a)

	if err != nil {
		fmt.Println(err)
	}

	a = []string{"help"}
	cmd.Run(a)

	// Output:
	// foo func
	// bar func
	// no subcommand found
	// foo long desc
}
