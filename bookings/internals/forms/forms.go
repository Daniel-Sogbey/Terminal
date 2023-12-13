package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	if x == "" {
		return false
	}
	return true

}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		fieldValue := f.Get(field)

		if strings.TrimSpace(fieldValue) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks the minimum length of a field
func (f *Form) MinLength(field string, length int) bool {
	value := f.Get(field)

	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("Minimum number of character is %d", length))
		return false
	}

	return true
}

// IsPhone checks for a valid phone number
func (f *Form) IsPhone(field string, length int) bool {
	value := f.Get(field)

	if len(value) < length || len(value) > length {
		f.Errors.Add(field, fmt.Sprintf("Enter a valid phone number of length %d", length))
		return false
	}
	return true
}

// IsEmail checks for a valid email address
func (f *Form) IsEmail(field string) bool {
	value := f.Get(field)
	if !govalidator.IsEmail(value) {
		f.Errors.Add(field, "Invalid email address")
		return false
	}
	return true
}
