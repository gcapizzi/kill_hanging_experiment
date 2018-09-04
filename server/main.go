package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// sigc := make(chan os.Signal, 1)
	// signal.Notify(sigc, syscall.SIGTERM)
	// go func() {
	// 	<-sigc
	// 	fmt.Println("SIGTERM received")
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("5 seconds passed")
	// 	os.Exit(0)
	// }()

	// outr, outw := io.Pipe()
	// defer outw.Close()
	//
	// errr, errw := io.Pipe()
	// defer errw.Close()

	var err error

	cmd := exec.Command("../idle/idle")
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Failed to create stdout pipe : %v\n", err)
	}

	// stdoutBuff := make([]byte, 256)
	// go func() {
	// 	for {
	// 		_, err := stdoutPipe.Read(stdoutBuff)
	// 		if err == io.EOF {
	// 			return
	// 		}
	// 		os.Stdout.Write(stdoutBuff)
	// 	}
	// }()
	var stdoutBuff bytes.Buffer
	mystdout := io.MultiWriter(os.Stdout, &stdoutBuff)
	go io.Copy(mystdout, stdoutPipe)

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Failed to create stderr pipe : %v\n", err)
	}

	if err = cmd.Start(); err != nil {
		panic("failed to start idle")
	}
	sigc := make(chan os.Signal, 1)

	signal.Notify(sigc, syscall.SIGTERM)
	go func() {
		<-sigc
		fmt.Println("SIGTERM received")

		if err := stdoutPipe.Close(); err != nil {
			fmt.Printf("Failed to close stdout pipe: %v\n", err)
		}
		if err := stderrPipe.Close(); err != nil {
			fmt.Printf("Failed to close stderr pipe: %v\n", err)
		}
		os.Exit(0)
	}()

	time.Sleep(30 * time.Second)

	go func() {
		cmd.Wait()
		fmt.Println("exit handler!")
		os.Exit(0)
	}()

	fmt.Println("finished")
}
