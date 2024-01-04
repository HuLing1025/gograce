package executer

import (
	"bytes"
	"os/exec"
)

// IExecuter interfaces of Executer which should implement.
type IExecuter interface {
	Execute() (outputBuffer bytes.Buffer, err error)
}

// Executer a entity that runs a test.
type Executer struct {
	command string            // command command name.
	args    []string          // args go test args.
	params  map[string]string // params go test params.
}

// defaultArgs default args.
var defaultArgs = map[string][]string{
	"go": {"test", "-cover", "-gcflags=all=-l"},
}

// NewExecuter creates a executer.
func NewExecuter(command string, args []string, params map[string]string) IExecuter {
	if len(args) == 0 {
		args = defaultArgs[command]
	}

	return &Executer{
		command: command,
		args:    args,
		params:  params,
	}
}

// Execute execute the command.
func (t *Executer) Execute() (outputBuffer bytes.Buffer, err error) {
	var args []string
	args = append(args, t.args...)
	for _, param := range t.params {
		args = append(t.args, param)
	}

	cmd := exec.Command(t.command, args...)
	cmd.Stdout = &outputBuffer

	err = cmd.Run()

	return
}
