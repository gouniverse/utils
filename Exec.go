package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Exec executes a system command using the the standard package "os/exec".
func Exec(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ExecLine executes a full system command line string containing the arguments using for the standard package "os/exec".
func ExecLine(cmd string) {
	if cmd == "" {
		log.Println("A blank command")
		return
	}
	cs := strings.Split(cmd, " ")
	Exec(cs[0], cs[1:]...)
	log.Println("Executed successfully")
}
