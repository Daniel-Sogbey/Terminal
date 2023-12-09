package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom for structs and embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{data,
		errors{}}
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, fieldLabel string, r *http.Request) bool {
	x := r.Form.Get(field)

	fmt.Println("Field Label : ", fieldLabel)

	if x == "" {
		f.Errors.Add(field, fmt.Sprintf("%s cannot be empty", fieldLabel))
		return false
	}
	return true

}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		fieldValue := f.Get(field)

		if strings.TrimSpace(fieldValue) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}
