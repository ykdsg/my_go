package main

func time(x, y int) int {
	return x * y
}

//有些场景存在一些高频使用的乘数，这个时候我们就没必要每次都传入这样的高频 乘数了。
//利用闭包可以

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	timesFunc := partialTimes(60)
	result := timesFunc(5)
	println(result)

}
