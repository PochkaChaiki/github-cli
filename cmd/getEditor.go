package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

const DefaultEditor = "nano"

func CaptureInputFromEditor(editor string) (string, error) {

	file, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return "", fmt.Errorf("CaptureInputFromEditor CreateTemp(): %v", err)
	}
	filename := file.Name()

	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return "", fmt.Errorf("CaptureInputFromEditor file.Close(): %v", err)
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return "", fmt.Errorf("CaptureInputFromEditor exec.LookPath(): %v", err)
	}

	command := exec.Command(executable, filename)
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	if err = command.Run(); err != nil {
		return "", fmt.Errorf("CaptureInputFromEditor command.Run(): %v", err)
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("CaptureInputFromEditor io.ReadAll(): %v", err)
	}
	return string(bytes), nil

}
