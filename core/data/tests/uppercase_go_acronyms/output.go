package proto

import (
    "encoding/json"
)

type AccountID string

type Foo struct {
	ID string `json:"id"`
	IDWithSuffix string `json:"id_with_suffix"`
	PrefixWithID string `json:"prefix_with_id"`
	Identity string `json:"identity"`
	LowercaseInputURL string `json:"lowercase_input_url"`
	UppercaseType AccountID `json:"uppercase_type"`
}
type BarTypes string
const (
	BarTypeVariantID BarTypes = "Id"
)
type Bar struct{ 
	Type BarTypes `json:"type"`
	content interface{}
}

func (b *Bar) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    BarTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	b.Type = enum.Tag
	switch b.Type {
	case BarTypeVariantID:
		var res string
		b.content = &res

	}
	if err := json.Unmarshal(enum.Content, &b.content); err != nil {
		return err
	}

	return nil
}

func (b Bar) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    BarTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = b.Type
    enum.Content = b.content
    return json.Marshal(enum)
}

func (b Bar) ID() string {
	res, _ := b.content.(*string)
	return *res
}

func NewBarTypeVariantID(content string) Bar {
    return Bar{
        Type: BarTypeVariantID,
        content: &content,
    }
}

