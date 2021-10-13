# 🧰 CLI toolbox

 > toolbox to create CLI in go


## Usage 

```go 

func ls_func(args []string) error {
  // execute ls
}

func main() {
  cmd := toolkit.New("foo")
  cmd.AddDesc("foot command description")

  subcmd := toolkit.New("bar")
  subls.AddDesc("bar subcommand description")
  
  cmd.AddSub(subcmd)
  
  if err := cmd.Run(os.Args); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
} 

```

