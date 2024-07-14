package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
			//golang 에서는 기본적으로 인터페이스를 통하면 어떤 타입이든 동일한 동작을 수행하는 것이 가능한데
			//여기서의 err.(PasswordError) 로직은 "타입 단언"이라는 기능을 사용한 케이스임 >이건 err 변수가 PasswordError 인지 체크하는 로직임
			//만약에 err 가 PasswordError 타입이면, ok 는 true 값이 되면서 errInfo 는 PasswordError 타입으로 설정됨
			//만약에 err 가 PasswordError 타입이 아니면, ok 는 false 가 되면서 errInfo 는 nil 으로 설정됨
			fmt.Printf("%v Len: %d RequiredLen: %d\n", errInfo, errInfo.Len, errInfo.RequiredLen)
		}
	} else {
		fmt.Println("회원 가입에 성공했습니다.")
	}

	//에러 래핑
	//에러를 감싸서 새로운 에러를 만들어야하는 경우도 존재함
	//예시) 파일에서 텍스트를 읽어서 특정 타입의 데이터로 변환하는 경우 파일 읽기에서 발생하는 에러도 필요하지만, 텍스트의 몇 번쨰 줄의 몇 번째 칸에서 에러가 발생하는지 알아야함...
	//MultipleFromString(), readNextInt(), readEq()
	readEq("123 3")   // 여기까진 성공
	readEq("123 abc") // 123까진 성공
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

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) //스캐너 생성 >> NewScanner() 는 io.Reader 인터페이스를 인수로 받기 때문에 string 타입을 io.Reader 로 만들어주려고 strings.NewReader() 를 사용
	scanner.Split(bufio.ScanWords)                      //한 단어씩 끊어서 읽기  >> Split 을 통해서 어떻게 끊어서 읽을지를 알려주고, bufio.ScanWords 을 통해 단어를 기준으로 읽게되고 bufio.ScanLines 을 사용하면 한 줄씩 끊어서 읽게됨

	//2번 잘라서 읽는다
	pos := 0
	a, n, err := readNextInt(scanner)
	//2번을 각각 수행하면서 에러 있으면 그때그때 뱉음
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%v err:%v", pos, err)
		//에러 감싸기
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	//2번을 각각 수행하면서 에러 있으면 그때그때 뱉음
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%v", pos, err)
	}
	return a + b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 리턴한다
// 변환된 숫자, 읽은 글자 수, 에러를 리턴한다
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { //단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word) //문자열을 숫자로 변환
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word: %v err: %v", word, err) //에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		//감싸진 에러를 꺼내오는 방법
		//As() 메소드를 통해서 err 안에 감싸진 에러 중 두 번째 인수의 타입인 *strconv.NumError 로 변환될 수 있는 에러가 있다면 변환하여 값을 넣고 true 를 반환
		//추가로 Is() 메소드를 통해서 단순하게 객체 타입만을 확인하는 것도 가능하니 As 로 아에 타입 캐스팅을 처리해서 확인하는 방법도 있고, Is() 을 통해서 객체 타입 체크만해서 확인하는 방법도 있다 이정도
		if errors.As(err, &numError) { //감싸진 에러가 NumError 인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}
