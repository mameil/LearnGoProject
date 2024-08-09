package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangeDataVer2(arg *Data) {
	arg.value = 9999
	arg.data[100] = 9999
}

func main() {
	var data Data    //data 라는 객체를 선언하고
	ChangeData(data) //함수의 파라미터로 data 라는 객체를 던지면 ChangeData 함수에서 받아서 인자로써 수행될 때 모든 변수값이 복사되기 때문에
	//ChangeData 의 인자인 arg 와 16라인에서 선언한 data 는 다른 메모리 공간을 갖는 변수

	fmt.Println("Value : ", data.value)                //data 라는 객체에다가 업데이트 친게 아니기 때문에 초기값 그대로 0 으로 나옴
	fmt.Println("Data's 101 index : ", data.data[100]) //data 라는 객체에다가 업데이트 친 게 아니기 때문에 초기값 그대로 0 으로 나옴
	// +심지어 객체를 복사하는 개념이기 때문에 짧은 시간내에 자주 호출되면 그만큼 메모리를 순간적으로 잡아먹을 예정이기 때문에 위험

	//이와 같은 문제를 해결하기 위해서 포인터라는 개념을 사용해서 처리해주면 된다
	dataPointer := &data                                      //data 라는 객체를 저장하고 있는 주소를 가리키는 포인터를 선언해주고
	ChangeDataVer2(dataPointer)                               //포인터를 업데이트 쳐준다 > 해당 주소에 있는 객체의 필드값들을 설정해준다("메모리 주소"는 8바이트를 차지하기 떄문에 인자를 넘기면서 8바이트만 복사됨)
	fmt.Println("Value : ", dataPointer.value)                //data 를 가리키는 포인터를 업데이트 쳐줬기 때문에 정상적으로 둘다 9999로 업데이트 쳐져있음
	fmt.Println("Data's 101 index : ", dataPointer.data[100]) //data 를 가리키는 포인터를 업데이트 쳐줬기 때문에 정상적으로 둘다 9999로 업데이트 쳐져있음

	//결론적으로는 포인터를 사용함으로써 효율적으로 구동되는 애플리케이션을 만들 수 있다는 것
	//근데 맨날 객체 생성하고, 포인터 쌍으로 만들어주고 이러한 과정을 한꺼번에 수행할 수 있는 방법도 있다
	var mp1 *Data = &Data{value: 1111}
	mp2 := &Data{value: 2222}
	fmt.Println(mp1.value, "/", mp2.value)

	//포인터를 기반으로 객체의 주소 정보로 객체를 조작하는 방법에 대해서 알아봤으니
	//이제는 실제로 메모리에 할당된 데이터에 접근하는 방법에 대해서 알아보자
	//"인스턴스"란, 메모리에 할당된 데이터의 실체를 의미한다
	//golang 에서는 그래서 자바에서 사용하는 것 처럼 new 라는 키워드를 기반으로 인스턴스를 생성하는 것을 지원한다
	var pointer1 *Data = new(Data)
	fmt.Println("===================")
	fmt.Println(pointer1) //초기값으로 초기화되어서 인스턴스화된 Data 객체가 출력됨
	fmt.Println("===================")

	//뭔가 C 언어처럼 메모리 할당해주고, 메모리 회수해주고 요런 작업도 필요한 것 같은데
	//golang 언어의 특징 중 하나로 GC 기능이 있다는 거
	//책에서 기억하라는 4가지는 아래와 같음
	//1. 인스턴스는 메모리에 생성된 데이터의 실체임
	//3. 포인터를 이용해서 인스턴스를 가리키게할 수 있음
	//3. 함수 호풀 시 포인터 인수를 통해서 인스턴스를 입력받고, 그 값을 변경할 수 있다 >> 위에 func ~~ (arg *Data) 이런 함수로 작업해서 된거 기억
	//4. 쓸모 없어진 인스턴스는 GC 가 알아서 판단해서 메모리를 회수해줌

	//프로그래밍에서 스택 메모리와 힙 메모리가 있는데,
	//기본적으로 스택 메모리가 힙 메모리보다 더 효율적이라서 스택 메모리에 프로그램에서 필요한 데이터들을 할당하는게 좋긴한데
	//스택 메모리는 함수 내부에서만 사용이 가능한 메모리 영역이다
	//함수 외부로 공개되는 메모리 공간이 힙 메모리에 적용된다

	//아레 메소드의 주석을 읽어보자
	userPointer := MakeUser("shim", 27)
	fmt.Println(userPointer)

	var actor = NewActor("shim", 100, 10.0)
	fmt.Println(actor.name)
	fmt.Println(actor.hp)
	fmt.Println(actor.speed)
}

type User struct {
	name string
	age  int
}

func MakeUser(name string, age int) *User {
	var u = User{name: name, age: age}
	return &u //아까 위에서 함수 내부에서 사용하는 메모리는 스택 영역의 메모리라고 했는데
	//이렇게 함수 내에서 스택 메모리에 올라간 데이터를 리턴해줘버리면
	//함수가 종료되면서 스택 메모리에서 할당된 메모리가 해제되기 때문에
	//어라 잘 안되어야 정상인데... 잘 데이터를 가리키고 출력도 잘됨
	//그 이유는 goLang 의 능력 덕분이다 > 함수 내에서 만든 변수를 외부로 공개하면 결국 스택 메모리를 오픈해버리는건데
	//이런 케이스를 goLang 이 찾아서 스택 메모리에 만들었던 아이들을 힙 메모리로 옮겨서 잘 되도록 처리해준다
	//결국은 goLang 은 어떤 타입의 변수이거나, 메모리의 수동 할당을 통해서 스택 영역을 사용할지, 힙 영역을 사용할지를 판단하는 것이 아니다
	//메모리 공간이 외부로 공개되는지, 아닌지를 기반으로 검사해서 스택 영역을 사용할지, 힙 영역을 사용할지를 판단한다

	//추가로 c/c++ 에서는 스택 메모리가 일정한 크기를 갖는 것에 비해서 goLang 은 스택 메모리의 크기가 계속해서 증가하는 동적 메모리 풀이다
	//이로 인해 메모리 효율성이 높으며, 재귀 호출 때문에 메모리가 고갈되는 일은 없다

}

type Actor struct {
	name  string
	hp    int
	speed float64
}

func NewActor(name string, hp int, speed float64) *Actor { //포인터를 던지라는건.. 결국은 함수 내에서 인스턴스를 만들고 그 인스턴스의 주소값을 던지라는 뜻이고
	//그렇게 해도 알아서 goLang 이 처리해준다니까 그냥 그렇게 해주면 될듯하다
	return &Actor{name: name, hp: hp, speed: speed}
}
