module demo1

go 1.18

require (
	com.yk/first/workpool v1.0.0
)

//利用 replace 指示符将 demo1 对 workerpool 的引用指向本地 workerpool1 路径
replace (
	com.yk/first/workpool v1.0.0 => ./../workerpool1
)