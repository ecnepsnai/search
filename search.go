/*
Package search provides a mechanism for searching fields within structures using the Levenshtein distance method,
commonly referred to as "fuzzy" searching.
*/
package search

import (
	"reflect"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type object struct {
	o      interface{}
	fields []string
}

type rankedObject struct {
	o    interface{}
	rank int
}

func (o object) rank(query string) int {
	highestRank := -1
	for _, field := range o.fields {
		for _, value := range valuesForField(o, field) {
			rank := fuzzy.RankMatch(query, value)
			if rank > highestRank {
				highestRank = rank
			}
		}
	}
	return highestRank
}

func valuesForField(o object, fieldName string) []string {
	field := reflect.ValueOf(o.o).FieldByName(fieldName)

	if strSlice, isStrSlice := field.Interface().([]string); isStrSlice {
		lc := make([]string, len(strSlice))
		for i, s := range strSlice {
			lc[i] = strings.ToLower(s)
		}
		return lc
	}

	return []string{strings.ToLower(field.String())}
}

// Search describes a search instance with loaded objects. Multiple individual queries can be performed against a single
// search instance.
type Search struct {
	objects []object
}

// Feed will load searchable data into the instance. Not all objects fed into the search need to be of the same type,
// However `o` must be a struct which must contain each field specified, otherwise it will panic. Specified fields
// should be a string or a []string, other types may not work as expected.
func (s *Search) Feed(o interface{}, fields ...string) {
	value := reflect.ValueOf(o)
	if value.Kind() != reflect.Struct {
		panic("can only feed struct type to search")
	}
	for _, field := range fields {
		if !value.FieldByName(field).IsValid() {
			panic("field doesn't exist " + field)
		}
	}

	s.objects = append(s.objects, object{
		o:      o,
		fields: fields,
	})
}

// Search will query the data and return a ordered list of matching objects, or an empty slice. Searches are
// case-insensitive.
func (s *Search) Search(query string) []interface{} {
	ranks := []rankedObject{}
	for _, o := range s.objects {
		distance := o.rank(strings.ToLower(query))
		if distance < 0 {
			continue
		}
		ranks = append(ranks, rankedObject{
			o:    o.o,
			rank: distance,
		})
	}
	if len(ranks) == 0 {
		return []interface{}{}
	}

	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].rank > ranks[j].rank
	})

	results := make([]interface{}, len(ranks))
	for i, r := range ranks {
		results[i] = r.o
	}
	return results
}
