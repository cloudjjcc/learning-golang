package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 5个哲学家做到一个桌子上就餐，桌上共有5只筷子，
//哲学只能拿左手边的筷子和右手边的筷子，且只有拿到两只筷子才能就餐

func main() {
	chopsticks := [5]sync.Mutex{}
	act := func(ctx context.Context, i int) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			left := i
			right := (i + 1) % 5
			if left < right {
				chopsticks[left].Lock()
				fmt.Printf("people %d pick up left chopstick %d \n", i, left)
				chopsticks[right].Lock()
				fmt.Printf("people %d pick up right chopstick %d \n", i, right)
			} else {
				chopsticks[right].Lock()
				fmt.Printf("people %d pick up right chopstick %d\n", i, right)
				chopsticks[left].Lock()
				fmt.Printf("people %d pick up left chopstick %d\n", i, left)
			}

			fmt.Printf("people %d eat \n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("people %d put down chopstick\n", i)
			chopsticks[left].Unlock()
			chopsticks[right].Unlock()
		}

	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for i := 0; i < 5; i++ {
		go act(ctx, i)
	}

	<-ctx.Done()
}
