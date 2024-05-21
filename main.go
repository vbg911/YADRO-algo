package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func canSortContainers(n int, containers [][]int) string {
	// Суммарное количество шаров каждого цвета
	colorSums := make([]int, n)
	// Суммарное количество шаров в каждом контейнере
	containerSums := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			colorSums[j] += containers[i][j]
			containerSums[i] += containers[i][j]
		}
	}

	// Сортируем оба списка и сравниваем
	sort.Ints(colorSums)
	sort.Ints(containerSums)

	if reflect.DeepEqual(colorSums, containerSums) {
		return "yes"
	} else {
		return "no"
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Чтение количества контейнеров
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// проверка ограничения из условия задачи
	if n <= 1 || n >= 100 {
		fmt.Print("Ограничения: 1<= n <= 100")
		return
	}

	// Чтение содержимого контейнеров
	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		scanner.Scan()
		values := strings.Fields(scanner.Text())
		for j := 0; j < n; j++ {
			amount, err := strconv.Atoi(values[j])
			if err != nil {
				fmt.Print(err.Error())
				return
			}
			// проверка ограничения из условия задачи
			if amount <= 0 || amount >= 1000000000 {
				fmt.Println("0 <= количество шаров одного цвета в одном контейнере <= 1000000000")
				return
			}
			containers[i][j] = amount
		}
	}

	// Проверка возможности сортировки
	result := canSortContainers(n, containers)
	fmt.Println(result)
}
