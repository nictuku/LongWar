package config

import (
	"testing"
)

type TestNested struct {
	Val1 int
	Val2 int `json:"valother2"`
}

type TestConfigJson struct {
	Value1 string   `json:"first"`
	Array  []string `json:"array"`

	NestedInline *struct {
		Val1 string `json:"val1"`
		Val2 int    `json:"val2"`
	} `json:"nested1"`

	Nested *TestNested `json:"nested2"`

	NestedInf interface{} `json:"nested3"`
}

func TestJsonLoader(t *testing.T) {
	l := NewJsonLoader("tests/config.json")
	c := &TestConfigJson{}

	err := l.Load(c)
	if err != nil {
		t.Fatal("Failed to get string from json config, ", err)
	}

	if c.Value1 != "test" {
		t.Error("Invalid value for test json config hello key expected test got, ", c.Value1)
	}

	if len(c.Array) != 3 {
		t.Error("Invalid number of elements from test config, len=", len(c.Array))
	}

	if c.NestedInline.Val1 != "Test string" {
		t.Error("Invalid value for NestedInline.Val1,", c.NestedInline.Val1)
	}

	if c.NestedInline.Val2 != 123 {
		t.Error("Invalid value for NestedInline.Val2,", c.NestedInline.Val2)
	}

	if c.Nested.Val1 != 321 {
		t.Error("Invalid value for Nested.Val1,", c.Nested.Val1)
	}

	if c.Nested.Val2 != 111 {
		t.Error("Invalid value for Nested.Val2,", c.Nested.Val2)
	}
}
