package generator

/*
	设计思想：
		1.相当于yield功能
		2.函数返回一个只读的 <-chan
		3.在函数内部开一个goruntine并发生成值放入chan中
*/

func Count(start, end int) <-chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	return ch
}
