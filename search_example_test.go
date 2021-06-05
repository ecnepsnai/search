package search_test

import (
	"fmt"

	"github.com/ecnepsnai/search"
)

func ExampleSearch_Search() {
	type Fruit struct {
		ID           int
		Name         string
		Translations []string
	}
	type Vegetable struct {
		ID           int
		Name         string
		Translations []string
	}

	fruits := []Fruit{
		{
			ID:           1,
			Name:         "Apple",
			Translations: []string{"Pomme", "Manzana"},
		},
		{
			ID:           2,
			Name:         "Banana",
			Translations: []string{"Banane", "Plátano"},
		},
	}
	vegetables := []Vegetable{
		{
			ID:           1,
			Name:         "Broccoli",
			Translations: []string{"Brocoli", "Brócoli"},
		},
		{
			ID:           2,
			Name:         "Carrot",
			Translations: []string{"Carotte", "Zanahoria"},
		},
	}

	s := search.Search{}
	for _, fruit := range fruits {
		s.Feed(fruit, "Name", "Translations")
	}
	for _, vegetable := range vegetables {
		s.Feed(vegetable, "Name", "Translations")
	}

	results := s.Search("z")

	for _, result := range results {
		if fruit, isFruit := result.(Fruit); isFruit {
			fmt.Printf("Fruit: id=%d name=%s\n", fruit.ID, fruit.Name)
		}
		if vegetable, isVegetable := result.(Vegetable); isVegetable {
			fmt.Printf("Vegetable: id=%d name=%s\n", vegetable.ID, vegetable.Name)
		}
	}

	// output: Vegetable: id=2 name=Carrot
	// Fruit: id=1 name=Apple
}
