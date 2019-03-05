package union

import (
	"encoding/json"
	"errors"
	"reflect"
)

// MemberUnion returns the Union of two structs or maps, including the union of any sub slices
func MemberUnion(s1, s2 interface{}) (interface{}, error) {
	var map1 map[string]interface{}
	var map2 map[string]interface{}
	var result = make(map[string]interface{})

	// convert s1 to map
	bytes1, err := json.Marshal(s1)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes1, &map1)
	if err != nil {
		return result, err
	}

	// convert s2 to map
	bytes2, err := json.Marshal(s2)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes2, &map2)
	if err != nil {
		return result, err
	}

	// iterate through map1, putting results into result map
	for k, v := range map1 {

		// if value is nil or key is _id, continue
		if v == nil || k == "_id" {
			continue
		}

		// perform slice union if v is a slice and also located in map2
		if _, ok := map2[k]; ok && map2[k] != nil && reflect.TypeOf(v).Kind() == reflect.Slice {
			slice1, err := InterfaceSlice(v)
			if err != nil {
				return result, err
			}
			slice2, err := InterfaceSlice(map2[k])
			if err != nil {
				return result, err
			}
			result[k] = SliceUnion(slice1, slice2)
		} else {
			result[k] = v
		}
	}

	// iterate through map2, putting results into result map
	for k, v := range map2 {

		// if value is nil or key is _id, continue
		if v == nil || k == "_id" {
			continue
		}

		// only add k,v to result if k is not a slice and is a different value
		// or not in result
		if _, ok := result[k]; !ok || (reflect.TypeOf(v).Kind() != reflect.Slice) {
			result[k] = v
		}
	}

	// return resulting merge
	return result, nil
}

// SliceUnion returns the union of two slices
func SliceUnion(a, b []interface{}) []interface{} {
	m := make(map[interface{}]bool)

	// iterate through slice a, adding values as
	// keys in m
	for _, v := range a {
		m[v] = true
	}

	// iterate through slice b, adding values not
	// in map m to slice a
	for _, v := range b {
		if _, ok := m[v]; !ok {
			a = append(a, v)
		}
	}

	// return union of slices a and b
	return a
}

// InterfaceSlice turns an interface{} or typed slice into a []interface{}
func InterfaceSlice(slice interface{}) ([]interface{}, error) {

	// ensure slice is of kind slice
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}

	// allocate result []interface
	result := make([]interface{}, v.Len())

	// iterate through slice, setting value in result
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}

	// return result and no error
	return result, nil
}
