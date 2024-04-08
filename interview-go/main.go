package main

import (
	"mydemos/interview-go/question"
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
	question.GetAbsPath()
}
