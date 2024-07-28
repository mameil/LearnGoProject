package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
VocaReadme.md 를 참고해서 프로그램을 만들어보자
*/
func main() {
	if len(os.Args) < 3 {
		fmt.Println("해당 프로그램을 수행하기 위해선 \"find\" + \"찾을 단어\" + \"파일명\" 을 입력해주셔야 합니다")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려고하는 단어 :", word)
	fmt.Println("찾으려고하는 파일 경로 :", files)
	PrintAllFiles(files)
}

//func GetFileList(path string) ([]string, error) {
//	return filepath.Glob(path)
//}

func PrintAllFiles(files []string) {
	for _, path := range files {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일 경로를 읽어오는데 실패했습니다 : ", err)
			return
		}

		fmt.Println("지정된 파일 경로에 존재하는 파일 리스트")
		for _, name := range fileList {
			fmt.Println(name)
		}
	}
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 여는데 실패했습니다..! :", err)
		return
	}
	defer file.Close() //항상 파일을 열면 꼭 함수가 끝나는 시점에 꼭꼭 파일을 닫아주는거 잊지말자

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
