package main

import (
	"fmt"
	"time"
)

// main은 실행시 환영 메시지와 현재 시스템 시간을
// 출력하여 프로그램의 정상 동작 여부를 확인하는 역할을 하는 함수
func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}
