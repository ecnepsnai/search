package search_test

import (
	"testing"

	"github.com/ecnepsnai/search"
)

type UserType struct {
	Username string
	Email    string
}

type WidgetType struct {
	ID   int
	Name string
}

func TestSearch(t *testing.T) {
	users := []UserType{
		{
			Username: "edwardswallow",
			Email:    "eswallow@leigon.com",
		},
		{
			Username: "roberthouse",
			Email:    "rhouse@vegas.biz",
		},
	}
	widgets := []WidgetType{
		{
			ID:   1,
			Name: "apple",
		},
		{
			ID:   2,
			Name: "orange",
		},
		{
			ID:   3,
			Name: "banana",
		},
	}

	s := search.Search{}
	for _, user := range users {
		s.Feed(user, "Username", "Email")
	}
	for _, widget := range widgets {
		s.Feed(widget, "Name")
	}

	doTest := func(query string, expectedResults int) {
		results := s.Search(query)
		actual := len(results)
		if expectedResults != actual {
			t.Errorf("Unexpected number of results for query '%s'. Expected %d got %d", query, expectedResults, actual)
		}
	}

	doTest("apple", 1)
	doTest("APPLE", 1)
	doTest("@", 2)
	doTest("a", 5)
	doTest("", 0)
}

func TestFeedNotStruct(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Did not panic when feeding slice")
		}
	}()

	s := search.Search{}
	s.Feed([]string{})
}

func TestFeedNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Did not panic when feeding nil")
		}
	}()

	s := search.Search{}
	s.Feed(nil)
	s.Search("")
}

func TestFeedStructWithoutField(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Did not panic when feeding struct without a field")
		}
	}()

	type example struct {
		foo string
	}

	s := search.Search{}
	s.Feed(example{
		foo: "",
	}, "bar")
	s.Search("")
}
