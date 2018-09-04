package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	buf := bytes.NewBuffer(nil)

	cmd := exec.Command("../server/server")
	cmd.Stdout = buf
	cmd.Stderr = buf

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Fprintln(os.Stderr, "failed to wait", err)
		}
	}

	fmt.Println(buf.String())
	fmt.Println("done")
}
