package main

import "fmt"

// interface 를 사용해서 메소드 구현을 포함해서 구조화된 객체가 아닌 추상화된 객체를 만들 수 있음
// 주의 사항이 몇 가지 있으니 확인
type Sample interface {
	String() string
	//String(int) string >> 에러 발생 >> 동일한 이름의 메소드가 존재할 수는 없음
	//_(x int) >> 에러 발생 >> 메소드에 반드시 메소드명이 있어야 함
	//+ 인터페이스에서는 메소드의 구현이 불가능하다
}

// 인터페이스를 선언하고, 메소드를 명시해주고
type Stringer interface {
	String() string
}

// 상속받아서 구현할 struct 을 만들고
type Student3 struct {
	Name string
	Age  int
}

type Student3By struct {
	Name string
	Age  int
}

func (s Student3By) String() string {
	return fmt.Sprintf("Name : %s, Age : %d", s.Name, s.Age)
}

// student3을 리시버로 가지고 인터페이스에서 명시한 메소드를 구현해준다
// 자바에서 "인터페이스"를 상속받으면 "클래스에서" 상속받고, 그 메소드를 강제적으로 "구현"해줘야 하는데 >> golang 에서는 "인터페이스를 상속받고 싶은 객체를 인터페이스에 대입" 해주면 원하는 메소드를 사용할 수 있음
func (s Student3) String() string {
	return fmt.Sprintf("Name : %s, Age : %d", s.Name, s.Age)
}

////////////////////////////////////////////////////////////////////////////////////

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

type FedexSender struct{}

func (f *FedexSender) Send(parcel string) {
	fmt.Println("Fedex sent", parcel)
}

type PostSender struct{}

func (f *PostSender) Send(parcel string) {
	fmt.Println("Post sent", parcel)
}

func main() {
	student := Student3{Name: "shim", Age: 28} //implement 받고 싶은 struct 을 생성
	student2 := Student3By{Name: "kyudo", Age: 27}

	var stringer Stringer //implement 할 인터페이스를 선언

	stringer = student //인터페이스에 struct 을 넣어줘서 implement 받는 구조
	fmt.Println(stringer.String())
	stringer = student2 //만약에 interface 에 구현한 메소드를 안적어두고 이렇게 interface 에 struct 를 넣어주면 "메소드 구현해!" 라고 에러가 발생하니까 같이 만들어주면 됨
	fmt.Println(stringer.String())

	/**
	인터페이스를 사용하는 이유
	인터페이스를 이용하면 구체화된 객체가 아닌 인터페이스만을 가지고 메소드를 호출하는 것이 가능하고, 수정이 필요해도 인터페이스만 수정하면 되기 때문에 유지보수가 편리해짐
	*/

	koreanPostSender := &PostSender{}
	SendBook("어린왕자", koreanPostSender)
	SendBook("그리스인 조르바", koreanPostSender)

	fedexSender := &FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	/**
	추상화 계층이란
	내부 동작을 감춰서 서비스를 제공하는 쪽과 사용하는 쪽 모두에게 자유를 주는 방식
	위의 koreanPostSender, fedexSender 모두 인터페이스를 통한 추상화로 인해 > 사용자는 아무런 생각없이 그냥 SendBook 메소드만 사용하고 객체만 잘 넣어주면 알아서 맞게 동작하는걸 봄
	*/

	/**
	추가적인 인터페이스 기능(아래)
	- 포함된 인터페이스
	- 빈 인터페이스
	- 인터페이스 디폴트값 >> 기본값은 유효하지 않은 메모 주소를 나타내는 nil 임
	*/
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student4{15})
}

////////////////////////////////////////////////////////////////////////////////////
/**
인터페이스 속 인터페이스
구조체 안에서 구조체를 포함한 필드를 가질 수 있듯이, 인터페이스 안에서도 인터페이스를 가질 수 있음
*/
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

/**
이렇게 되어있는 상황에서
Read(), Write(), Close() 메소드를 모두 가지고 있는 타입은, Reader, Writer, ReadWriter 모두 사용이 가능
Read(), Close() 메소드를 모두 가지고 있는 타입은 Reader 만 사용 가능
Write(), Close() 메소드를 모두 가지고 있는 타입은 Writer 만 사용 가능
Read(), Write() 메소드를 모두 가지고 있는 타입은 Close() 메소드가 없기(구현안했기) 때문에 Reader, Writer 둘 다 사용할 수 없음
*/

////////////////////////////////////////////////////////////////////////////////////
/**
interface { }
위처럼 선언해주면, 메소드를 가지지 않은 빈 인터페이스라고 볼 수 있다
이걸 언제쓰냐? > 어떤 값이든 받을 수 있는 함수, 메소드, 변숫값을 만들 때 사용
*/
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("int : ", int(t))
	case float64:
		fmt.Println("float64 : ", float64(t))
	case string:
		fmt.Println("string : ", string(t))
	default:
		fmt.Println("Not supported type")
	}
}

type Student4 struct {
	Age int
}
