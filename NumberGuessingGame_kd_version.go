package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	//숫자 맞추기 게임을 만들어본다
	//프로그램 실행 시, 랜덤한 숫자를 하나 선정하고
	//사용자가 입력한 숫자가 랜덤한 숫자와 일치하는지 계속해서 돌아가며 확인하는 프로그램

	//0 ~ 99 사이의 숫자를 생성하기 위해서 100을 나눈 나머지를 사용
	targetNum := rand.Int() % 100

	var userInput int
	var userCount int
	fmt.Println("===========================================================")
	fmt.Println("숫자 맞추기 게임을 시작합니다")
	fmt.Println("숫자는 애플리케이션에서 랜덤으로 선정하며 숫자는 0 ~ 99사이로 설정됩니다")
	fmt.Println("===========================================================")
	var stdin = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("숫자값을 입력하세요: ")
		_, err := fmt.Scan(&userInput)

		if err != nil {
			fmt.Println("숫자를 입력해주셔야 합니다")
			stdin.ReadString('\n')
			continue
		}

		switch {
		case userInput == 9999:
			{
				fmt.Println("===========================================================")
				fmt.Println("Admin 모드 >>>> 랜덤으로 선정된 숫자는 아래와  같습니다")
				fmt.Println(targetNum)
				fmt.Println("===========================================================")
			}
		case targetNum == userInput:
			{
				userCount++
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수 : %v", userCount)
				return
			}

		case targetNum > userInput:
			{
				userCount++
				fmt.Println("입력하신 숫자가 더 작습니다.")
				continue
			}

		case targetNum < userInput:
			{
				userCount++
				fmt.Println("입력하신 숫자가 더 큽니다.")
				continue
			}
		}
	}
}
