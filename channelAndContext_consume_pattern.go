package main

import (
	"fmt"
	"sync"
	"time"
)

/*
*
고루틴에서 뮤텍스를 사용하지 않는 방법에 대해서 이전에 적어봤는데
애초에 같은 데이터를 바라보는 상황이 나와서 뮤텍스의 Lock 을 피한다는 내용이였다
그리고 그 방법으로는 영역을 나누는 방법, 그리고 역할을 나누는 방법이라고 했었는데
이번에 그 "역할을 나누는 방법" 에 대해서 살펴보자

예를 들어서
자동차 공장에서
차체 생산 > 바퀴 설치 > 도색 > 완성
위와 같은 단계를 거쳐서 생산한다고 가정하자

각 공장에 1초가 걸린다고 보면 자동차 한 대를 만드는데 3초가 걸려야 "완성"까지 간다
근데 3명이 공정 하나씩 처리하면, 첫 차를 생산하늗네 3초가 걸리고, 그 뒤론 1초마다 하나씩 처리되면서 툭툭 튀어나옴
그림으로 보면 이해가 더 잘됨
A    B    C
T1   T1   T1
T2   T2   T2
T3   T3   T3

1초대 : "자동차"라는 결과물이 아직은 없음
A 공정에서는 T1 작업 진행
B 공정에서는 받은게 없으니 대기
C 공정에서는 받은게 없으니 대기

2초대 : "자동차"라는 결과물이 아직은 없음
A 공정에서는 T1 을 B 공정으로 넘기고 / T2 작업을 진행
B 공정에서는 T1 작업을 A 공정으로부터 받아서 진행
C 공정에서는 아직 받은게 없으니 대기

3초대 : "자동차"라는 결과물이 아직은 없음
A 공정에서는 T2 을 B 공정으로 넘기고 / T3 작업을 진행
B 공정에서는 T1 을 C 공정으로 넘기고 / T2 작업을 A 공정으로부터 받아서 진행
C 공정에서는 T1 작업을 B 공정으로부터 받아서 진행

4초대 : T1 가 튀어나옴

5초대 : T2 가 튀어나옴

6초대 : T3 가 튀어나옴

>>>>길고 길었지만 결과적으로 처음에 3초가 걸리고, 이후에는 1초마다 하나씩 자동차가 생산되는걸 볼 수 있음
이런 것을 보고 "컨베이어 벨트 시스템"이라고 함 > "공정"들 사이를 컨베이어 벨트로 이어서 생산하는 방식
여기서 컨베이어 벨트를 "채널"이라고 생각하면 된다
*/

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wgc sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wgc.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wgc.Wait()
	fmt.Println("End Factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			//차체를 쭉 만들고
			car := &Car{}
			car.Body = "Sports Car"
			//타이어 채널로 던진다
			tireCh <- car
		case <-after:
			//쭉 다음 벨트로 넘기고 타이어 벨트를 닫아줌(더 이상 넣을 게 없으니)
			close(tireCh)
			wgc.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	//타이어 채널에서 받은 자동차
	for car := range tireCh {
		time.Sleep(time.Second)
		//타이어를 장착
		car.Tire = "Winter Tire"
		//페인트 채널로 넘긴다
		paintCh <- car
	}
	//주루룩 벨트에 넘어온 자동차를 작업한 뒤
	wgc.Done()
	close(paintCh) //채널을 닫아준다
}

func PaintCar(paintCh chan *Car) {
	//페인트 채널에서 받은 자동차
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Blue"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f 차 생성 완료 : CarBody - %s / CarTire - %s / CarColor - %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wgc.Done()
}

/**
이렇게 채널을 사용해서 역할을 나누어서 작업하면
굳이 락처리를 하면서 서로 기다리면서 작업하는 것보다 더 빠르게 작업하는 것이 가능하고, 굳이 동시에 접근할 필요가 없으니 뮤텍스의 락 작업도 필요 없음

이렇게 채널을 사용해서 한쪽에서 데이터를 생성해서 넣어주고, 그걸 받아서 처리하고 다음으로 넘기고, 이러한 방식을
"생산자 소비자 패턴"이라고 함
위 예시에서는 MakeBody() 가 생산자, InstallTire() 가 소비자이고 / InstallTire() 가 생산자이고, PaintCar() 가 소비자이다
*/
