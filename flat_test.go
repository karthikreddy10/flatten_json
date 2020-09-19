package flat

import (
	"encoding/json"
	"reflect"
	"testing"
	"io/ioutil"
	"fmt"
)

func TestFlat(t *testing.T) {
	tests := []struct{
		fixtureName string
		expected map[string]interface{} 
	}{
		{
			"simple",
			map[string]interface{}{"hello": "world"},
		},
		{
			"simple_nested",
			map[string]interface{}{"user.details.name": "karthik"},
		},
		{
			"deep_nested",
			map[string]interface{}{"user.details.name": "karthik", "user.details.location.coordinates.lat": 34.5, "user.details.location.coordinates.long": 21.6},
		},
		{
			"simple_array",
			map[string]interface{}{"users.0": "pogba", "users.1": "bruno", "users.2": "van de beek"},
		},
		{
			"nested_array",
			map[string]interface{}{"club.name": "manchester", "club.location.coordinates.lat": 34.5, "club.location.coordinates.long": 21.6, "club.squad.players.0": "pogba", "club.squad.players.1": "bruno", "club.squad.players.2": "van de beek"},
		},
	}

	for testCase, test := range tests {
		jsonData, err := getJSONFileData(t, test.fixtureName)
		if err != nil {
			t.Errorf("%d: failed to get JSON data: %v", testCase+1, err)
		}
		
		flatMap, err := Flat(jsonData)
		if err != nil {
			t.Errorf("%d: failed to flatten map: %v", testCase+1, err)
		}

		if !reflect.DeepEqual(flatMap, test.expected) {
			t.Errorf("%s: name, %d: mismatch, got: %v want: %v",test.fixtureName, testCase+1, flatMap, test.expected)
		}
	}
}

func getJSONFileData(t *testing.T, fixtureName string) (result map[string]interface{}, err error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("./fixtures/%s.json", fixtureName))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &result) 
	if err != nil {
		return nil, err
	}

	return result, nil
}