package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages map[string]int
	fmt.Println(ages)
	ages = make(map[string]int)
	fmt.Println(ages)
	ages["yogesh"] = 30
	fmt.Println(ages)
	delete(ages, "yogesh")
	fmt.Println(ages)

	temperatures := map[string]int{
		"mumbai":  30,
		"delhi":   35,
		"pune":    28,
		"chennai": 38,
	}
	fmt.Println(temperatures)

	for key, value := range temperatures {
		fmt.Println(key, value)
	}

	//sort
	cities := make([]string, 0)
	for key := range temperatures {
		cities = append(cities, key)
	}
	sort.Strings(cities)
	fmt.Println("sorted cities:", cities)
	for _, city := range cities {
		fmt.Printf("%s has temparature %d\n", city, temperatures[city])
	}

	if mumTemp, ok := temperatures["mumbaii"]; ok {
		fmt.Println(mumTemp)
	} else {
		fmt.Println("temperatures does not contain mumbai")
	}
}
