package main

import (
	"sync"

	"github.com/gogf/gf/frame/g"
)

// 并行3个协程测试

type checkResult struct {
	Code   int
	Msg    string
}

func main() {
	ch := make(chan *checkResult, 5)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go do1("协程1", ch, wg)
	go do2("协程2", ch, wg)
	go do3("协程3", ch, wg)
	wg.Wait()
	close(ch)

	for v := range ch {
		g.Dump(v.Msg,v.Code)
	}

}

func do1(s string, ch chan<- *checkResult, wg *sync.WaitGroup) {
	result := &checkResult{}
	defer wg.Done()
	result.Code = 1
	result.Msg = s + " success do1 "
	ch <- result
	return
}

func do2(s string, ch chan<- *checkResult, wg *sync.WaitGroup) {
	result := &checkResult{}
	defer wg.Done()
	result.Code = 2
	result.Msg = s + " success do2 "
	ch <- result
	return
}

func do3(s string, ch chan<- *checkResult, wg *sync.WaitGroup) {
	result := &checkResult{}
	defer wg.Done()
	result.Code = 3
	result.Msg = s + " success do3 "
	ch <- result
	return
}

// 执行顺序不一定
// 协程3 success do3 3

// 协程1 success do1 1

// 协程2 success do2 2
