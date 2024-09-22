package main

import (
	"fmt"
)

func location(name, city string) (region, continent string) {

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("Matt", "LA")
	fmt.Printf("Matt lives in %s, %s", region, continent)
}
