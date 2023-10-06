// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/utils"
	"time"
)

type Person struct {
	Address    string     `json:"address"`
	Age        int64      `json:"age"`
	Birthday   types.Date `json:"birthday"`
	Birthtime  time.Time  `json:"birthtime"`
	Name       string     `json:"name"`
	PersonType string     `json:"personType"`
	UID        string     `json:"uid"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Person) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *Person) GetAge() int64 {
	if o == nil {
		return 0
	}
	return o.Age
}

func (o *Person) GetBirthday() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.Birthday
}

func (o *Person) GetBirthtime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Birthtime
}

func (o *Person) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Person) GetPersonType() string {
	if o == nil {
		return ""
	}
	return o.PersonType
}

func (o *Person) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}