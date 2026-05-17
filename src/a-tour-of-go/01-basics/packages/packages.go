package main

import (
	"fmt"
	"math/rand"
)

// 프로그램의 실행 시작점 함수.
// 0 ~ 9의 무작위 정수를 생성하여 출력합니다.
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
