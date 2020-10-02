package main

import "fmt"

/*
*Даны два массива (слайса) А и В, каждый из них содержит 2 значения.
Первая цифра - начало входа пользователя в сеть, вторая — выход пользователя из сети.
Необходимо вывести на экран время нахождения пользователей А и В в сети вместе.
Пример: Входные данные: A[10,18] B[9, 13] Выходные данные: Общее время нахождения в сети с 10 до 13
*/

/*
1. продумать реализацию с замкнутым пространством
2. тесты
*/

func main() {
	a := []int{10, 18}
	b := []int{9, 13}

	start, end := find(a, b)
	fmt.Println(start, end)

}

func find(a, b []int) (int, int) {
	var maxStart int
	var maxEnd int
	if a[1] > b[0] || b[1] > a[0] {
		if a[0] >= b[0] {
			maxStart = a[0]
			if a[1] >= b[1] {
				maxEnd = b[1]
			} else {
				maxEnd = a[1]
			}
		} else if a[0] < b[0] {
			maxStart = b[0]
			if a[1] >= b[1] {
				maxEnd = b[1]
			} else {
				maxEnd = a[1]
			}
		}
	}
	return maxStart, maxEnd
}
