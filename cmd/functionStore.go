package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var activeProcesses []*exec.Cmd

type commandType struct {
	command string
	typ     string
	content string
}

var cwd string

func executeCommand(command commandType) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command.command) // Windows shell
	} else {
		cmd = exec.Command("bash", "-c", command.command) // Unix shell
	}
	cmd.Stderr = os.Stderr
	activeProcesses = append(activeProcesses, cmd)
	if err := cmd.Run(); err != nil {
		return err
	}

	for i, proc := range activeProcesses {
		if proc == cmd {
			activeProcesses = append(activeProcesses[:i], activeProcesses[i+1:]...)
			break
		}
	}
	return nil
}

func executeGeneral(commands []commandType) error {

	for _, one := range commands {
		switch one.typ {
		case "mkdir":
			if err := os.Mkdir(one.command, os.ModePerm); err != nil && !os.IsExist(err) {
				return err
			}
		case "cd":
			if err := os.Chdir(cwd + one.command); err != nil {
				return err
			}
		case "exec":
			Program.Send(logMsg{msg: "Executing: " + one.command + "...", remove: false})

			if err := executeCommand(one); err != nil {
				return err
			} else {
				Program.Send(logMsg{msg: "Executing: " + one.command + "...Done", remove: true})
			}
		case "write":
			Program.Send(logMsg{msg: "Writing to file : " + one.command + "...", remove: false})
			if err := os.WriteFile(one.command, []byte(one.content), 0644); err != nil {
				return err
			}
			Program.Send(logMsg{msg: "Writing to file : " + one.command + "...Done", remove: true})
		case "prepend":
			Program.Send(logMsg{msg: "Prepending headers to file : " + one.command + "...", remove: false})
			existingContent, err := os.ReadFile(one.command)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			// Combine the new content with the existing content
			newContent := []byte(one.content + string(existingContent))
			// Write the combined content back to the file
			if err := os.WriteFile(one.command, newContent, 0644); err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}

			Program.Send(logMsg{msg: "Prepending headers to file : " + one.command + "...Done", remove: true})

		}

	}
	return nil
}

func nodeBackendFunction(index int, name string) error {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	commands := backendCommands(name)

	switch index {
	case 0:
		return executeGeneral(commands[:3])
	case 1:
		return executeGeneral(commands[:7])
	case 2:
		return executeGeneral(commands[:8])
	case 3:
		return executeGeneral(commands)
	}
	return nil
}

func reactNativeFunc(index int, name string) error {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	commands := reactNativeCommands(name)

	switch index {
	case 0:
		return executeGeneral(commands[:6])
	case 1:
		return executeGeneral(commands[:13])
	case 2:
		return executeGeneral(commands[:19])
	case 3:
		return executeGeneral(commands[:30])
	case 4:
		return executeGeneral(commands)
	}
	return nil
}

func reactFunc(index int, name string) error {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	commands := reactCommands(name)

	switch index {
	case 0:
		return executeGeneral(commands[:4])
	case 1:
		return executeGeneral(commands[:12])
	case 2:
		return executeGeneral(commands[:21])
	case 3:
		return executeGeneral(commands)
	case 4:
		var mui = commandType{command: "npm install @mui/material @emotion/react @emotion/styled @fontsource/roboto @mui/icons-material", typ: "exec"}
		newComm := append(commands, mui)
		return executeGeneral(newComm)
	case 5:
		var antd = commandType{command: "npm install antd --save", typ: "exec"}
		newComm := append(commands, antd)
		return executeGeneral(newComm)
	}
	return nil
}

func mernFunc(index int, name string) error {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	reactCommands, backCommands := mernCommands(name)
	var commands = []commandType{
		{command: name, typ: "mkdir"},
		{command: "/" + name, typ: "cd"},
	}
	switch index {
	case 0:
		commands = append(commands, reactCommands[:4]...)
		commands = append(commands, commandType{command: "/" + name, typ: "cd"})
		commands = append(commands, backCommands[:3]...)
		return executeGeneral(commands)
	case 1:
		commands = append(commands, reactCommands[:21]...)
		commands = append(commands, commandType{command: "/" + name, typ: "cd"})
		commands = append(commands, backCommands...)
		return executeGeneral(commands)
	case 2:
		commands = append(commands, reactCommands...)
		commands = append(commands, commandType{command: "/" + name, typ: "cd"})
		commands = append(commands, backCommands...)
		return executeGeneral(commands)
	case 3:
		commands = append(commands, reactCommands...)
		commands := append(commands, commandType{command: "npm install @mui/material @emotion/react @emotion/styled @fontsource/roboto @mui/icons-material", typ: "exec"})
		commands = append(commands, commandType{command: "/" + name, typ: "cd"})
		commands = append(commands, backCommands...)
		return executeGeneral(commands)
	case 4:
		commands = append(commands, reactCommands...)
		commands := append(commands, commandType{command: "npm install antd --save", typ: "exec"})
		commands = append(commands, commandType{command: "/" + name, typ: "cd"})
		commands = append(commands, backCommands...)
		return executeGeneral(commands)

	}
	return nil
}
