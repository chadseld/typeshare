package proto

import (
    "encoding/json"
)

type AddressDetails struct {
}
type AddressTypes string
const (
	AddressTypeVariantFixedAddress AddressTypes = "FixedAddress"
	AddressTypeVariantNoFixedAddress AddressTypes = "NoFixedAddress"
)
type Address struct{ 
	Type AddressTypes `json:"type"`
	content interface{}
}

func (a *Address) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AddressTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AddressTypeVariantFixedAddress:
		var res AddressDetails
		a.content = &res
	case AddressTypeVariantNoFixedAddress:
		return nil

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a Address) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AddressTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a Address) FixedAddress() *AddressDetails {
	res, _ := a.content.(*AddressDetails)
	return res
}

func NewAddressTypeVariantFixedAddress(content *AddressDetails) Address {
    return Address{
        Type: AddressTypeVariantFixedAddress,
        content: content,
    }
}
func NewAddressTypeVariantNoFixedAddress() Address {
    return Address{
        Type: AddressTypeVariantNoFixedAddress,
    }
}

