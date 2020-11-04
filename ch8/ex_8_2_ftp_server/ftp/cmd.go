package ftp

import "strings"

type Cmd struct {
	cmd  string
	args []string
}

func (cmd *Cmd) Exec() (string, error) {
	return "exec", nil
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
