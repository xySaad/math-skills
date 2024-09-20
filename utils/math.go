package utils

import "math"

func Average(list []int) int {
	counter := 0
	sum := 0

	for _, v := range list {
		counter++
		sum += v
	}

	average := sum / counter

	return average
}

func Median(list []int) int {
	orderedList := Order(list)
	result := 0

	if len(orderedList)%2 == 0 {
		result = (orderedList[(len(orderedList)/2)-1] + orderedList[(len(orderedList)/2)]) / 2
	} else {
		result = orderedList[(len(orderedList)/2)-1]
	}

	return result
}

func Variance(list []int) int {
	mean := Average(list)
	deviations := []int{}
	for i, v := range list {
		deviations = append(deviations, v-mean)
		deviations[i] = deviations[i] * deviations[i]
	}
	result := Average(deviations)
	return result
}

func StandardDeviation(list []int) int {
	variance := Variance(list)
	return int(math.Sqrt(float64(variance)))
}
