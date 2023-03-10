// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

package main

import (
	"flash_box_server/internal/flashBox"
	"os"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := flashBox.NewFlashBoxCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
