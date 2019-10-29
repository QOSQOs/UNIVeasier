package model

import (
	"errors"
	"regexp"

	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
)

var lettersRegex = regexp.MustCompile(`^[a-zA-Z]+(\s*[a-zA-Z]*)+$`)
var emailsRegex = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z.\-]+\.[a-z]{2,4}$`)

type Person struct {
	Id               common.NullInt64      `json:"id"`
	Code             common.NullString     `json:"code"`
	Gender           common.Gender         `json:"gender"`
	FirstName        common.NullString     `json:"first_name"`
	LastName         common.NullString     `json:"last_name"`
	Phone            common.NullString     `json:"phone"`
	Email            common.NullString     `json:"email"`
	Avatar           []byte                `json:"avatar"`
	Type             common.TypePerson     `json:"type"`
	HomeCity         common.NullString     `json:"home_city"`
	CurrentCity      common.NullString     `json:"current_city"`
	Ethnic           common.NullInt32      `json:"ethnic"`
	Nationality      common.NullString     `json:"nationality"`
	BirthDate        common.NullTime       `json:"birth_date"`
	AdmissionYear    common.NullInt32      `json:"admission_year"`
	Period           common.NullInt32      `json:"period"`
	IsVerified       common.InstanceStatus `json:"is_verified"`
	DocVerifier      []byte                `json:"doc_verifier"`
	CreatedDate      common.NullTime       `json:"created_date"`
	LastModifiedDate common.NullTime       `json:"last_modified_date"`
}

func (t *Person) Validate() error {
	ok := lettersRegex.MatchString(t.FirstName.String)
	if len(t.FirstName.String) == 0 || !ok {
		return errors.New("Invalid first name")
	}

	ok = lettersRegex.MatchString(t.LastName.String)
	if len(t.LastName.String) == 0 || !ok {
		return errors.New("Invalid last name")
	}

	ok = emailsRegex.MatchString(t.Email.String)
	if len(t.Email.String) == 0 || !ok {
		return errors.New("Invalid email")
	}

	err := t.Gender.IsValid()
	if err != nil {
		return err
	}

	err = t.Type.IsValid()
	if err != nil {
		return err
	}

	err = t.IsVerified.IsValid()
	if err != nil {
		return err
	}
	return nil
}
