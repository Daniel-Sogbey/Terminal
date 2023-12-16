package forms

import (
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors{},
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		fieldValue := f.Get(field)

		if strings.TrimSpace(fieldValue) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

func (f *Form) IsUrl(field string) {

	if !govalidator.IsURL(f.Get(field)) {
		f.Errors.Add(field, "Please provide a valid URL.")
	}
}
