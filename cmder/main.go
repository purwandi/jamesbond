package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const (
	PORT     = 8081
	GIN_MODE = "release"
	RAHASIA  = "agent"
)

func main() {
	args := os.Args[1:]

	logrus.WithFields(logrus.Fields{
		"command": args[0],
		"args":    args[1:],
	}).Info("err")

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PORT=%d", PORT))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GIN_MODE=%s", GIN_MODE))
	cmd.Env = append(cmd.Env, fmt.Sprintf("RAHASIA=%s", RAHASIA))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Fatal(err)
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	if err := cmd.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
