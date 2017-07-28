// 这个示例程序展示如何使用
// 有缓冲的通道和固定数目的
// goroutine 来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的goroutine 的数量
	taskLoad         = 50 // 要处理的工作的数量
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// init 初始化包，Go 语言运行时会在其他代码执行之前
// 优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// main 是所有Go 程序的入口
func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	fmt.Println("sleep")
	time.Sleep(time.Duration(10000) * time.Millisecond)
	fmt.Println("wakeup")
	wg.Wait()
	close(tasks)
	for {
		task, ok := <-tasks
		if !ok {
			break
		}
		fmt.Println(task)
	}
}

// worker 作为goroutine 启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	fmt.Println("worker:", worker)
	// 通知函数已经返回
	defer wg.Done()

	for i := 1; i < 10; i++ {
		tasks <- fmt.Sprintf("This is from worker:%d_%d", worker, i)
	}
}
