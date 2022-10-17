package main

import (
	"errors"
	"fmt"
)

func validateDate(y, m, d int) error {
	if y <= 0 {
		return errors.New("year must greater than 0")
	}
	if m < 1 || m > 12 {
		return errors.New("month must between 1~12")
	}
	if d < 1 || d > 31 {
		return errors.New("day must between 1~31")
	}
	return nil
}

func calcYear(y int) int {
	y -= 1
	y = y % 400             // 윤년 제외
	result := (y / 100) * 5 // 100년 단위로 나눈값 저장
	y = y % 100
	result += y + y/4
	return result % 7
}

func calcMonth(m int) int {
	m -= 1
	result := 0
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 0; i < 12; i++ {
		if m >= i+1 {
			result += months[i]
		} else {
			break
		}
	}
	return result % 7
}

func calcDay(d int) int {
	return d % 7
}

func calcDate(y, m, d int) int {
	result := calcYear(y) + calcMonth(m) + calcDay(d)
	result %= 7
	return result
}

func main() {
	fmt.Println("년월일을 yyyy mm dd 형태로 입력해주세요.")
	var y, m, d int
	_, err := fmt.Scanf("%d %d %d", &y, &m, &d)
	if err != nil {
		panic(err)
	}

	if err := validateDate(y, m, d); err != nil {
		panic(err)
	}

	dayOfTheWeek := [7]string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"}
	fmt.Printf("%d년 %d월 %d일은 %s입니다.\n", y, m, d, dayOfTheWeek[calcDate(y, m, d)])
}
