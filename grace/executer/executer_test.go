package executer

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	resultBuf, err := NewExecuter("go", []string{}, map[string]string{
		"path": `./examples`,
	}).Execute()

	fmt.Println(resultBuf.String())

	if err != nil {
		t.Error(err)
	}
}
