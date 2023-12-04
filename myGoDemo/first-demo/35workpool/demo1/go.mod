module demo1

go 1.21

// 使用workspace 会优先使用workspace下的module 源码
//require com.yk/demo/workpool v1.0.0

//replace com.yk/demo/workpool v1.0.0 => ../workpool1
