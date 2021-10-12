# 🧰 CLI toolbox

 > toolbox to create CLI in go


## Usage 

```go 

func ls_func(args []string) error {
  // execute ls
}

func main() {
  cmd := toolkit.New("command")
  cmd.AddDesc("command description")

  subls := cmd.AddSub("ls", ls_func)
  subls.AddDesc("ls subcommand description")

  err := cmd.Run(os.Args)
  
  if err := nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
} 

```

