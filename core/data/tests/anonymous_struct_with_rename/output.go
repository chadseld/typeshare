package proto

import "encoding/json"

// Generated type representing the anonymous struct variant `List` of the `AnonymousStructWithRename` Rust enum
type AnonymousStructWithRenameListInner struct {
	List []string `json:"list"`
}
// Generated type representing the anonymous struct variant `LongFieldNames` of the `AnonymousStructWithRename` Rust enum
type AnonymousStructWithRenameLongFieldNamesInner struct {
	SomeLongFieldName string `json:"some_long_field_name"`
	And bool `json:"and"`
	ButOneMore []string `json:"but_one_more"`
}
// Generated type representing the anonymous struct variant `KebabCase` of the `AnonymousStructWithRename` Rust enum
type AnonymousStructWithRenameKebabCaseInner struct {
	AnotherList []string `json:"another-list"`
	CamelCaseStringField string `json:"camelCaseStringField"`
	SomethingElse bool `json:"something-else"`
}
type AnonymousStructWithRenameTypes string
const (
	AnonymousStructWithRenameTypeVariantList AnonymousStructWithRenameTypes = "list"
	AnonymousStructWithRenameTypeVariantLongFieldNames AnonymousStructWithRenameTypes = "longFieldNames"
	AnonymousStructWithRenameTypeVariantKebabCase AnonymousStructWithRenameTypes = "kebabCase"
)
type AnonymousStructWithRename struct{ 
	Type AnonymousStructWithRenameTypes `json:"type"`
	content interface{}
}

func (a *AnonymousStructWithRename) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AnonymousStructWithRenameTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AnonymousStructWithRenameTypeVariantList:
		var res AnonymousStructWithRenameListInner
		a.content = &res
	case AnonymousStructWithRenameTypeVariantLongFieldNames:
		var res AnonymousStructWithRenameLongFieldNamesInner
		a.content = &res
	case AnonymousStructWithRenameTypeVariantKebabCase:
		var res AnonymousStructWithRenameKebabCaseInner
		a.content = &res

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a AnonymousStructWithRename) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AnonymousStructWithRenameTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a AnonymousStructWithRename) List() *AnonymousStructWithRenameListInner {
	res, _ := a.content.(*AnonymousStructWithRenameListInner)
	return res
}
func (a AnonymousStructWithRename) LongFieldNames() *AnonymousStructWithRenameLongFieldNamesInner {
	res, _ := a.content.(*AnonymousStructWithRenameLongFieldNamesInner)
	return res
}
func (a AnonymousStructWithRename) KebabCase() *AnonymousStructWithRenameKebabCaseInner {
	res, _ := a.content.(*AnonymousStructWithRenameKebabCaseInner)
	return res
}

func NewAnonymousStructWithRenameTypeVariantList(content *AnonymousStructWithRenameListInner) AnonymousStructWithRename {
    return AnonymousStructWithRename{
        Type: AnonymousStructWithRenameTypeVariantList,
        content: content,
    }
}
func NewAnonymousStructWithRenameTypeVariantLongFieldNames(content *AnonymousStructWithRenameLongFieldNamesInner) AnonymousStructWithRename {
    return AnonymousStructWithRename{
        Type: AnonymousStructWithRenameTypeVariantLongFieldNames,
        content: content,
    }
}
func NewAnonymousStructWithRenameTypeVariantKebabCase(content *AnonymousStructWithRenameKebabCaseInner) AnonymousStructWithRename {
    return AnonymousStructWithRename{
        Type: AnonymousStructWithRenameTypeVariantKebabCase,
        content: content,
    }
}

