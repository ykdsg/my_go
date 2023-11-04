# 基准测试
命令行</br>
go test -bench=. -benchmem

也可以通过ide直接运行。

# 结合pprof分析
cpu：</br>
go test -bench=. -benchmem -cpuprofile=cpu.out

mem:</br>
go test -bench=. -benchmem -memprofile=mem.out

查看具体情况：</br>
go tool pprof [file]

然后在pprof中通过lsit *** 查看消耗情况。