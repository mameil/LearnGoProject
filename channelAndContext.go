package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	//go 언어에서 Channel 과 Context 는 동시성 프로그래밍을 도와주는 기능임
	//채널은 고루틴 사이에서 메세지를 전달하는 메세지 큐임
	//이걸 사용하면 뮤텍스 없이 동시성 프로그래밍이 가능
	//컨텍스트는 고루틴에 작업을 요청할 때 작업 취소나 작업 시간 등을 설정할 수 있는 "작업 명세서" 역할을 수행
	//채널과 컨텍스트를 기반으로 특정 데이터를 고루틴 사이에서 전달하거나 특정 시간 동안만 작업을 요청하거나 작업 도중에 작업 취소를 날리거나 할 수 있음

	//채널 사용하기
	//위에서 이야기했지만 채널은 고루틴 사이에서 메세지를 전달할 수 있는 메세지 큐
	//위에서 메세지 큐라고 했다시피, 큐 형식으로 FIFO 형식으로 처리됨

	//채널 만들기 >> 채널 인스턴스를 생성하는데, make() 함수를 통해서 생성한다
	//messages >> 채널 인스턴스명, chan string >> string 채널 타입, make(chan string) >> 채널 안에는 string 타입의 메세지가 들어감
	var messages chan string = make(chan string)

	//채널에 메세지 넣기
	//메세지를 넣을 때는 "<-" 연산자를 사용해서 message 라는 채널 인스턴스에 뒤에 오는 데이터를 넣어주겠다는 의미
	messages <- "This is a First Message"

	//채널에서 메세지 빼기
	var msg string = <-messages
	fmt.Println("msg :", msg)

	//채널을 사용해서 데이터를 넣고 빼기
	//square()
	var wg sync.WaitGroup
	ch := make(chan int) //채널을 int 타입이 들어가게끔 만들고

	wg.Add(1)          //고루틴을 넣을 작업 준비해주고
	go square(&wg, ch) //고루틴을 돌리고
	ch <- 9            //채널에 데이터를 넣어준다
	wg.Wait()

	//채널의 크기
	//일반적으로 채널을 처음 만들면, 크기가 0이다 >> 크기가 0이라는 것의 의미는 채널에 들어온 데이터를 담을 공간이 없다는 것
	//채널의 크기란 택배기사가 택배를 전달하는데, 택배를 담을 곳이 없다면 수신자가 와서 가져갈 때까지 택배를 들고 기다려야함
	//그냥 택배기사가 보관함에 던지고 가면 되는걸 보관함이 없으면 기다려야하는 사고가 생김 >> 해당 에러 케이스는 channelAndContext_error.go 에서 확인

	//내부에 데이터를 보관할 수 있는 메모리 영역을 "버퍼"라고 하는데, 그 "버퍼" 의 사이즈를 채널을 만들 때 지정해주는 것이 가능하다
	//채널을 선언해줄 때 버퍼의 크기를 초장부터 선언해주는 케이스
	myChannel := make(chan string, 2)
	myChannel <- "Hello"
	myChannel <- "World"
	fmt.Println(<-myChannel) //Hello
	fmt.Println(<-myChannel) //World
	// +추가로 버퍼의 사이즈가 다 차면, 빈자리가 생길 때 까지 기다리기 떄문에 또 deadlock 이 발생함
	//<-myChannel 2개를 주석처리해보고 돌리면 fatal error: all goroutines are asleep - deadlock! 에러가 발생함

	//채널에서 데이터 대기
	//channelAndContext_blocking.go 에서 확인
	//채널을 다 사용하고나면 정상적으로 close 해줘야 메모리 누수에서부터 해방될 수 있다 >> 기껏 고루틴으로 작업하는거면 제대로 쓰자

	//select 문
	//channelAndContext_select.go 에서 확인

	//일정 간격으로 수행하는 로직
	//channelAndContext_term.go 에서 확인

	//채널로 생산자 소비자 패턴으로 구현하기
	//channelAndContext_consume_pattern.go 에서 확인

	//컨텍스트 사용하기
	/**
	context 은 자체적인 패키지에서 제공해주는 기능으로, 작업을 지시할 때 작업 가능 시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할을 함
	새로운 고루틴으로 작업을 시작할 때 "일정 시간 동안만 작업을 지시하거나 외부에서 작업을 취소할 때 사용" 또는 작업 설정에 관한 데이터를 전달할 수도 있음
	*/

	//작업 취소가 가능한 컨텍스트 >> 이걸 만들어서 작업자에게 전달하면 작업을 지시한 지시자가 원할 때 작업 취소를 알릴 수 있음
	//PrintEverySecond() 확인
	wg2.Add(1)

	//WithCancel 함수는 상위 컨텍스트를 인수로 넣어주면, 그 컨텍스트를 감싼 컨텍스트를 생성
	//만약에 상위 컨텍스트가 없다면, 기본적으인, context.Background() 을 넣어줘야 함
	//context.WithCancel() 함수는 값을 2개를 반환해주는데, 첫번째 리턴값(ctx) 은 컨텍스트 객체, 두번째 리턴값(cancel)은 취소 함수로, 이걸 사용해서 원할때 작업 내용 취소 가능
	ctx, cancel := context.WithCancel(context.Background()) //create context

	go PrintEverySecond(ctx) //초마다 계속해서 Tick 을 부르는 로직이 수행되다가
	//5초 이후에
	time.Sleep(5 * time.Second)
	//취소를 날려서 작업을 취소
	cancel()

	wg2.Wait()

	//작업 시간을 설정한 컨텍스트
	//취소랑 비슷하게 사용하면 되는데, 단순이 WithCancel() 대신에 WithTimeout() 을 사용하고, 함수의 첫번째 인자는 동일하게 context 넣어주면되구 두번째 인자로 설정하고 싶은 시간을 넣어주면 됨
	//ctx, cancel:= context.WithTimeout(context.Background(), 3*time.Second)

	//특정 값을 설정한 컨텍스트
	//작업을 지시하는데 있어서, 특정 값에 따라서 분기를 쳐서 수행시키고 싶은 경우가 있을텐데
	//이때 컨텍스트에 특정 키를 기반으로 값을 읽어오게 설정하는 것이 가능한데, 이 기능을 사용한다
	//그게 바로 context.WithValue() 임
	//square27 check
	ctx2 := context.WithValue(context.Background(), "number", 9) //number 라는 키의 값을 9로 설정한 컨텍스트를 만들어준다
	go square27(ctx2)                                            //위에서 만든 키-값 기반의 컨텍스트를 넘겨서 사용한다
	wg3.Wait()

	//굉장히 어지러웠지만, 채널과 컨텍스트를 기반으로 동시성프로그래밍을 하는 방법에 대해서 알아봤는데
	//Go 언어에서는채널과 컨텍스트를 사용해서 동시성 프로그래밍을 하는데 활발하게 사용된다는 점
	//구글에 golang concurrency patterns 검색하면 잘 나온다넹
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg2.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

func square27(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int) //context.Context.Value() 메소드의 기본적인 리턴타입은 interface{} 이기 때문에 꼭 타입 변환작업을 해주고 사용해야함
		fmt.Printf("Square: %d\n", n*n)
	}
	wg3.Done()
}
