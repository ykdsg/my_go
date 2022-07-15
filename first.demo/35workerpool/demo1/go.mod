module demo1

go 1.18

require (
	com.yk/first/workpool v1.0.0
)

replace (
	com.yk/first/workpool v1.0.0 => ./../workerpool1
)