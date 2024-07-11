package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

/*
*
에러 핸들링이란, 프로그램에서 에러를 처리하는 방법
어디에서나 발생할 수 있으며 이걸 적절하게 처리하는 것도 프로그램을 만드는 사람의 몫
*/
func main() {

	/*
		에러를 처리하는 가장 기본적인 방식은 에러를 리턴하고 잘 처리하는 방법임
		특정 로직을 수행하다가 에러가 발생해서 프로그램이 그냥 픽 죽어버리는 것보다는 에러가 발생하면 해당 에러를 노출해주고 다른 액션을 취해서 프로그램이 이어지도록 유지하는 것이
		사용자 입장에서 더 좋은 경험을 제공하는 것임
	*/

	//파일의 내용을 읽어서 노출하는 로직인데, 파일을 읽는 과정에서 에러 발생 시, 파일을 생성해서 읽어주는 것까지 수행하도록 고려한 로직
	//ReadFile(), WriteFile() 참고
	fileName := "ErrorHandlingTest.txt"
	line, err := ReadFile(fileName)
	if err != nil {
		err = WriteFile(fileName, "Hello World!")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다 - ", err)
			return
		}
		line, err = ReadFile(fileName)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다 - ", err)
			return
		}
		fmt.Println("파일 내용 - ", line)
	}

	//직접 사용자 에러를 정의해서 리턴하는 방식으로 구현도 가능하다
	//Sqrt(), New() 참고
	sqrt, err := Sqrt(-2)
	if err != nil {
		//fmt.Println("[ERROR]", err)
		//아래 함수에서처럼 fmt.Errorf() 을 통해서 에러를 정의해서 날리는 방법도 있고
		//아래서 재정의한 New() 을 통해서 에러를 전달하는 것도 가능
		errors.New("이것도 에러를 발생시키는 방법 중 하나!")
	} else {
		fmt.Println("Sqrt -2 => ", sqrt)
	}

	//에러 타입
	//에러는 interface 로
	/*
		type error interface {
			Error() string
		}
	*/
	//이렇게 생겼음, 그렇기 떄문에 어떤 타입이든 문자열을 리턴하는 Error() 를 포함한다면, 에러로 사용이 가능하다
	//사용 예시 >> PasswordError 쪽 확인
	err = RegisterAccount("myId", "myPw") //ID, PW 입력
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { //인터페이스 변환
			fmt.Printf("%v Len: %d RequiredLen: %d\n", errInfo, errInfo.Len, errInfo.RequiredLen)
		}
	} else {
		fmt.Println("회원 가입에 성공했습니다.")
	}
}

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close() //해당 함수가 종료되기 전에 무조건적으로 파일을 닫는다

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(fileName string, line string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close() // 해당 함수가 종료되기 전에 무조건적으로 파일을 닫는다

	_, err = fmt.Fprintln(file, line)
	return err
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//fmt.Errorf() 을 통해서 에러를 특별하게 만들어서 반환하는 것도 가능
		return 0, fmt.Errorf("루트로 나누는 건 양수만 가능합니다, 입력된 값 : %v", f)
	}
	return math.Sqrt(f), nil
}

// errors 패키지의 New() 함수를 통해서 에러를 생성하는 것도 가능하다
func New(text string) error {
	return errors.New(text)
}

type PasswordError struct {
	Len         int
	RequiredLen int
}

func (err PasswordError) Error() string {
	return "길이의 암호가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 0 {
		return PasswordError{len(password), 8}
	}
	return nil
}
