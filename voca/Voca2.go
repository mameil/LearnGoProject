package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNum int
	line    string
}

type FindInfo struct {
	fileName string
	lines    []LineInfo
}

/*
*
* VocaReadme.md 를 참고해서 프로그램을 만들어보자 ver 3
* 기존엔 모든 파일 검색을 하나의 main() 고루틴에서 실행했지만, 검색할 파일이 늘어나면 결국 늘어난 만큼 다음 파일들이 기다리는 이슈가 발생
* 여기서 고루틴을 활용해서 파일 갯수가 늘어나더라도, 빠르게 검색되도록 수정해보자
* 파일 당 고루틴 하나씩 설정해서 병렬로 동시에 수행되게끔하고, 각각의 고루틴 결과를 채널로 모아서 결과를 던지는 방식으로 구현해보자
* 이렇게 작업을 분배하고, 다시 거둬들이는 방식을 보면서 "뿌리고 거두기"(Scatter - Gather) 방식이라고 부름

>> 실행해보니까 파일이 2개인데도 빠른 듯한데 ㅋㅋㅋㅋ 신기하다
*/
func main() {
	if len(os.Args) < 3 {
		fmt.Println("해당 프로그램을 수행하기 위해선 \"find\" + \"찾을 단어\" + \"파일명\" 을 입력해주셔야 합니다")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	fmt.Println("================================================")
	fmt.Println("파일 검색 프로그램이 시작됩니다......")
	fmt.Println("찾으려는 단어 : ", word)
	fmt.Println("찾을 대상의 파일 수 : ", len(files))
	var findInfos []FindInfo
	for _, path := range files {
		fmt.Println("================================================")
		fmt.Println("파일 검색을 시작합니다...")
		fmt.Println("파일 경로 : ", path)
		findInfos = append(findInfos, FindWordInAllFiles2(word, path)...)
		fmt.Println("파일 검색이 완료되어서 정리 중입니다")
		fmt.Println("================================================")
	}

	for _, findInfo := range findInfos {
		fmt.Println("================================================")
		fmt.Println("파일명 : ", findInfo.fileName)
		fmt.Println("아래는 \"" + word + "\" 단어가 포함된 라인입니다")
		for _, lineInfo := range findInfo.lines {
			fmt.Printf("%s:%d | %s\n", findInfo.fileName, lineInfo.lineNum, lineInfo.line)
		}
		fmt.Println("================================================")
	}
}

func FindWordInAllFiles2(word, path string) []FindInfo {
	var findInfos []FindInfo
	fileList, err := filepath.Glob(path)

	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다 : ", err)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(fileList)
	recvCnt := 0

	for _, fileName := range fileList {
		go FindWordInFile2(word, fileName, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}

	return findInfos
}

func FindWordInFile2(word, fileName string, ch chan FindInfo) {
	findInfo := FindInfo{fileName, []LineInfo{}}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다! ", fileName)
		ch <- findInfo
		return
	}

	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}
