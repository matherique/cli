# 🧰 CLI toolbox

 > toolbox to create CLI in go


## Usage 

```go 

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

