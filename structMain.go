package main

import (
	"fmt"
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
	//fmt.Println(student.classList[2].className) //배열에 2개만 집어넣고 3 번째껄 접근하면 uncheckedException 으로 인한 Exception 발생
	//fmt.Println(student.classList[2].grade) //배열에 2개만 집어넣고 3 번째껄 접근하면 uncheckedException 으로 인한 Exception 발생
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

}
