package main

import (
	"errors"
	"os/exec"
	"runtime"
)

func getPwd() (string, error) {
	switch runtime.GOOS {
	case "darwin", "linux":
		out, err := exec.Command("pwd").Output()
		if err != nil {
			return "", err
		}
		return string(out), nil
	case "windows":
		out, err := exec.Command("echo", "%cd%").Output()
		if err != nil {
			return "", err
		}
		return string(out), nil
	}

	return "", errors.New("your OS is not allowed, wtf?")
}
