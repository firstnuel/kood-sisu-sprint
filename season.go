package sprint

import "fmt"

func Season(month string) string {
	season := fmt.Sprintf("invalid input: %v", month)

	switch month {
	case "jan", "feb", "dec":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	}
	return season
}
