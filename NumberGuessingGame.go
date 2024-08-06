package main

//
//import (
//	"bufio"
//	"fmt"
//	"math/rand"
//	"os"
//	"time"
//)
//
//var stdinq = bufio.NewReader(os.Stdin)
//
//func InputIntValue2() (int, error) {
//	var n int
//	_, err := fmt.Scanln(&n)
//	if err != nil {
//		stdinq.ReadString('\n')
//	}
//	return n, err
//}
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	r := rand.Intn(100) //랜덤한 숫자를 범위를 지정해서 생성하고 싶으면 해당 함수를 통해서 만들어주면 됨
//	//근데 문제는 이때 생성되는 랜덤값은 완전한 랜덤이 아닌 유사 랜덤값이라고 한다
//	//랜덤값이 산출되는 초기값이 같아서 그렇다
//	//그래서 이를 해결해주기 위해서 현재 시간을 기준으로 랜덤한 값을 뽑아낼 수 있도록 현재시간을 Seed 로 매겨줌
//	cnt := 1
//	for {
//		fmt.Print("숫자값을 입력하세요: ")
//		n, err := InputIntValue2()
//		if err != nil {
//			fmt.Println("숫자를 입력해주세요")
//		}
//		if n == r {
//			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수 : ", cnt)
//			break
//		} else if n > r {
//			fmt.Println("입력하신 숫자가 더 큽니다.")
//		} else {
//			fmt.Println("입력하신 숫자가 더 작습니다.")
//		}
//		cnt++
//	}
//}
