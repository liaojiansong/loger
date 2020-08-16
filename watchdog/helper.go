package watchdog

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func execCommand(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)
	// 创建一个标准输入给人家
	cmd.Stdin = strings.NewReader("标准输入")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", errors.New("读取失败" + err.Error())
	}
	return out.String(), nil
}