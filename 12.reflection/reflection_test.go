package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "struct with one string field",
			Input:         struct{ Name string }{"Alexis"},
			ExpectedCalls: []string{"Alexis"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Alexis", "Salvador"},
			ExpectedCalls: []string{"Alexis", "Salvador"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Alexis", 34},
			ExpectedCalls: []string{"Alexis"},
		},
		{
			Name:          "nested fields",
			Input:         Person{"Alexis", Profile{34, "Salvador"}},
			ExpectedCalls: []string{"Alexis", "Salvador"},
		},
		{
			Name:          "pointers to things",
			Input:         &Person{"Alexis", Profile{33, "Salvador"}},
			ExpectedCalls: []string{"Alexis", "Salvador"},
		},
		{
			Name:          "slices",
			Input:         []Profile{{34, "Salvador"}, {33, "London"}},
			ExpectedCalls: []string{"Salvador", "London"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
