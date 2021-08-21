package utils

import (
	"errors"
	"friend-management-v1/model"
	"regexp"
)

func RetainSlices(slice1 []string, slice2 []string) []string {
	results := make([]string, 0) // slice tostore the result

	for i := 0; i < len(slice1); i++ {
		for k := 0; k < len(slice2); k++ {
			if slice1[i] != slice2[k] {
				continue
			}
			results = append(results, slice1[i])
		}
	}
	return results
}

func GetEmailsFromText(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(text, -1)

	return match
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 || len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func ValidateAddComonRequest(rq model.AddAndGetCommonRequest) error {
	// if rq.Friends == nil {
	// 	return errors.New("friends list must not be null")
	// }
	if len(rq.Friends) != 2 {
		return errors.New("must contain 2 emails")
	}
	if !IsEmailValid(rq.Friends[0]) || !IsEmailValid(rq.Friends[1]) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateSubcribeAndBlockRequest(rq model.SubcribeAndBlockRequest) error {
	if rq.Requestor == "" || rq.Target == "" {
		return errors.New("requestor and target must not be empty")
	}
	if !IsEmailValid(rq.Requestor) || !IsEmailValid(rq.Target) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateRetrieveRequest(rq model.RetrieveRequest) error {
	if rq.Sender == "" {
		return errors.New("sender must not empty")
	}
	if !IsEmailValid(rq.Sender) {
		return errors.New("invalid email format")
	}
	return nil
}
