package main

import (
	"fmt"
	"math"
)

func SquareRootResolver() {
	var a float32
	var b float32
	var c float32
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
	var a int32
	var b int32
	_, err := fmt.Scanf("%d %d", a, b)
	if err != nil {
		return
	}
	if b != 0 {
		fmt.Println(a / b)
	} else {
		fmt.Println("Impossible")
	}
}
