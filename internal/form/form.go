package form

import (
	"GoInActionAssignment/internal/database"
	"fmt"
	"net/url"
	"regexp"
)

// Form struct
type Form struct {
	url.Values
	Errors errors
}
type errors map[string][]string

// New creates new Form instance
func New(data url.Values) *Form {
	return &Form{data, make(errors)}
}

// Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	if _, ok := e[field]; !ok || e == nil {
		return ""
	}
	return e[field][0]
}

//Input validation of required fields
//Checks for if fields are empty for firstname, lastname.
//Username should be at least 8 characters long
//Password should not be empty
func (f *Form) Required(fields ...string) {
	firstNameRegex := "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	re, _ := regexp.Compile(firstNameRegex)
	for i := 0; i < len(fields); i++ {
		value := f.Get(fields[i])
		if i == 0 || i == 1 {
			if len(value) > 0 && len(value) < 40 && re.MatchString(value) {
				return
			} else {
				f.Errors.Add(fields[i], fmt.Sprintf("%s should not be empty", fields[i]))
			}
		} else if i == 2 {
			if len(value) < 8 {
				f.Errors.Add(fields[i], fmt.Sprintf("%s should be at least 8 characters long", fields[i]))
			}
		}
		if i == 3 {
			if value == "" {
				f.Errors.Add(fields[i], fmt.Sprintf("%s should not be empty", fields[i]))
			}
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

//Check if User entered an Existing user
//Returns False if the name exists in
//Bool default value is false in Golang
func (f *Form) ExistingUser() bool {
	username := f.Get("username")
	_, ok := database.Users[username]
	return ok
}
