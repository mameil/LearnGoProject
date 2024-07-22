package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//select 문은
	//여러 채널을 동시에 기다리는 것이 가능하게 해주는 구문이다
	//select 문 안에 선언한 어떤 채널이라도 하나의 채널에서 데이터를 읽어오면, 해당 구문을 수행하고 select 문을 종료
	//추가로 계속해서 select 처리를 해주고 싶다면, for 문을 통해서 처리해주면 됨

	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square3(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Add()
}

func square3(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		//ch, quit 둘 다 기다리면서 select 문을 수행
		select {
		case n := <-ch:
			fmt.Printf("Square : %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}
