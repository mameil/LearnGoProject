package main

import "fmt"

/*
*
함수는 func 키워드와 메소드 명 사이에 소괄호로 리시버라는걸 정의해준다
"리시버"란 해당 메소드가 어떤 구조체에 속하는지에 대해서 명시해주는 역할을 한다 >> 자바나 코틀린에서는 class 라는 개념이 있지만, golang 에서는 클래스 개념이 없어서 이걸 통해서 명시해준다고 생각
*/
type account struct {
	balance int
}

// 기존에 사용했던, 리시버가 없어서 해당 객체를 적용하고자 하는 객체를 함수의 파라미터로 넘겨서 작업하는 방법
// 전에 작성했던 이야기인데, 함수의 파라미터로 객체를 넘긴다는 것은 그만큼 메모릐를 사용한다는 의미이기 때문에 좋지 않은 방법이라고 보임
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 이게 "리시버"를 사용한 메소드 생성방법
// 결국은 account 이라는 struct 에 대한 메소드이고, int 값을 받으면 그 값을 balance 에 더해주는 메소드
func (a *account) plusMethod(amount int) {
	a.balance += amount
}

//당연한거지만, 리시버를 직접 생성한 struct 을 지정할 수도 있지만, 기본적으로 제공해주는 int 나 string 같은 타입을 리시버로 지정하는 것이 가능하다
//근데 아래처럼 하면 컴파일 에러가 발생하니
//func (a int) plusMethod(amount int) int {

// 상위 struct 를 만들고 처리해주면된다 >> 과연 자주 사용할지는 모르겠지만 아무튼 존재
type myInt int

func (a myInt) plusMethod(amount int) int {
	return int(a) + amount
}

/**
메소드가 왜 필요하냐?
메소드의 소속을 정해서 객체 지향 언어에 맞게 객체에 맞는 메소드를 가질 수 있도록 설정해주는 과정
객체 지향 언어에서 말하는 좋은 방향의 프로그래밍이란 결국 "결합도"를 낮추고 "응집도"를 높이는 것을 말하는데, 리시버라는 개념을 통해서 "특정 struct"에 맞는 메소드들을 묶음으로써 "응집도"를 높힘
*/

/**
리시버에 "포인터" / "갑 타입" 각각을 넣었을 때의 차이를 확인해보자
*/
//포인터를 리시버로 넣었을 때
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입을 리시버로 넣었을 때
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 값 타입을 리시버로 받고 그 값을 반환할 때
func (a3 account) withdrawValue2(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	a := &account{1000}

	withdrawFunc(a, 1000)  //리시버가 없으니 그냥 바로 함수를 사용할 수 있음
	fmt.Println(a.balance) //0

	a.plusMethod(1000)     //리시버가 있는 함수는 클래스 내부의 함수처럼 객체.메소드명 요렇게 사용
	fmt.Println(a.balance) //1000

	//리시버의 값타입 vs 포인터타입
	var mainA *account = &account{1000}

	//포인터를 타입으로 가진 메소드를 호출하면, 포인터가 가리키고 있는 메모리의 주소값이 복사
	mainA.withdrawPointer(500)
	fmt.Println(mainA.balance) //500

	//값을 타입으로 가진 메소드를 호출하면, 리시버 타입의 모든 값이 복사됨 >> 결국 파라미터로 mainA 을 넘겼지만 이게 복사되어서 적용되었기 때문에 실제 mainA 의 값은 변한게 없음
	mainA.withdrawValue(500)
	fmt.Println(mainA.balance) //500

	//값을 타입으로 가지지만, 값을 반환하는 메소드를 호출해버리면 > 함수가 수행되고 적용된 사항을 가진 객체를 반환하고, 그 값은 당연하게 잘 되어있겠지
	mainB := mainA.withdrawValue2(500)
	fmt.Println(mainB.balance) //0
}
