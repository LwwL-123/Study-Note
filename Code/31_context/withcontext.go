
package main

import (
	"context"
	"fmt"
	"time"
)

var c = 1

func speakMemo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Done")
			return
		default:
			fmt.Printf("exec default func:%d\n",c)
			c++
		}
	}
}

func timeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("进程结束")
			return
		default:
			time.Sleep(time.Millisecond)
			fmt.Printf("我是子进程")
		}

	}
}

func main() {
	rootContext := context.Background()

	//ctx, cancelFunc := context.WithCancel(rootContext)
	//go speakMemo(ctx)
	//cancelFunc()

	ctx,_ := context.WithTimeout(rootContext,time.Second)
	go timeout(ctx)
	time.Sleep(time.Second * 2)
}
