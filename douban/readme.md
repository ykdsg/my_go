douban网页解析版本，所以数据获取有点慢，这个跟豆瓣网页打开速度有关。

使用go 语言开发，比python 版本好的一点是直接一个可运行程序搞定，不再依赖本地是否安装了合适的Python版本，以及相应的lxml依赖。
当然缺点是包略微有点大5M+，因为自带了go运行时环境，不过也还可以接受吧。

通过以下命令能降低打包大小：
go build -ldflags '-w -s'

打包不同平台和架构下的包
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-w -s' -o douban-workflow-mac-amd64-1.2
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags '-w -s' -o douban-workflow-mac-arm64-1.2

