package main

import (
	"context"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// 二者都可以触发
// 	// cmd := exec.CommandContext(ctx, "bash", "/Users/aaron/sleep.sh")
// 	cmd := exec.CommandContext(ctx, "bash", "-c", "echo hello && sleep 1200")
// 	out, err := cmd.CombinedOutput()
// 	fmt.Printf("ctx.Err : [%v]\n", ctx.Err())
// 	fmt.Printf("error   : [%v]\n", err)
// 	fmt.Printf("out     : [%s]\n", string(out))
// }

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "/Users/aaron/sleep.sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	go func() {
		select {
		case <-ctx.Done():
			err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			if err != nil {
				fmt.Printf("kill error   : [%v]\n", err)
			}
		}
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("output:", string(output))
	fmt.Printf("ctx.Err : [%v]\n", ctx.Err())
	fmt.Printf("error   : [%v]\n", err)
}
