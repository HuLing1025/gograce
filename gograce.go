package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// func main() {
// 	// fileprocessor.NewSourceCodeProcessor().BuildASTTree(".")
// }

func main() {
	cmd0 := exec.Command("go", "env")
	var outputBuf1 bytes.Buffer
	cmd0.Stdout = &outputBuf1
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
		return
	}
	if err := cmd0.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf1.Bytes())

}
