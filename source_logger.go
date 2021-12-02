package main

import (
	"fmt"
	"github.com/creack/pty"
	"io"
	"syscall"
	"os"
	"os/exec"
	"os/signal"    
)

func main() {
	args := os.Args[2:]

	// Calling the srcds_linux executable in the same directory
	cmd := exec.Command(os.Args[1], args...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = append(cmd.Env, "LD_LIBRARY_PATH=.:bin:"+os.Getenv("LD_LIBRARY_PATH"))

	// Redirecting the SIGINT or SIGKILL signal to the srcds_linux
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			err := cmd.Process.Signal(sig)
			if err != nil {
				fmt.Printf("[sourcelogger] couldn't redirect signal %v to srcds_linux\n", sig)
			}
		}
	}()

	// Starting the pseudo terminal for catching the stdout of gmod
	file, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("[sourcelogger] could't start the srcds_linux executable")
		panic(err)
	}

	// Redirecting the output of gmod to the stdout
	go func() {
		fmt.Println("[sourcelogger] Started redirecting stdin")
		_, err = io.Copy(file, os.Stdin)
		if err != nil {
			fmt.Printf("[sourcelogger] Error redirecting stdin: %s\n", err)
		}
		fmt.Println("[sourcelogger] Finished redirecting stdin")
	 }()
	// Redirecting the stdin to input of gmod
	_, _ = io.Copy(os.Stdout, file)
}
