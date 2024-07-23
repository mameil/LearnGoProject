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
	go square4(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}

func square4(wg *sync.WaitGroup, ch chan int) {
	//time.Tick() 메소드는 일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환하는 함수
	//이 함수가 반환한 채널에서 데이터를 읽어오면 일정 시간 간격으로 현재 시각을 나타내는 Time 객체를 리턴
	tick := time.Tick(time.Second) //1초 간격으로 시그널

	//time.After() 메소드는 현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환하는 함수
	//이 함수가 반환한 채널에서 데이터를 읽으면 일정 시간 경과 후에 현재 시각을 나타내는 Time 객체를 리턴
	terminate := time.After(10 * time.Second) //10초 간격으로 시그널

	for {
		//select 문을 사용해서 tick, terminate, ch 순서로 채널에서 데이터 읽기를 시도
		select {
		case <-tick:
			fmt.Println("Tick") //tick 에서 메세지를 받으면 Tick 출력
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return //terminate 에서 메세지를 받으면 바로 종료
		case n := <-ch:
			fmt.Printf("Square : %d\n", n*n) //tick 이랑 terminate 둘다로부터 메세지를 받아오지 못하면 ch 에서 데이터를 읽어옴
			time.Sleep(time.Second)
		}
	}
}
