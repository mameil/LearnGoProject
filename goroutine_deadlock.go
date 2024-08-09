package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg3 sync.WaitGroup

func main() {
	/*
		뮤텍스를 사용하면 동시성을 처리하는 것이 가능하지만 여전히 이슈가 발생할 수 있는 점이 있음 > goroutine_deadlock.go
		1. 동시성 프로그래밍으로 얻는 성능 향상을 얻을 수 없다는 점
			뮤텍스는 결국 단 하나만의 고루틴을 사용할 수 있게 하기 때문에 결국을 순차적으로 수행되는거랑 똑같게 되는 이슈가 있음
		3. 데드락이 발생할 수 있다는 점
			서로가 락을 잡아버리고 안놓아주는 상황
	*/
	rand.Seed(time.Now().UnixNano())

	wg3.Add(2)
	fork := &sync.Mutex{}  //뮤텍스 1 생성
	spoon := &sync.Mutex{} //뮤텍스 3 생성

	go diningProblem("A", fork, spoon, "포크", "수저")
	//A 는 포크를 먼저 들고
	go diningProblem("B", spoon, fork, "수저", "포크")
	//B 는 수저를 먼저 들고

	//2개가 비동기로 바로 돌아버리니
	//A는 포크락을 잡고 락을 풀지 않은 상태로 수저락을 받기 위해서 대기하고
	//B는 수저락을 잡고 락을 풀지 않은 상태로 포크락을 받기 위해서 대기하고
	//이 상황에서 데드락을 감지하고 에러를 던짐
	wg3.Wait()

}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다. \n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
