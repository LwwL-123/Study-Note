package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {

	ctx,cancel := context.WithCancel(context.Background())
	go find(ctx,cancel)
	go can(ctx,cancel)

	for  {
		
	}
}

func find(ctx context.Context,cancelFunc context.CancelFunc)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("任务取消")
			return
		default:
			fmt.Println("runing")
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func can(ctx context.Context,cancelFunc context.CancelFunc) {
	time.Sleep(time.Second)
	cancelFunc()
}
