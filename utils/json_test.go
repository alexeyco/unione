package utils_test

import (
	"reflect"
	"testing"

	"github.com/alexeyco/unione/utils"
)

type toJsonTestData struct {
	original interface{}
	result   string
}

func TestToJson(t *testing.T) {
	data := []toJsonTestData{
		{
			original: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
			result: `{"foo":{"bar":"baz"}}`,
		},
		{
			original: struct {
				Foo struct {
					Bar string `json:"bar"`
				} `json:"foo"`
			}{
				Foo: struct {
					Bar string `json:"bar"`
				}{
					Bar: "baz",
				},
			},
			result: `{"foo":{"bar":"baz"}}`,
		},
	}

	for _, row := range data {
		given, err := utils.ToJson(row.original)
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if given != row.result {
			t.Fatalf(`Result should be %s, %s given`, row.result, given)
		}
	}
}

type jsonMapTestData struct {
	original string
	result   map[string]interface{}
}

func TestJsonToMap(t *testing.T) {
	data := []jsonMapTestData{
		{
			original: `{"foo":{"bar":"baz"}}`,
			result: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": "baz",
				},
			},
		},
	}

	for _, row := range data {
		given, err := utils.JsonToMap(row.original)
		if err != nil {
			t.Fatalf(`Error should be nil, "%s" given`, err)
		}

		if !reflect.DeepEqual(row.result, given) {
			t.Fatal(`Results should be equal`)
		}
	}
}

func TestJsonIsEqual(t *testing.T) {
	expected := `{"foo":[{"bar":"baz"}]}`
	given := `{
		"foo": [
			{
				"bar":"baz"
			}
		]
	}`

	utils.JsonIsEqual(t, expected, given)
}
