package flat

import (
	"reflect"
	"strconv"
	// "strings"

	// "github.com/imdario/mergo"
)

// Flat returns a flat map by taking original nested map
func Flat(nestedMap map[string]interface{}) (result map[string]interface{}, err error) {
	result, err = mapKeyVals("", nestedMap)

	return 
}

// mapKeyVals returns flatMap
func mapKeyVals(prefixKey string, nestedMap interface{}) (flatMap map[string]interface{}, err error) {
	flatMap = make(map[string]interface{})
	Delimiter := "."

	switch nestedMap := nestedMap.(type) {
		case map[string]interface{}:
			if reflect.DeepEqual(nestedMap, map[string]interface{}{}) {
				flatMap[prefixKey] = nestedMap
				return
			}

			for key, val := range nestedMap {
				newKey := key

				if prefixKey != "" {
					newKey = prefixKey + Delimiter + newKey
				}

				tempMap, mapErr := mapKeyVals(newKey, val)
				if mapErr != nil {
					err = mapErr
					return
				}

				updateMap(flatMap, tempMap)
			}
		case []interface{}:
			if reflect.DeepEqual(nestedMap, []interface{}{}) {
				flatMap[prefixKey] = nestedMap
				return
			}

			for key, val := range nestedMap {
				newKey := strconv.Itoa(key)
				if prefixKey != "" {
					newKey = prefixKey + Delimiter + newKey
				}

				tempMap, mapErr := mapKeyVals(newKey, val)
				if mapErr != nil {
					err = mapErr
					return
				}

				updateMap(flatMap, tempMap)
			}
		default:
			flatMap[prefixKey] = nestedMap
		}
		return
}

// updateMap takes a input map and return the map 
func updateMap(output map[string]interface{}, input map[string]interface{}) {
	for key, val := range input {
		output[key] = val
	}
}