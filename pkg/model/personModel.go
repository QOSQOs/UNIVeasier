package model

import (
	"errors"
	"regexp"

	"github.com/QOSQOs/UNIVeasier/pkg/model/common"
	"github.com/QOSQOs/UNIVeasier/pkg/model/types"
)

var lettersRegex = regexp.MustCompile(`^[a-zA-Z]+(\s*[a-zA-Z]*)+$`)
var emailsRegex = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z.\-]+\.[a-z]{2,4}$`)

type Person struct {
	Id               types.Int64           `json:"id"`
	Code             types.String          `json:"code"`
	Gender           common.Gender         `json:"gender"`
	FirstName        types.String          `json:"first_name"`
	LastName         types.String          `json:"last_name"`
	Phone            types.String          `json:"phone"`
	Email            types.String          `json:"email"`
	Avatar           []byte                `json:"avatar"`
	Type             common.TypePerson     `json:"type"`
	HomeCity         types.String          `json:"home_city"`
	CurrentCity      types.String          `json:"current_city"`
	Ethnic           types.Int32           `json:"ethnic"`
	Nationality      types.String          `json:"nationality"`
	BirthDate        types.Time            `json:"birth_date"`
	AdmissionYear    types.Int32           `json:"admission_year"`
	Period           types.Int32           `json:"period"`
	IsVerified       common.InstanceStatus `json:"is_verified"`
	DocVerifier      []byte                `json:"doc_verifier"`
	CreatedDate      types.Time            `json:"created_date"`
	LastModifiedDate types.Time            `json:"last_modified_date"`
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
