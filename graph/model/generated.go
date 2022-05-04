// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Meta data about the service
type Info struct {
	// The name of the service
	Name string `json:"name"`
	// The currently running version in semver.
	Version string `json:"version"`
	// Are things up and working? Useful for status pages.
	Running bool `json:"running"`
}

// All known genders in the galaxy.
type Gender string

const (
	// No gender? You're probably a droid.
	GenderNone Gender = "None"
	// Male, men, man, boy... You know, dudes!
	GenderMale Gender = "Male"
	// Female, women, woman, girl... You know, ladies!
	GenderFemale Gender = "Female"
	// Both!
	GenderHermaphrodite Gender = "Hermaphrodite"
)

var AllGender = []Gender{
	GenderNone,
	GenderMale,
	GenderFemale,
	GenderHermaphrodite,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderNone, GenderMale, GenderFemale, GenderHermaphrodite:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
