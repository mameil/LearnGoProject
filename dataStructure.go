package main

import (
	"container/list"
	"fmt"
)

/**
자료 구조란, 데이터를 저장하는 구조를 의미한다, 지금까지 사용하던, 배열, 슬라이드도 자료 구조 중 하나
추가적으로 학습할 내용
- 리스트 : 비연속 메모리를 사용해 요소를 저장한다, 요소 삽입과 삭제가 배열보다 빠름
- 큐 : First In First Out 구조(FIFO)
- 스택 : Last In First Out 구조(LIFO)
- 링 : 처음과 끝이 연결된 리스트로 크기가 고정된 구조
- 맵 : key-value 형식으로 자료가 저장되는 구조
*/

func main() {
	/**
	리스트와의 배열
	배열은 연속된 메모리에 데이터를 저장하는 반면
	리스트는 불연속된 메모리에 데이터를 저장한다

	리스트의 구조를 살펴보자면, 포인터를 기반으로 다음, 이전의 데이터를 가리키고 있다 >> 이렇게 생겼기 떄문에 LinkedList 라고 부르기도 함
	@see Element
	예를 들어서 배열을 선언하고 데이터 1, 2, 3 을 넣었다고 생각해보면
	데이터1의 주소값, 데이터2의 주소값, 데이터3의 주소값 모두 청므 선언하는데 있어서 알아서 챡챡 메모리 어디가에 올라갈꺼고
	리스트에서 단순히 이 각각의 데이터를 "연결해서" 리스트 데이터로써 관리하기 떄문에 "불연속된 데이터를 저장한다"라고 이야기할 수 있음
	추가로 배열은 "연속된 데이터를 저장한다"라고 되어있음
	*/

	v := list.New()
	e4 := v.PushBack(4)   //리스트의 가장 마지막에 4를 넣어준다
	e1 := v.PushFront(1)  //리스트의 가장 앞에 1을 넣어준다
	v.InsertAfter(2, e1)  //리스트의 1(가장 앞) 뒤에 2를 넣어준다
	v.InsertBefore(3, e4) //리스트의 4(가장 뒤) 앞에 3을 넣어준다

	//전위순회
	for i := v.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ") //1 2 3 4
	}
	fmt.Println()

	//후위순회
	for i := v.Back(); i != nil; i = i.Prev() {
		fmt.Print(i.Value, " ") //4 3 2 1
	}

}

// 사실 shiftx2 을 통해서 non-project items 활성화시키고 보면 실제 Element 가 어떻게 구성되어있는지 확인 가능
type Element struct {
	Value interface{} //리스트를 구성하는 데이터 타입으로써 실제 리스트에 들어가는 값, interface{} 형태로 되어있기 때문에 아무런 타입의 값이 들어갈 수 있음
	Next  *Element    //리스트의 다음 요소 주소를 가리킴
	Prev  *Element    //리스트의 이전 요소 주소를 가리킴
	//"다음"과 "이전" 요소를 가리키고 있기 때문에 양방향성으로 볼 수 있음
}

type Animal struct {
	Name string
}
