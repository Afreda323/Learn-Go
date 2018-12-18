package main

import "strings"

func isValidString(s string) bool {
	return len(s) > 0
}

func isFruit(s string) bool {
	fruits := []string{"pear", "apple", "banana"}
	for _, fruit := range fruits {
		if strings.ToLower(s) == fruit {
			return true
		}
	}
	return false
}

func isGender(g int) bool {
	return g == 2 || g == 1
}

func isSport(s string) bool {
	sports := []string{"football", "basketball", "tennis"}
	for _, sport := range sports {
		if strings.ToLower(s) == sport {
			return true
		}
	}
	return false
}
