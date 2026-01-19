// gen.go
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 跨平台生成 proto 文件
func main() {
	//在控制台输出文字
	os.Stdout.Write([]byte("✅ apiLink: 正在生成...\n"))

	moduleName := "github.com/wegge0857/PolyrepoGoMicros-ApiLink"
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".proto") || strings.Contains(path, "third_party") {
			return nil
		}
		cmd := exec.Command("protoc",
			"--proto_path=.",
			"--proto_path=./third_party",
			"--go_out=.", "--go_opt=module="+moduleName,
			"--go-grpc_out=.", "--go-grpc_opt=module="+moduleName,
			"--go-http_out=.", "--go-http_opt=module="+moduleName, // 必须加这一行
			path,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	})
	//在控制台输出文字
	os.Stdout.Write([]byte("✅ apiLink: 所有proto文件已生成\n"))
}
