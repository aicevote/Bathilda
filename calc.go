package main

import (
	"math"
)

// [2400, 7200, 21600, 64800, 194400] * 1000
func formula(val float64, meltingRate int) float64 {
	val /= float64(meltingRate)

	return (4*val + 5) / (val*val + 4*val + 5)
}

func calcResult(votes []vote, theme theme, milliseconds int) []float64 {
	sumOfPoints := 0.0
	points := make([]float64, len(theme.Choices))

	for _, vote := range votes {
		if vote.ThemeID == theme.ThemeID &&
			milliseconds >= vote.CreatedAt &&
			(vote.ExpiredAt == 0 || vote.ExpiredAt > milliseconds) {
			point := formula(float64(milliseconds-vote.CreatedAt), theme.MeltingRate)
			points[vote.Answer] += point
			sumOfPoints += point
		}
	}

	for i := 0; i < len(theme.Choices); i++ {
		points[i] = math.Round(points[i]/sumOfPoints*1000000) / 10000
		if math.IsNaN(points[i]) {
			points[i] = 0
		}
	}
	return points
}

func calcTransition(votes []vote, theme theme, now int) transition {
	var shortTransition []result
	var longTransition []result

	for i := 0; i < 60; i++ {
		timestamp := now - i*theme.SaveInterval
		percentage := calcResult(votes, theme, timestamp)

		shortTransition = append(shortTransition, result{timestamp, percentage})
	}

	for i := 0; i < 60; i++ {
		timestamp := now - i*theme.SaveInterval*24
		percentage := calcResult(votes, theme, timestamp)

		longTransition = append(longTransition, result{timestamp, percentage})
	}

	return transition{
		ShortTransition: shortTransition,
		LongTransition:  longTransition,
	}
}
