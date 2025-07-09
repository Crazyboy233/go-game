package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

// 这里应该返回一个切片，返回解析出来的命令和操作参数
func Parser() Command {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
		fmt.Printf("input = %s\n", input)
	}

	// 检查读取错误
	if err := scanner.Err(); err != nil {
		// 如果有错误发生，将错误信息输出到标准错误输出
		fmt.Fprintln(os.Stderr, "读取错误：", err)
	}

	parts := strings.Fields(input)

	// 读取命令个数校验
	if len(parts) == 0 {
		return Command{}
	}

	cmdName := parts[0]
	args := parts[1:]
	command := &Command{cmdName, args}

	return *command
}

// 参数个数校验
// func ParameterVerification(name string) int {
// 	switch name {
// 	case "/Login":
// 		return 1
// 	case "/Logout":
// 		return 1
// 	case "register":
// 		return 1
// 	default:
// 		return -1
// 	}
// }
