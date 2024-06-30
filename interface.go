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

func main() {
	student := Student3{Name: "shim", Age: 28} //implement 받고 싶은 struct 을 생성
	student2 := Student3By{Name: "kyudo", Age: 27}

	var stringer Stringer //implement 할 인터페이스를 선언

	stringer = student //인터페이스에 struct 을 넣어줘서 implement 받는 구조
	fmt.Println(stringer.String())
	stringer = student2 //만약에 interface 에 구현한 메소드를 안적어두고 이렇게 interface 에 struct 를 넣어주면 "메소드 구현해!" 라고 에러가 발생하니까 같이 만들어주면 됨
	fmt.Println(stringer.String())

}
