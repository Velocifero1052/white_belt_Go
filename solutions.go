package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math"
	"strconv"
)

func SquareRootResolver() {
	var a, b, c float32
	_, err := fmt.Scanf("%f %f %f", &a, &b, &c)
	if err != nil {
		return
	}

	if a != 0.0 && b != 0.0 {
		underRootExpression := b*b - 4*a*c
		if underRootExpression < 0.0 {
			return
		} else if underRootExpression == 0.0 {
			fmt.Printf("%.2f", -b/(2.0*a))
		} else {
			underRootExpression = float32(math.Sqrt(float64(underRootExpression)))
			x1 := (-b + underRootExpression) / (2.0 * a)
			x2 := (-b - underRootExpression) / (2.0 * a)
			fmt.Printf("%.2f %.2f", x1, x2)
		}
	} else if a == 0.0 && b != 0.0 && c != 0.0 {
		fmt.Printf("%.2f", -c/b)
	} else if a != 0.0 && b == 0.0 && c == 0.0 || a == 0.0 && b != 0.0 && c == 0.0 {
		println(0.0)
	} else if a != 0.0 && b == 0.0 && c != 0.0 && -c/a >= 0.0 {
		res := float32(math.Sqrt(float64(-c / a)))
		fmt.Printf("%.2f %.2f", res, -res)
	}
}

func soCalledImpossibleFunc() {
	var a, b int32
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if b != 0 {
		fmt.Println(a / b)
	} else {
		fmt.Println("Impossible")
	}
}

func priceCalculations() {
	var n, a, b, x, y, finalPrice float32
	_, err := fmt.Scanf("%f %f %f %f %f", &n, &a, &b, &x, &y)
	finalPrice = n
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if n > a && n > b {
		finalPrice = n - n*(y/100.0)
	} else if n > a && n <= b {
		finalPrice = n - n*(x/100.0)
	}
	fmt.Println(finalPrice)
}

func printEvenNumbers() {
	var x, y, i int
	first := true

	_, err := fmt.Scanf("%d %d", &x, &y)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if x == y && x%2 == 0 {
		fmt.Println(x)
		return
	} else if x%2 == 0 {
		i = x
	} else {
		i = x + 1
	}

	for ; i <= y; i += 2 {
		if first {
			fmt.Printf("%d", i)
			first = false
		} else {
			fmt.Printf(" %d", i)
		}
	}
}

func secondOccurrenceOfF() {
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		fmt.Println(err)
		return
	}
	foundIndex := -2
	first := true

	for i := 0; i < len(word); i++ {
		if string(word[i]) == "f" && first {
			foundIndex = -1
			first = false
		} else if string(word[i]) == "f" && !first {
			foundIndex = i
			break
		}
	}
	fmt.Println(foundIndex)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(a, a%b)
}

func integerToBinaryConversion() {
	var n int64
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(strconv.FormatInt(n, 2))
}

func lower_than_average() {
	var size_of_array int
	_, err := fmt.Scanf("%d\n", &size_of_array)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	temps := make([]int, size_of_array)

	for i := 0; i < size_of_array; i++ {
		var num int
		_, err = fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		temps[i] = num
	}

	sum := 0
	for temp := range temps {
		sum += temp
	}

	average := sum / len(temps)
	var higherThanAverage []int
	for temp := range temps {
		if temp > average {
			higherThanAverage = append(higherThanAverage, temp)
		}
	}

	fmt.Println(len(higherThanAverage))
	first := true
	for average := range higherThanAverage {
		if first {
			first = false
			fmt.Printf("%d", average)
		} else {
			fmt.Printf(" %d", average)
		}
	}

}

func worry_count() {
	var people []int
	var numberOfCommands int
	_, err := fmt.Scanf("%d\n", &numberOfCommands)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defaultValue := 1
	worryValue := -1
	var command string
	for i := 0; i < numberOfCommands; i++ {
		_, err = fmt.Scanf("%s", &command)
		if command == "WORRY" {
			var index int
			_, err = fmt.Scanf("%d\n", &index)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			people[index] = worryValue
		} else if command == "COME" {
			var number int
			_, err = fmt.Scanf("%d\n", &number)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			if number > 0 {
				for j := 0; j < number; j++ {
					people = append(people, defaultValue)
				}
			} else if number < 0 {
				number = -number
				for j := 0; j < number; j++ {
					people = people[:len(people)-1]
				}
			}

		} else if command == "WORRY_COUNT" {
			worryCount := slices.IndexFunc(people, func(elem int) bool { return elem == worryValue })
			fmt.Println(worryCount)
		}
	}
}
