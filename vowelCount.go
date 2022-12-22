package main

import "fmt"

func countVowels(s string) (int, map[rune]int) {
	vowelDistribution := make(map[rune]int)

	numVowels := 0

	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			numVowels++

			vowelDistribution[r]++
		}
	}

	return numVowels, vowelDistribution
}
func main() {
	numVowels, vowelDistribution := countVowels("I am learning GoLang")

	fmt.Println("Number of vowels:", numVowels)
	fmt.Println("Vowel distribution:", vowelDistribution)

}
