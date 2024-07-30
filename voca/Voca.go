package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo2 struct {
	lineNum int
	line    string
}

type FindInfo2 struct {
	fileName string
	lines    []LineInfo2
}

/*
*
* VocaReadme.md 를 참고해서 프로그램을 만들어보자
* goLang 을 이용해서 프로그램을 만든건 처음이니 혹시 몰라서 적어둠
* goLang 에서 run > edit configuration 에 들어가서 "output Directory" 에다가 build 해서 나온 "실행파일"이 어디에 저장됬으면 하는지 지정해주고
* terminal 에서 프로그램을 실행 시 사용할 환경변수와 함께 넣어줘서 실행시키면 잘만됨
 */
func main() {
	if len(os.Args) < 3 {
		informReturn("해당 프로그램을 수행하기 위해선 \"find\" + \"찾을 단어\" + \"파일명\" 을 입력해주셔야 합니다")
	}

	word := os.Args[1]
	files := os.Args[2:]

	fmt.Println("================================================")
	fmt.Println("파일 검색 프로그램이 시작됩니다......")
	fmt.Println("찾으려는 단어 : ", word)
	fmt.Println("찾을 대상의 파일 수 : ", len(files))
	findInfos := []FindInfo2{}
	for _, path := range files {
		fmt.Println("================================================")
		fmt.Println("파일 검색을 시작합니다...")
		fmt.Println("파일 경로 : ", path)
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
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

func informReturn(errMsg string) {
	fmt.Println("================================================")
	fmt.Println("프로그램 수행 중, 에러가 발생해서 프로그램을 종료합니다......")
	fmt.Println(errMsg)
	fmt.Println("================================================")
	return
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo2 {
	findInfos := []FindInfo2{}

	fileList, err := GetFileList(path)
	if err != nil {
		informReturn("파일 읽어오는데 실패했습니다 : " + err.Error())
	}

	for _, fileName := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, fileName))
	}

	return findInfos
}

func FindWordInFile(word, fileName string) FindInfo2 {
	findInfo := FindInfo2{fileName, []LineInfo2{}}
	file, err := os.Open(fileName)

	if err != nil {
		informReturn("파일을 여는데 실패했습니다 : " + err.Error())
	}
	defer file.Close()

	lineNum := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo2{lineNum, line})
		}
		lineNum++
	}

	return findInfo
}
