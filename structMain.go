package main

import (
	"fmt"
	"unsafe"
)

// 구조체는 여러 필드를 하나로 묶어서 만드는 기능이고 다른 타입들의 값을 변수 하나로 묶는다
type Person struct {
	name   string
	age    int
	height float64
	gender string
}

func printPerson(person Person) {
	fmt.Println("============================")
	fmt.Println("name:", person.name)
	fmt.Println("age:", person.age)
	fmt.Println("height:", person.height)
	fmt.Println("gender:", person.gender)
	fmt.Println("============================")
}

type GradeInfo struct {
	className string
	grade     string
}

type Student struct {
	person    Person
	classList []GradeInfo
}

func printStudent(student Student) {
	printPerson(student.person)
	fmt.Println(student.classList)
	fmt.Println(student.classList[0].className)
	fmt.Println(student.classList[0].grade)
	fmt.Println(student.classList[1].className)
	fmt.Println(student.classList[1].grade)
	//fmt.Println(student.classList[3].className) //배열에 2개만 집어넣고 3 번째껄 접근하면 uncheckedException 으로 인한 Exception 발생
	//fmt.Println(student.classList[3].grade) //배열에 2개만 집어넣고 3 번째껄 접근하면 uncheckedException 으로 인한 Exception 발생
}

type memTest struct {
	a int8 //1바이트
	b int  //8바이트
	c int8 //1바이트
	d int  //8바이트
	e int8 //1바이트
}

type memTest2 struct {
	a int8 //1바이트
	c int8 //1바이트
	e int8 //1바이트
	b int  //8바이트
	d int  //8바이트
}

func main() {
	var kdshim Person //객체를 선언하고 필드에 값을 설정해준다
	kdshim.name = "심규도"
	kdshim.age = 27
	kdshim.height = 170.3
	kdshim.gender = "MALE"

	printPerson(kdshim)

	var ordinary Person //객체를 그냥 선언하면 기본적으로 필드 타입의 초기값으로 설정된다
	printPerson(ordinary)

	initPerson := Person{"First", 10, 180.1, "MALE"} //자바에서의 AllArgConstructor 처럼 전체 필드를 가진 생성자 사용가능
	printPerson(initPerson)
	partialPerson := Person{name: "Second", gender: "FEMALE", age: 20} //코틀린의 Data 클래스에서 몇개의 필드만을 설정해주는 생성자 사용 가능
	printPerson(partialPerson)

	student1 := Student{ //객체 in 객체와 같은 구조로 설정하는 것도 가능하다
		person: kdshim,
		classList: []GradeInfo{
			{className: "컴퓨터구조", grade: "A"},
			{className: "알고리즘", grade: "C"},
		},
	}
	printStudent(student1)

	doppelganger := student1 //구조체를 복사하는건 그냥 "대입"만 해줘도 전체적으로 다 복사됨
	printStudent(doppelganger)

	//객체의 메모리 정렬 : string(16) + int(8) + float48(8) + string(16)
	//go 에서는 메모리 정렬이라는 개념이 존재한다 >
	//컴퓨터의 레지스터 크기에 따라서 컴퓨너의 비트를 측정하는데 레지스터 크기가 8바이트라는 이야기는 한 번 연산에 8바이트 크기를 연산할 수 있다는 이야기
	//기본적으로 컴퓨터는 8의 배수로 돌아야 효율적으로 데이터가 수행되기 때문에 알아서 컴퓨터가 8바이트를 맞추고, 이 맞출 때 데이터가 8의 배수보다 부족하면 메모리를 강제로 8의 배수로 맞춘다 그리고 이 맞추면서 추가되는 공간이 메모리 패딩이라는 개념이다

	//예를 들어서 memTest 객체를 보면 1-8-1-8-1 순서로 되어있는데
	//코드를 라인 바이 라인으로 읽는 과정에서 1바이트를 읽으면 8의 배수로 해야하기 때문에 실제로 변수별로 메모리가 할당되는걸 보면 아래와 같은 구조로 설정된다
	//1 byte+메모리패딩(7 byte)
	//8byte
	//1 byte+메모리패딩(7 byte)
	//8byte
	//1 byte+메모리패딩(7 byte) >> total : 40
	//위와 같은 불필요한 메모리 패딩읠 추가를 피하기 위해서 코드의 순서를 바꿔준다 > 8 바이트의 배수가 맞도록 8의 배수가 아닌 바이트들을 한꺼번에 몰았다
	//그렇게되면
	//1 byte + 1 byte + 1 byte + 메모리패딩(5 byte)
	//8byte
	//8byte >> total :  24
	//이렇게 보면 객체의 순서를 변경한걸로 > 객체당 빈공간으로 21바이트를 사용하던걸, 5바이트로 줄일 수 었다는 점이다
	//결론 : 불필요한 메모리 낭비를 줄이기 위해 "작은 크기 필드값을 앞에 배치하자"
	//자바나 코틀린에서도 이러한 개념이 적용되는지가 궁금하긴 하다 > 나중에 확인해보자
	fmt.Println(unsafe.Sizeof(memTest{1, 2, 3, 4, 5}))  //40
	fmt.Println(unsafe.Sizeof(memTest2{1, 2, 3, 4, 5})) //24

}
