package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

type commandType struct {
	command string
	typ     string
}

var cwd string
var Comm *exec.Cmd

func executeCommand(command commandType) error {
	cmd := exec.Command("cmd", "/C", command.command) // For Windows shell
	// Use ".bash" or "sh" for Linux/Mac, replace above with: exec.Command("bash", "-c", command)
	Comm = cmd
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		// InstallMinorError += ErrorStyles.Render(fmt.Sprintf("Failed to install %v, please install manually\n", command.command))
		return err
	}
	return nil
}

func executeGeneral(commands []commandType, done chan bool) error {

	for _, one := range commands {
		select {
		case <-done:
			return errors.New("key")
		default:
			switch one.typ {
			case "mkdir":
				if err := os.Mkdir(one.command, os.ModePerm); err != nil && !os.IsExist(err) {
					return err
				}
			case "cd":
				if err := os.Chdir(cwd + one.command); err != nil {
					fmt.Print(cwd + one.command)
					return err
				}
			case "exec":
				fmt.Print("Executing: " + one.command + "...")
				if err := executeCommand(one); err != nil {
					// InstallMinorError += fmt.Sprintf("Error while installing %v. Please install this manually!", one.command)
					return err
				} else {
					fmt.Println("Done")

				}
			}
		}
	}
	return nil
}

func nodeBackendFunction(index int) error {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	commands := []commandType{
		{command: "server", typ: "mkdir"},
		{command: "/server", typ: "cd"},
		{command: "echo console.log(`Hellow`); > test.js", typ: "write"},
		{command: "npm init -y", typ: "exec"},
		{command: "npm install express", typ: "exec"},
		{command: "npm install dotenv", typ: "exec"},
		{command: "npm install cors", typ: "exec"},
		{command: "npm install mongoose", typ: "exec"},
		{command: "models", typ: "mkdir"},
		{command: "controllers", typ: "mkdir"},
		{command: "routes", typ: "mkdir"},
		{command: "config", typ: "mkdir"},
		{command: "/server/config", typ: "cd"},
		// {command: "type console.log('Hellow') > database.js", typ: "exec"},
	}

	signalChan := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nReceived an interrupt, stopping execution...")
		done <- true
	}()
	switch index {
	case 0:
		return executeGeneral(commands, done)
	}
	<-done
	return nil
}
