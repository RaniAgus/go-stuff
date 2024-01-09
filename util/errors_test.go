package util_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RaniAgus/go-starter/util"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestGetValidationErrorFields(t *testing.T) {
	var tests = []struct {
		name   string
		errors error
		fields []string
	}{
		{
			name: "should return a slice of fields with validation errors",
			errors: validator.ValidationErrors{
				FakeFieldError{field: "field1"},
				FakeFieldError{field: "field2"},
			},
			fields: []string{"field1", "field2"},
		},
		{
			name: "should return a slice of fields with validation errors, removing the key from the field name if the error is a slice or map",
			errors: validator.ValidationErrors{
				FakeFieldError{field: "field1[0]"},
				FakeFieldError{field: "field2[1]"},
			},
			fields: []string{"field1", "field2"},
		},
		{
			name:   "should return an empty slice if there are no validation errors",
			errors: validator.ValidationErrors{},
			fields: []string{},
		},
		{
			name:   "should return an empty slice if the error is not a slice of validation errors",
			errors: util.APIError{},
			fields: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fields := util.GetValidationErrorFields(test.errors)

			if len(fields) != len(test.fields) {
				t.Errorf("Expected length of fields to be %d, got %d", len(test.fields), len(fields))
			}

			for i, field := range fields {
				if field != test.fields[i] {
					t.Errorf("Expected field %s to be in the slice, got %s", test.fields[i], field)
				}
			}
		})
	}
}

type FakeFieldError struct {
	field string
}

func (err FakeFieldError) Tag() string { return "" }

func (err FakeFieldError) ActualTag() string { return "" }

func (err FakeFieldError) Namespace() string { return "" }

func (err FakeFieldError) StructNamespace() string { return "" }

func (err FakeFieldError) Field() string { return err.field }

func (err FakeFieldError) StructField() string { return "" }

func (err FakeFieldError) Value() interface{} { return "" }

func (err FakeFieldError) Param() string { return "" }

func (err FakeFieldError) Kind() reflect.Kind { return reflect.String }

func (err FakeFieldError) Type() reflect.Type { return reflect.TypeOf("") }

func (err FakeFieldError) Translate(translator ut.Translator) string { return "" }

func (err FakeFieldError) Error() string { return fmt.Sprintf("Field %s is invalid", err.field) }
