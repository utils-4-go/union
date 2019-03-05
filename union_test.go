package union

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Test struct {
	Name    string
	Hobbies []string
}

func TestMemberUnion(t *testing.T) {
	a := Test{
		"Test",
		[]string{"hobby1", "hobby2"},
	}
	b := Test{
		"Test2",
		[]string{"hobby3", "hobby4"},
	}
	expected := Test{
		"Test2",
		[]string{"hobby1", "hobby2", "hobby3", "hobby4"},
	}
	m, err := MemberUnion(a, b)
	result := Test{}
	bytes1, err := json.Marshal(m)
	err = json.Unmarshal(bytes1, &result)
	if err != nil {
		t.Error("Received an error", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestSliceUnion(t *testing.T) {
	a := make([]interface{}, 2)
	a[0] = "one"
	a[1] = "two"
	b := make([]interface{}, 2)
	b[0] = 2
	b[1] = 3
	expected := make([]interface{}, 4)
	expected[0] = "one"
	expected[1] = "two"
	expected[2] = 2
	expected[3] = 3
	result := SliceUnion(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}

func TestInterfaceSlice(t *testing.T) {
	a := []string{"hello", "test"}
	expected := make([]interface{}, 2)
	expected[0] = "hello"
	expected[1] = "test"
	result, err := InterfaceSlice(a)
	if err != nil {
		t.Error("Received an error", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Error("Expected", expected, "got", result)
	}
}
