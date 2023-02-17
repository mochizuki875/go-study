package main

import (
	"fmt"
	"os/exec"
)

func main(){
	target := "/tmp"
	command := exec.Command("umount", target)
	output, err := command.CombinedOutput()
	fmt.Println("output: ",string(output))
	fmt.Println("err: ",err)

}