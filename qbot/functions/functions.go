package functions

import "os"

func FunctionList() string {
	b, _ := os.ReadFile("/root/software/qqbot/functionList.txt")
	return string(b)
}
