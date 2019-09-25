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


func execConfig(path string)  {

	env = make(map(string[string]))

	// Execute Starlark program in a file.
	thread := &starlark.Thread{Name: "my thread"}
	globals, err := starlark.ExecFile(thread, path, nil, nil)
	if err != nil {  }

	// Retrieve a module global.
	commands := globals["commands"]

	// Call Starlark function from Go.
	v, err := starlark.Call(thread, commands, nil, nil)
	if err != nil {  }
	fmt.Printf("hello() = %v\n", v)

}
