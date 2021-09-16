package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	court := make(chan int)

	//等待两个线程
	wg.Add(2)
	go player("1号", court)
	go player("2号", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			//如果通道被关闭，则我们胜利
			fmt.Printf("玩家%s胜利\n", name)
			return
		}

		//选随机数
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("玩家%s输了\n", name)
			close(court)
			return
		}
		fmt.Printf("玩家%s击回了第%d球\n", name, ball)
		ball++

		time.Sleep(time.Second)
		court <- ball
	}
}
