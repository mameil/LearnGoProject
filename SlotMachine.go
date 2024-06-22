package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

/**
 * @brief		SlotMachine
 * @details		SlotMachine 은 슬롯머신 게임을 구현한 함수이다
 * 				초기의 금액은 1000원이고
 * 				1~5 사이 숫자를 입력받고, 1~5 사이에서 랜덤으로 숫자를 돌려서
 * 				만약에 2개의 값이 같으면 >> 가진 돈에 500을 추가 | 축하한다는 메세지와 현재 보유금액 노출
 * 				만약에 2개의 값이 다르면 >> 가진 돈에 100을 차감 | 아쉽다는 메세지와 현재 보유금액 노출
 * 				종료되는 시점을 사용자의 보유금액이 0원 이하가 되거나 5000 이상이 되면 종료(걍 숫자는 알아서 조절)
 */

func InputIntValue() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		stdin.ReadString('\n')
	}
	return input, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	money := 1000

	fmt.Println("===========================================================")
	fmt.Println("슬롯머신 게임을 시작합니다")
	fmt.Printf("초기 금액은 %v 원 입니다\n", money)
	fmt.Println("보유 금액이 500원 이하가 되거나, 1500이 되면 게임이 종료됩니다")
	fmt.Println("===========================================================")

	for {
		fmt.Print("1 ~ 5 사이의 숫자를 입력하세요: ")
		var input int
		input, _ = InputIntValue()
		magicNum := rand.Intn(5) + 1

		switch {
		case input == magicNum:
			{
				money += 500
				fmt.Printf("축하합니다! 2개의 숫자가 일치합니다, 현재 잔액 %v\n", money)
			}
		case input != magicNum:
			{
				money -= 100
				fmt.Printf("아쉽습니다! 2개의 숫자가 일치하지 않습니다(랜덤숫자 : %v), 현재 잔액 %v\n", magicNum, money)
			}
		}

		if money >= 1500 {
			fmt.Println("축하합니다! 1500원을 달성하셨습니다 게임을 종료합니다")
			return
		}
		if money <= 500 {
			fmt.Println("아쉽군요 보유금액이 적어서 게임을 진행할 수 없습니다 게임을 종료합니다")
			return
		}
	}
}
