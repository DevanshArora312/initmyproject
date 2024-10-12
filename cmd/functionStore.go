package cmd

import (
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

	commands := backendCommands

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
		return executeGeneral(commands[:12])
	case 2:
		return executeGeneral(commands[:18])
	case 3:
		return executeGeneral(commands[:29])
	case 4:
		return executeGeneral(commands)
	}
	return nil
}
