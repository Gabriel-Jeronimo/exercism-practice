package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(CalculateWorkingCarsPerMinute(productionRate, successRate)) * 60
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var sum int = 0

	sum += carsCount / 10 * 95000
	sum += carsCount % 10 * 10000

	return uint(sum)
}
