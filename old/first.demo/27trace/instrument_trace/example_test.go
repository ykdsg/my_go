package trace_test

import "hz.com/yk/instrument_trace"

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()()
}

//ExampleXXX 形式的函数表示一个示例，go test 命 令会扫描 example_test.go 中的以 Example 为前缀的函数并执行这些函数。
func ExampleTrace() {
	a()
	// Output:
	//g[00001]:    ->hz.com/yk/instrument_trace_test.a
	//g[00001]:        ->hz.com/yk/instrument_trace_test.b
	//g[00001]:            ->hz.com/yk/instrument_trace_test.c
	//g[00001]:                ->hz.com/yk/instrument_trace_test.d
	//g[00001]:                <-hz.com/yk/instrument_trace_test.d
	//g[00001]:            <-hz.com/yk/instrument_trace_test.c
	//g[00001]:        <-hz.com/yk/instrument_trace_test.b
	//g[00001]:    <-hz.com/yk/instrument_trace_test.a
}
