package main

import "fmt"

func main() {
	//slice 란 golang 에서 제공하는 동적으로 사이즈가 조절되는 배열 + 배열의 일부를 슬라이스 형식으로 잘라서 사용할 수도 있음

	//우리가 기억하는 배열은 사이즈가 선언할 때 고정해서 선언하기 때문에 배열을 선언하는 순간 뭐 그 이상을 추가하거나 하는것이 불가능하다 >> 만약에 추가하고 싶다면 배열을 새로 만들어야함
	//위 example : var array [10]int

	//slice 는 단순하게 배열의 사이즈를 입력해주지 않으면 동적 배열로써 선언됨
	var mySlice []int //초기값으로써 배열의 사이즈가 0인 슬라이스
	if len(mySlice) == 0 {
		fmt.Println("슬라이스는 초기화되면 길이가 0인 배열이다")
	}

	//slice 의 가장 초기값 사이즈는 0읻네 1번째에 숫자를 넣어줄려고하면 index out of range [1] with length 0 "runtime error" 에러가 발생한다
	//mySlice[1] = 9999

	//slice 을 선언하면서 함께 초기값들을 설정해주는 것도 가능
	var mySlice2 = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice2)
	var mySlice3 = []int{1, 2: 2, 4: 3} //배열의 2번째에는 2를 넣어주고 4번째에는 3을 넣어주고 + 나머지 선언해주지 않은 사이의 자리값 index 에는 int 의 초기값인 0 을 넣어준다
	fmt.Println(mySlice3)

	//make() 함수를 통해서 slice 를 생성하는 방법
	var mySlice4 = make([]int, 3) //make 함수의 첫 번째 parameter 에는 어떤 타입의 slice 를 만들 것인지 타입을 작성해주고, 두 번째 parameter 에는 사이즈를 작성해준다
	fmt.Println(mySlice4)

	//slice 를 순회하는 기능 #1
	for i := 0; i < len(mySlice4); i++ {
		fmt.Print(mySlice4[i])
	}
	fmt.Println()
	//slice 를 순회하는 기능 #2
	for i, v := range mySlice4 {
		fmt.Println(i, ":", v)
	}

	//slice 에 값을 추가하는 방법 >> java 의 ArrayList 에 add 을 통해서 추가하는 것 처럼 append() 을 통해서 숫자를 추가
	mySlice5 := make([]int, 3)     // 초기값 으로 인해 [0, 0, 0]
	mySlice5 = append(mySlice5, 1) //거기에 1을 추가했으니 slice 의 사이즈는 늘어나고, 마지막 index 에 1이 추가된다
	fmt.Println(len(mySlice5))     //4
	fmt.Println(mySlice5)          //[0, 0, 0, 1]

	//slice 에 값을 추가하는 방법2 >> append() 함수에서는 parameter 가 여러개가 들어갈 수 있음
	mySlice6 := make([]int, 1)
	mySlice6 = append(mySlice6, 1, 2, 3, 4, 5)
	fmt.Println(len(mySlice6)) // 6
	fmt.Println(mySlice6)      //[0, 1, 2, 3, 4, 5]

	/**
	slice 의 내부구현체는 string 처럼 struct 의 형태로 되어있음
	type SliceHeader struct {
		Data uintptr //실제 배열을 가리키는 포인터
		Len  int     //요소의 개수
		Cap  int	 //실제 배열의 길이
	}
	그래서 make() 함수를 통해서 slice 를 생성할 때 3번째 파라미터에 cap 도 추가해서 선언할 수 있음
	*/
	mySlice7 := make([]int, 3, 5)                //이 의미는 int 타입의 요소를 가지는 5개 크기의 배열을 만들고 우선 3개만 기본값으로 초기화를 해둔다는 의미
	fmt.Println("mySlice7", mySlice7)            //[0, 0, 0]
	fmt.Println("mySlice7's len", len(mySlice7)) //3
	fmt.Println("mySlice7's cap", cap(mySlice7)) //5

	//슬라이스와 배열의 동작차이
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	//각각 0번째 인덱스에 있는 숫자를 바꿔줌
	changeArray(array)
	changeSlice(slice)

	fmt.Println(array) //[1,2,3,4,5] 왜 여긴 반영이 안되고
	fmt.Println(slice) //[999,2,3,4,5] 왜 여긴 반영이 되었을까
	//array 가 복사되지 않는건, 저번에 포인터할 때 나왔던 이야기이지만, golang 에서 함수에 파라미터로 넘기는 순간 값이 복사되어서 들어가기 때문에 함수 내부에서 변경된 값은 결국 전혀 상관없는(대입하면서 복사된) 값에 적용된다
	//하지만 slice 는 왜 반영이 될까? >> 이건 애초에 slice 내부에서 객체에 대한 포인터를 가지고 있는거고, 포인터가 가리키고 있는 곳의 값을 변경하는 것이기 때문에 함수 외부에서 인자로 넘긴 값을 확인해봐도 포인터가 가리킨 곳은 변경됬기 때문에 잘 반영되걸로 보임

	//append 를 잘못 사용하는 경우
	//append() 메소드는 호출되면 우선 slice 에 갑슬 추가할 수 있는 빈공간이 있는지 확인하고, 남은 빈 공간을 계산하는 방법은 cap - len 이다
	//남은 빈 공간의 개수가 추가하는 값의 개수보다 크거나 같은 경우 배열 뒤에 추가한다
	mySlice8 := append(make([]int, 0, 5), 1, 2, 3)
	//mySlice 는 메모리 공간까지 해서 보면 [1,2,3,0,0] 형식으로 되어있음
	fmt.Println(mySlice8)
	mySlice81 := append(mySlice8, 4, 5)

	fmt.Println("mySlice81:", mySlice81, len(mySlice81), cap(mySlice81)) //[1,2,3,4,5] mySlice81 과 mySlice8 은 모두 같은 배열을 바라보는데,
	fmt.Println("mySlice8:", mySlice8, len(mySlice8), cap(mySlice8))     //[1,2,3] 이렇게 처리되는데 이게 append() 함수를 잘못 사용한 경우이다
	// mySlice8 의 요소갯수는 3개  + 전체배열길이는 5개 / mySlice81 의 요소갯수는 5개 + 전체배열길이는 5개
	//개인적으로는 어? 왜 같은걸 바라보고 있는데 mySlice8 을 조회했을 때 1,2,3 만 나오나 싶었는데 len 으로 잡혀있어서 그런듯

	//이 상황에서
	mySlice81[1] = 9999                                                  //이렇게 대입해버리면?
	fmt.Println("mySlice8:", mySlice8, len(mySlice8), cap(mySlice8))     //[1,9999,3] 이렇게 mySlice8 에도 같이 영향을 받는다 > 왜? 같은 배열을 바라보고 있으니까
	fmt.Println("mySlice81:", mySlice81, len(mySlice81), cap(mySlice81)) //[1,9999,3,4,5] 이렇게 mySlice81 은 당연히 바뀌어야 하는거고

	//또한 값을 추가해보면?
	mySlice8 = append(mySlice8, 500)
	fmt.Println("mySlice8:", mySlice8, len(mySlice8), cap(mySlice8))     //[1,9999,3,500] 우선 mySlice8 을 기준으로는 cap 이 5니까 값을 추가할 수 있으니, 정상적으로 500 이란 값을 추가해주었다
	fmt.Println("mySlice81:", mySlice81, len(mySlice81), cap(mySlice81)) //[1,9999,3,500,5] 여기가 문제인데, mySlice81 과 mySlice8 은 같은 배열을 바라보기 때문에 mySlice8 에서 값을 변경해버리면 mySlice81 에도 영향을 미치게 된다
}

func changeArray(array [5]int) {
	array[0] = 999
}

func changeSlice(slice []int) {
	slice[0] = 999
}
