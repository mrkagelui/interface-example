package main

import "testing"

// Context: This is a part of a configureable user authentication app.
//   There are "features" that are supposed to extract different attributes of a user, which are of different data types (let's assume only bool and int for simplicity).
//   An `AllFeatures()` func is returning all supported features, every one of which is supposed to correspond to an "extractor", a function.
//   Now I want to make sure all features are implemented, i.e., each one of the features actually corresponds to an extrator func.
//   Is the test below correctly asserting that?

func Test_allFeaturesAreImplemented(t *testing.T) {
	features := AllFeatures()
	var extractor any
	for _, f := range features {
		switch f.resultDataType {
		case "bool":
			extractor = boolRegistry(f.name)
		case "int":
			extractor = intRegistry(f.name)
		default:
			t.Errorf("please test data type %v", f.resultDataType)
		}

		if extractor == nil {
			t.Errorf("please implement feature %v", f.name)
		}

	}
}

func AllFeatures() []feature {
	return []feature{
		{
			name:           "is_bot",
			resultDataType: "bool",
		},
		{
			name:           "age",
			resultDataType: "int",
		},
	}
}

type user struct {
	name string
	age  int
}

type feature struct {
	resultDataType string
	name           string
}

type boolExtractor func(user) bool
type intExtractor func(user) int

func boolRegistry(featureName string) boolExtractor {
	switch featureName {
	case "is_bot":
		return isBot
	default:
		return nil
	}
}

func intRegistry(featureName string) intExtractor {
	switch featureName {
	case "age":
		return age
	default:
		return nil
	}
}

func isBot(u user) bool {
	return u.name == "bot"
}

func age(u user) int {
	return u.age
}
