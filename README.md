因为当一个工程依赖了多个库时，这些被依赖的库的 init 方法是按顺序执行的。
所以对于在 wechat 库之前 init 的库，它们对 logrus 的设置会被 wechat 库的 init 方法覆盖掉。

使用 make run 命令复现这个 bug。
若未安装 make 则可以手动执行 Makefile 中记录的命令。手动执行命令时需要注意路径问题。