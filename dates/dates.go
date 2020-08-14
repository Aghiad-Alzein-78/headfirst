package dates

//DaysInWeek the number of days within a week
const DaysInWeek int = 7

//WeeksToDays convert weeks to days
func WeeksToDays(weeks int) int {
	return weeks * DaysInWeek
}

//DaysToWeeks convert days to weeks
func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
