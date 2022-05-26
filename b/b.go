package b

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("zz/b 库将 logrus 的输出设置成了 io.Discard。")
	fmt.Println("所有 logrus 函数都不会向控制台打印信息。")
	logrus.SetOutput(io.Discard)
}
