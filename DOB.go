package main

import (
	"fmt"
	"time"
)

type FamilyMember struct {
	Name     string
	Relation string
	DOB      time.Time
	Age      int
}

func (f *FamilyMember) CalculateAge() {
	now := time.Now()
	f.Age = now.Year() - f.DOB.Year()
	if now.YearDay() < f.DOB.YearDay() {
		f.Age--
	}
}
func main() {
	dad := FamilyMember{
		Name:     "John",
		Relation: "Father",
		DOB:      time.Date(1975, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	dad.CalculateAge()

	mom := FamilyMember{
		Name:     "Jane",
		Relation: "Mother",
		DOB:      time.Date(1978, time.March, 15, 0, 0, 0, 0, time.UTC),
	}
	mom.CalculateAge()

	sister := FamilyMember{
		Name:     "Samantha",
		Relation: "Sister",
		DOB:      time.Date(2001, time.December, 31, 0, 0, 0, 0, time.UTC),
	}
	sister.CalculateAge()

	fmt.Printf("My dad is %d years old.\n", dad.Age)
	fmt.Printf("My mom is %d years old.\n", mom.Age)
	fmt.Printf("My sister is %d years old.\n", sister.Age)

}
