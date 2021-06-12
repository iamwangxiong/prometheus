// go语言运行shell命令
package ping

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/gogap/logrus"
)

func Ping() {
	var err error
	var cmd *exec.Cmd

	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("ping www.baidu.com")
	if _, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		log.Info("Succeed to exec cmd ping")
	}
}
