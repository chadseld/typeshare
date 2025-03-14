package proto

import (
    "encoding/json"
)

type CustomType struct {
}
type Types struct {
	S string `json:"s"`
	StaticS string `json:"static_s"`
	Int8 int `json:"int8"`
	Float float32 `json:"float"`
	Double float64 `json:"double"`
	Array []string `json:"array"`
	FixedLengthArray [4]string `json:"fixed_length_array"`
	Dictionary map[string]int `json:"dictionary"`
	OptionalDictionary *map[string]int `json:"optional_dictionary,omitempty"`
	CustomType CustomType `json:"custom_type"`
}
