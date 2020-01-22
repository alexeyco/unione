package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

// ToJson marshals any value to json string.
func ToJson(v interface{}) (s string, err error) {
	var b []byte
	if b, err = json.Marshal(v); err != nil {
		return
	}

	s = string(b)

	return
}

// JsonToMap parses json string and returns map.
func JsonToMap(s string) (m map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(s), &m)
	return
}

// JsonIsEqual checks if two json strings are equal.
func JsonIsEqual(t *testing.T, expectedJson, givenJson string) {
	var expectedMap map[string]interface{}
	var err error
	if expectedMap, err = JsonToMap(expectedJson); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
		return
	}

	var givenMap map[string]interface{}
	if givenMap, err = JsonToMap(givenJson); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
		return
	}

	if !reflect.DeepEqual(expectedMap, givenMap) {
		t.Errorf(`JSON should be %s, %s given`, expectedJson, givenJson)
	}
}
