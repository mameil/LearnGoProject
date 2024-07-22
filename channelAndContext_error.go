package main

import "fmt"

func main() {
	ch := make(chan int) //크기가 0 인 채널 생성(일반적으로 생성하면 크기가 0인 채널이 생성됨)

	ch <- 9 //그래서 채널에 데이터를 넣었지만, 채널에서 데이터를 보관할 곳이 없기 때문에 고루틴에서 데이터를 빼가는 것을 기다림
	//하지만 아무도 해당 채널에서 데이터를 빼가지 않기 때문에
	//고루틴은 계속해서 채널에서 데이터를 빼가는 것을 기다림 >> 그래서 deadLock 이 발생하면서 프로그램이 강제로 종료됨
	fmt.Println("This Should NEVER PRINT!")

	//그래서 해당 이런 이슈를 해결하기 위해선 초장에 채널의 크기를 설정해주는게 좋음
	//다시 channelAndContext.go
}
