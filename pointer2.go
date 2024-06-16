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

}
