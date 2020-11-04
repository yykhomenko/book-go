package ftp

import (
	"fmt"
	"strings"
)

type Cmd struct {
	Name string
	Args []string
}

func (cmd *Cmd) Exec() (string, error) {
	return fmt.Sprintf("exec %s(%s)", cmd.Name, strings.Join(cmd.Args, ", ")), nil
}

func NewCmd(line string) *Cmd {
	words := strings.Split(line, " ")
	cmd := strings.TrimSpace(words[0])
	args := words[1:]
	for i, arg := range args {
		args[i] = strings.TrimSpace(arg)
	}
	return &Cmd{cmd, args}
}
