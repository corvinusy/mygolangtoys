package main

import (
    "fmt"
)

func main() {

	x_values := make([]float64,0)
	y_values := make([]float64,0)
	l_values := make([]float64,0)

	for i := 0; i < 10; i++ {
		x_values = append(x_values, float64(i+1))
		y_values = append(y_values, oF(x_values[i]))
		
		lp := lagrangePolynome(float64(i+2), x_values, y_values)
		
		if lp == oF(float64(i+2)) {
			break
		}

		l_values = append(l_values, lp)
	}
	
//	fmt.Println(l_values)

	result := 0.0

	for _, lv := range l_values {
		result += lv
	}

	fmt.Println(int64(result))

	return
}
/*----------------------------------------------------------------------------*/
func oF(x float64) float64 {


	result := 1 - x + x*x - x*x*x + x*x*x*x - x*x*x*x*x + x*x*x*x*x*x - 
		x*x*x*x*x*x*x + x*x*x*x*x*x*x*x - x*x*x*x*x*x*x*x*x + x*x*x*x*x*x*x*x*x*x

//	result := x * x * x
	return result
}
/*----------------------------------------------------------------------------*/
func lagrangePolynome(x float64, x_values, y_values []float64) float64 {

	result := 0.0

	for i := 0; i < len(x_values); i++ {
		base := 1.0
		for j := 0; j < len(x_values); j++ {
			if i == j { continue }
			base *= (x - x_values[j])/(x_values[i] - x_values[j])
		}
		result += base*y_values[i]
	}

	return result
	
}
