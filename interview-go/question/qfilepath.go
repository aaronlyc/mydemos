package question

import (
	"fmt"
	"os"
)

func GetAbsPath() {
	executable, err := os.Executable()
	if err != nil {
		return
	}
	fmt.Printf("exec path is: %s\n", executable)
}
