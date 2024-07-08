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

	//배열 vs 리스트
	//예를 들어서 이미 데이터가 어느 정도 차 있는 배열에서 가장 앞에다가 데이터를 추가하려고 하면
	//기존에 있는 모든 배열을 하나씩 미루고, 맨 앞에다가 추가하려는 데이터를 추가해줘야 한다 > 그리고 이건 결국 시간복잡도가 O(n) 이 된다
	//하지만 리스트의 맨 앞에 데이터를 넣어주려고 하면, 단순하게 Prev 에 넣으려는 데이터를 설정해주고, Next 에 기존에 첫 번쨰였던 요소를 설정해주면 된다
	//결국 계산식에 따르면 "배열 시작 주소 + (인덱스 X 타입 크기)" 만큼 소요되는데, 그 의미는 "처리하려는 요소 갯수과 상관 없이 추가하려는 값만 처리해주면 되니" O(1) 이 된다
	//빅오 표기법 > 알고리즘이 걸리는 시간을 측정하는 방법
	//예를 들어 aN^3 + bN^2 + cN^1 + b 이렇게 계산이 측정되면 빅오표시법으로 O(n^3) 이렇게 표기

	//이렇게만 보면 배열을 쓸 이유가 없어보이는데,
	//배열은 인덱싱에 최적화되어 있다
	//배열에서 데이터를 찾는건 O(1) 만큼 걸리지만 데이터를 넣고 뺴는건 O(n) 걸림
	//리스트는에서 데이터를 찾는건 O(n) 만큼 걸리지만 데이터를 넣고 뺴는건 O(1) 걸림
	//특정 데이터들을 묶어서 관리하는데 있어서 그 특정 데이터들, 즉 배열이나 리스트에 들어가는 값들이 "자주 변경이 일어나지 않고 인덱싱만 한다면" 배열이 유리하고 "자주 변경이 일어나면" 리스트가 유리하다

	//큐를 직접 구현해보자
	//큐의 특징을 생각해보면 선입선출, 즉 먼저 들어온 놈이 먼저 튀어나가는 구조이다
	//나오는 곳과 들어가는 곳이 다르고 가로로 누워있는 일자형 통을 생각해보면 된다
	//큐 안의 내부에 있는 데이터들의 순서는 보장된다
	//이걸 구현하는데 있어서는 근데 가장 먼저 들어간 데이터를 빼내기 때문에
	//배열로 처리한다고 가정하고 맨 첫번째가 빠지면 첫 번째가 비게되니 그 뒤에 있는 아이들을 한꺼번에 앞으로 땡겨줘야하고 이렇게 재인덱싱되는 과정에서 시간복잡도가 O(n) 이 계속 발생한다
	//리스트로 처리하면 이 과정이 O(1) 으로 처리되기 떄문에 오래걸리지 않음
	//아래 Queue 클래스 확인
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Push(i) //큐에 밀어넣고
	}

	for i := 3; i > 0; i-- { //하나씩 뽑아본다
		fmt.Println("Popped From Queue[STEP1] : ", queue.Pop())
	}

	for i := 0; i < queue.v.Len(); i++ { // 하나씩 뽑아본다
		fmt.Println("Popped From Queue[STEP2] : ", queue.Pop())
	}

	//스택을 직접 구현해보자
	//스택의 특징을 생각해보면 후입선출, 즉 마지막으로 들어온놈이 먼저 튀어나오는 구조이다
	//나오는 곳과 들어가는 곳이 같고 세로로 서있는 일자형 통을 생각해보면 된다
	//나오는 곳과 들어가는 곳이 같으니 기본적으로 아래까지 가기 위해선 위에 있는 아이들을 하나하나 제거하고 꺼내야한다
	//결국 가장 데이터가 튀어나오는 순서가 최신순 desc 으로 튀어나온다고 생각하면 된다
	//이걸 사용할떄는 "최신 것부터 처리가 필요한 상황"이다
	//예를 들면 "함수 호출 시" a() -> b() -> c() 순서로 함수 내에서 함수를 호출한다고 생각해보면 c() 함수가 종료되면 다시 b() 함수로 돌아가야하고, b() 함수가 종료되면 a() 함수로 돌아가야하는 그런 개념
	//아래 Stack 클래스 확인
	stack := NewStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")

	val := stack.Pop()
	for val != nil { //stack 에서 데이터를 뽑는데 nil 이 아닐때 까지 계속해서 루프를 돌면서 stack 에서 데이터를 뽑는다
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}

}

// 사실 shiftx2 을 통해서 non-project items 활성화시키고 보면 실제 Element 가 어떻게 구성되어있는지 확인 가능
type Element struct {
	Value interface{} //리스트를 구성하는 데이터 타입으로써 실제 리스트에 들어가는 값, interface{} 형태로 되어있기 때문에 아무런 타입의 값이 들어갈 수 있음
	Next  *Element    //리스트의 다음 요소 주소를 가리킴
	Prev  *Element    //리스트의 이전 요소 주소를 가리킴
	//"다음"과 "이전" 요소를 가리키고 있기 때문에 양방향성으로 볼 수 있음
}

// list 을 통해서 큐 생성
type Queue struct { //구조체 정의
	v *list.List
}

func (q *Queue) Push(val interface{}) { //요소 추가
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front) //요소를 반환하면서 삭제
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val) //세로형 통에서 값을 넣는건 맨 위에서 넣으니까 PushBack()
}

func (s *Stack) Pop() interface{} {
	front := s.v.Back() //세로형 통해서 값을 빼는건 맨 위에서 빼니까 Back()
	if front != nil {
		return s.v.Remove(front)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}
