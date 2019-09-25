package main

import (
	"fmt"
	"go.starlark.net/starlark"
	"os"
)

var TestFileSysRepo = "C:/Users/chris/dev/opcg/goz/mock/filesysrepo"


func main() {
	args := os.Args[1:]

	if args[0] == "exec" {
		runConfig(args[1])
	}
}

type Env struct {
	starlark.HasAttrs
}

func (t *Env) Append(key, value string) {
	fmt.Printf("Got %s %s", key, value)
}

func execConfig(path string)  {

	env := Env{}

	args := starlark.Tuple{env}

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "package thread"}
	globals, err := starlark.ExecFile(thread, path, nil, nil)
	if err != nil {  }

	// Retrieve a module global.
	commands := globals["commands"]

	// Call Starlark function from Go.
	v, err := starlark.Call(thread, commands, args, nil)
	if err != nil {  }
	fmt.Printf("commands() = %v\n", v)

}
