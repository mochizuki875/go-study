package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("uname", "-r") // dummy command return status 0
	output, err := command.CombinedOutput()

	// if err := command.Run(); err != nil {
	// 	fmt.Println("err: ", err)
	// }

	fmt.Printf("output: %#v\n", string(output))
	fmt.Printf("err: %#v\n", err)
	fmt.Println("command.ProcessState", command.ProcessState)
	fmt.Printf("command.ProcessState: %#v\n", command.ProcessState)

	// fmt.Printf("command: %#v\n", command)
	fmt.Printf("*command: %#v\n", *command)
	fmt.Printf("*command.ProcessState: %#v\n", *command.ProcessState)
}
