package main

import (
	"fmt"
	_ "zz/b"

	_ "github.com/silenceper/wechat/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("导入了 wechat 包时的行为")
	logrus.Println("因为 zz/b 库的 init 方法。这句 log 应该不会显示在控制台上。")
	fmt.Println("aa/a 刚刚使用 logrus 向控制台打印了一条信息。")
}
