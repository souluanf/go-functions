package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0
	for _, value := range lista {
		total += value
	}
	return total / float64(len(lista))
}

func main() {
	fmt.Println(media([]float64{98, 93, 77, 82, 83}))
}
