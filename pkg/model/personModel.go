package model

import (
	"github.com/QOSQOs/UNIVeasier/pkg/model/errors"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"

	"regexp"
)

var lettersRegex = regexp.MustCompile(`^[a-zA-Z]+(\s*[a-zA-Z]*)+$`)
var emailsRegex = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z.\-]+\.[a-z]{2,4}$`)

type Person struct {
	Id               types.Int64  `json:"id"`
	Code             types.String `json:"code"`
	Gender           types.Gender `json:"gender"`
	FirstName        types.String `json:"first_name"`
	LastName         types.String `json:"last_name"`
	Phone            types.String `json:"phone"`
	Email            types.String `json:"email"`
	Avatar           []byte       `json:"avatar"`
	Type             types.Person `json:"type"`
	HomeCity         types.String `json:"home_city"`
	CurrentCity      types.String `json:"current_city"`
	Ethnic           types.Int32  `json:"ethnic"`
	Nationality      types.String `json:"nationality"`
	BirthDate        types.Time   `json:"birth_date"`
	AdmissionYear    types.Int32  `json:"admission_year"`
	Period           types.Int32  `json:"period"`
	IsVerified       types.Status `json:"is_verified"`
	DocVerifier      []byte       `json:"doc_verifier"`
	CreatedDate      types.Time   `json:"created_date"`
	LastModifiedDate types.Time   `json:"last_modified_date"`
}

func (m *Person) Validate() error {
	ok := lettersRegex.MatchString(m.FirstName.String)
	if len(m.FirstName.String) == 0 || !ok {
		return &errors.InvalidNameError{m.FirstName.String}
	}

	ok = lettersRegex.MatchString(m.LastName.String)
	if len(m.LastName.String) == 0 || !ok {
		return &errors.InvalidLastNameError{m.LastName.String}
	}

	ok = emailsRegex.MatchString(m.Email.String)
	if len(m.Email.String) == 0 || !ok {
		return &errors.InvalidEmailError{m.Email.String}
	}

	err := m.Gender.IsValid()
	if err != nil {
		return err
	}

	err = m.Type.IsValid()
	if err != nil {
		return err
	}

	err = m.IsVerified.IsValid()
	if err != nil {
		return err
	}
	return nil
}
