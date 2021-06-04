package search_test

import (
	"fmt"

	"github.com/ecnepsnai/search"
)

func ExampleSearch_Search() {
	type Fruit struct {
		ID   int
		Name string
	}
	type Vegetable struct {
		ID   int
		Name string
	}

	fruits := []Fruit{
		{
			ID:   1,
			Name: "Apple",
		},
		{
			ID:   2,
			Name: "Orange",
		},
		{
			ID:   3,
			Name: "Banana",
		},
		{
			ID:   4,
			Name: "Tomato",
		},
	}
	vegetables := []Vegetable{
		{
			ID:   1,
			Name: "Broccoli",
		},
		{
			ID:   2,
			Name: "Carrot",
		},
		{
			ID:   3,
			Name: "Coiflour",
		},
		{
			ID:   4,
			Name: "Pepper",
		},
	}

	s := search.Search{}
	for _, fruit := range fruits {
		s.Feed(fruit, "Name")
	}
	for _, vegetable := range vegetables {
		s.Feed(vegetable, "Name")
	}

	results := s.Search("B")

	for _, result := range results {
		if fruit, isFruit := result.(Fruit); isFruit {
			fmt.Printf("Fruit: id=%d name=%s\n", fruit.ID, fruit.Name)
		}
		if vegetable, isVegetable := result.(Vegetable); isVegetable {
			fmt.Printf("Vegetable: id=%d name=%s\n", vegetable.ID, vegetable.Name)
		}
	}

	// output: Vegetable: id=1 name=Broccoli
	// Fruit: id=3 name=Banana
}
