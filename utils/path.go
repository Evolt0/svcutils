package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// os.Getwd()获取当前目录路径

// GetExecPath 获取执行路径
func GetExecPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
