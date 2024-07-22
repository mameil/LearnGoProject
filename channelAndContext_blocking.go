package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 //데이터를 넣음
	}

	close(ch) //채널을 수기로 닫아줌 >> 이래야 main 메소드에서 무한 루프의 for 가 끝남

	wg.Wait() //여긴 실행 안됨
	//여기 부분이 문제인게, 채널은 데이터가 들어오기를 기다리고 있고, wg.Wait() 은 채널에서 종료되었다고 시그널을 줘야하는데 서로가 다른 곳을 바라보고 기다리고 있기 때문에 deadLock 이 발생

	//이걸 해결하는 방법은 >> 채널을 다 사용하고 "close()" 메소드를 통해서 채널을 닫아주면 됨 >> 채널이 닫힌걸 알면 for 문이 종료되고 다음이 수행될 수 있음
}

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //데이터를 계속해서 기다림
		//이렇게 처리하면 채널에서 데이터를 계속해서 기다리는 것이 가능
		//채널을 계속해서 바라보다가, 채널에 데이터가 인입되면 그때 n 변수에 값을 복사하고 아래 로직이 수행
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}

	wg.Done() //여긴 실행안됨
}
