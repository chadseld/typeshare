package proto

import "encoding/json"

// Struct comment
type ItemDetailsFieldValue struct {
}
// Enum comment
type AdvancedColorsTypes string
const (
	// This is a case comment
	AdvancedColorsTypeVariantString AdvancedColorsTypes = "String"
	AdvancedColorsTypeVariantNumber AdvancedColorsTypes = "Number"
	AdvancedColorsTypeVariantUnsignedNumber AdvancedColorsTypes = "UnsignedNumber"
	AdvancedColorsTypeVariantNumberArray AdvancedColorsTypes = "NumberArray"
	// Comment on the last element
	AdvancedColorsTypeVariantReallyCoolType AdvancedColorsTypes = "ReallyCoolType"
)
type AdvancedColors struct{ 
	Type AdvancedColorsTypes `json:"type"`
	content interface{}
}

func (a *AdvancedColors) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AdvancedColorsTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AdvancedColorsTypeVariantString:
		var res string
		a.content = &res
	case AdvancedColorsTypeVariantNumber:
		var res int
		a.content = &res
	case AdvancedColorsTypeVariantUnsignedNumber:
		var res uint32
		a.content = &res
	case AdvancedColorsTypeVariantNumberArray:
		var res []int
		a.content = &res
	case AdvancedColorsTypeVariantReallyCoolType:
		var res ItemDetailsFieldValue
		a.content = &res

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a AdvancedColors) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AdvancedColorsTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a AdvancedColors) String() string {
	res, _ := a.content.(*string)
	return *res
}
func (a AdvancedColors) Number() int {
	res, _ := a.content.(*int)
	return *res
}
func (a AdvancedColors) UnsignedNumber() uint32 {
	res, _ := a.content.(*uint32)
	return *res
}
func (a AdvancedColors) NumberArray() []int {
	res, _ := a.content.(*[]int)
	return *res
}
func (a AdvancedColors) ReallyCoolType() *ItemDetailsFieldValue {
	res, _ := a.content.(*ItemDetailsFieldValue)
	return res
}

func NewAdvancedColorsTypeVariantString(content string) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantString,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantNumber(content int) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantNumber,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantUnsignedNumber(content uint32) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantUnsignedNumber,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantNumberArray(content []int) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantNumberArray,
        content: &content,
    }
}
func NewAdvancedColorsTypeVariantReallyCoolType(content *ItemDetailsFieldValue) AdvancedColors {
    return AdvancedColors{
        Type: AdvancedColorsTypeVariantReallyCoolType,
        content: content,
    }
}

type AdvancedColors2Types string
const (
	// This is a case comment
	AdvancedColors2TypeVariantString AdvancedColors2Types = "string"
	AdvancedColors2TypeVariantNumber AdvancedColors2Types = "number"
	AdvancedColors2TypeVariantNumberArray AdvancedColors2Types = "number-array"
	// Comment on the last element
	AdvancedColors2TypeVariantReallyCoolType AdvancedColors2Types = "really-cool-type"
)
type AdvancedColors2 struct{ 
	Type AdvancedColors2Types `json:"type"`
	content interface{}
}

func (a *AdvancedColors2) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AdvancedColors2Types   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AdvancedColors2TypeVariantString:
		var res string
		a.content = &res
	case AdvancedColors2TypeVariantNumber:
		var res int
		a.content = &res
	case AdvancedColors2TypeVariantNumberArray:
		var res []int
		a.content = &res
	case AdvancedColors2TypeVariantReallyCoolType:
		var res ItemDetailsFieldValue
		a.content = &res

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a AdvancedColors2) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AdvancedColors2Types   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a AdvancedColors2) String() string {
	res, _ := a.content.(*string)
	return *res
}
func (a AdvancedColors2) Number() int {
	res, _ := a.content.(*int)
	return *res
}
func (a AdvancedColors2) NumberArray() []int {
	res, _ := a.content.(*[]int)
	return *res
}
func (a AdvancedColors2) ReallyCoolType() *ItemDetailsFieldValue {
	res, _ := a.content.(*ItemDetailsFieldValue)
	return res
}

func NewAdvancedColors2TypeVariantString(content string) AdvancedColors2 {
    return AdvancedColors2{
        Type: AdvancedColors2TypeVariantString,
        content: &content,
    }
}
func NewAdvancedColors2TypeVariantNumber(content int) AdvancedColors2 {
    return AdvancedColors2{
        Type: AdvancedColors2TypeVariantNumber,
        content: &content,
    }
}
func NewAdvancedColors2TypeVariantNumberArray(content []int) AdvancedColors2 {
    return AdvancedColors2{
        Type: AdvancedColors2TypeVariantNumberArray,
        content: &content,
    }
}
func NewAdvancedColors2TypeVariantReallyCoolType(content *ItemDetailsFieldValue) AdvancedColors2 {
    return AdvancedColors2{
        Type: AdvancedColors2TypeVariantReallyCoolType,
        content: content,
    }
}

