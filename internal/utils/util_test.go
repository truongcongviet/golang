package utils

import (
	"friend-management-v1/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetainSlices(t *testing.T) {
	slice1 := []string{"1", "2"}
	slice2 := []string{"2", "3"}
	actual := RetainSlices(slice1, slice2)

	expect := []string{"2"}

	if len(actual) != len(expect) {
		t.Errorf("expected slice with length: %d instead of %d", len(expect), len(actual))
	}

	if actual[0] != expect[0] {
		t.Errorf("expected slice with value: %s instead of %s", expect[0], actual[0])
	}

}

func TestGetEmailsFromText(t *testing.T) {
	slice1 := []string{"1", "2", "1"}
	actual := Unique(slice1)

	expect := []string{"1", "2"}

	if len(actual) != len(expect) {
		t.Errorf("expected slice with length: %d instead of %d", len(expect), len(actual))
	}

	if actual[0] != expect[0] {
		t.Errorf("expected slice with value: %s instead of %s", expect[0], actual[0])
	}

}

func TestUnique(t *testing.T) {
	text := "ada hau@gmail.com"
	actual := GetEmailsFromText(text)

	expect := []string{"hau@gmail.com"}

	if len(actual) != len(expect) {
		t.Errorf("expected slice with length: %d instead of %d", len(expect), len(actual))
	}

	if actual[0] != expect[0] {
		t.Errorf("expected slice with value: %s instead of %s", expect[0], actual[0])
	}

}

func TestIsEmailValidTrue(t *testing.T) {
	text := "hau@gmail.com"
	actual := IsEmailValid(text)

	assert.Equal(t, true, actual)
}

func TestIsEmailValidFalse(t *testing.T) {
	text := "ha"
	actual := IsEmailValid(text)

	assert.Equal(t, false, actual)
}

func TestValidateAddComonRequestOK(t *testing.T) {
	friends := []string{"da@gmail.com", "yas@gmail.com"}
	rq := model.AddAndGetCommonRequest{
		Friends: friends,
	}
	err := ValidateAddComonRequest(rq)

	assert.Nil(t, err)

}
func TestValidateAddComonRequestEmptyList(t *testing.T) {
	friends := []string{}
	rq := model.AddAndGetCommonRequest{
		Friends: friends,
	}
	err := ValidateAddComonRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "must contain 2 emails", err.Error())
}

func TestValidateAddComonRequestInvalidEmail(t *testing.T) {
	friends := []string{
		"quan12yt@gmail.com",
		"da",
	}
	rq := model.AddAndGetCommonRequest{
		Friends: friends,
	}
	err := ValidateAddComonRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid email format", err.Error())
}

func TestValidateSubcribeAndBlockRequestEmpty(t *testing.T) {

	rq := model.SubcribeAndBlockRequest{
		Requestor: "",
		Target:    "",
	}
	err := ValidateSubcribeAndBlockRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "requestor and target must not be empty", err.Error())
}

func TestValidateSubcribeAndBlockRequestInvalidEmail(t *testing.T) {

	rq := model.SubcribeAndBlockRequest{
		Requestor: "weq",
		Target:    "qwe",
	}
	err := ValidateSubcribeAndBlockRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid email format", err.Error())
}

func TestValidateRetrieveRequestOk(t *testing.T) {

	rq := model.RetrieveRequest{
		Sender: "qu@gmail.com",
		Text:   "qwe",
	}
	err := ValidateRetrieveRequest(rq)

	assert.Nil(t, err)
}

func TestValidateRetrieveRequestEmptySender(t *testing.T) {

	rq := model.RetrieveRequest{
		Sender: "",
		Text:   "qwe",
	}
	err := ValidateRetrieveRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "sender must not empty", err.Error())
}

func TestValidateRetrieveRequestInvalidEmail(t *testing.T) {

	rq := model.RetrieveRequest{
		Sender: "weqw",
		Text:   "qwe",
	}
	err := ValidateRetrieveRequest(rq)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid email format", err.Error())
}
