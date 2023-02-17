package main

import (
	"log"
	"os/exec"
)

func main(){
	// Read and display the capabilities of the running process
	// c := cap.GetProc()
	// log.Printf("this process has these caps:", c)
	command := exec.Command("getpcaps", "$$")
	output, err := command.CombinedOutput()
	log.Printf("output: ", string(output))

	if err != nil {
		log.Printf("error: ", err) 
	}
}
