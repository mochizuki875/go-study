package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("uname", "-r") // dummy command return status 0
	if _, err := command.CombinedOutput(); err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Printf("command: %#v\n", command)
	fmt.Printf("*command: %#v\n", *command)
	fmt.Printf("*command.ProcessState: %#v\n", *command.ProcessState)
}
