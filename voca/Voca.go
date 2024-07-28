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
	findInfos := []FindInfo{}
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

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	fileList, err := GetFileList(path)
	if err != nil {
		informReturn("파일 읽어오는데 실패했습니다 : " + err.Error())
	}

	for _, fileName := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, fileName))
	}

	return findInfos
}

func FindWordInFile(word, fileName string) FindInfo {
	findInfo := FindInfo{fileName, []LineInfo{}}
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
			findInfo.lines = append(findInfo.lines, LineInfo{lineNum, line})
		}
		lineNum++
	}

	return findInfo
}
