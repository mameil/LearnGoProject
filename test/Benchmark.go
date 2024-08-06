package test

/**
Go 의 특징으로 테스트 코드이외에 코드의 성능을 측정해주는 벤치마크 기능도 있다!
이것도 동일하게 테스트 패키지를 통해서 지원하고 동일하게 표현의 규약이 존재
1. 파일 명이 _test.go 로 끝나야 한다
	> 테스트 코드의 규약과 동일함
2. testing 패키지를 임포트 받아야 한다
	> 테스트 코드의 규약과 동일함
3. 벤치마크 코드는 func BenchmarkXxxx(b *testing.B) 형식이여야 한다
	> 테스트 코드의 규약과 동일하지만 이름이 명확하게 테스팅과 벤치마킹이 분리되어있음




*/
