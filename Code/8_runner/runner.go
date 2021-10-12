package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//创建一个无缓冲的通道
	baton := make(chan int)

	wg.Add(1)

	//第一位持有接力棒
	go Runner(baton)

	//开始
	baton <- 1

	//等待比赛结束
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("第 %d 位开始跑步\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("第%d位准备起跑", newRunner)
		go Runner(baton)
	}

	time.Sleep(time.Second)

	if runner == 4 {
		fmt.Printf("第 %d 位跑完了，赢了比赛！\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("第%d位把接力棒递给了第%d位\n", runner, newRunner)

	baton <- newRunner

}
