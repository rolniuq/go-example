package main

import (
	"os"
	"os/exec"
	"syscall"
)

// some time we want to replace current go process to another one
// this is an example to do that

func main() {
	binary, lookerr := exec.LookPath("ls")
	if lookerr != nil {
		panic(lookerr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

//Note that Go does not offer a classic Unix fork function. Usually this isn’t an issue though, since starting goroutines, spawning processes, and exec’ing processes covers most use cases for fork
