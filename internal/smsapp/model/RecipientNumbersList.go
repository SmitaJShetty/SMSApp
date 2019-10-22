package model

import "strings"

// RecipientList construct for recipient numbers
type RecipientList struct {
	NumberList []string `json:"number_list"`
}

// NewRecipientList generates a new recipient list
func NewRecipientList(list []string) *RecipientList {
	return &RecipientList{
		NumberList: list,
	}
}

// GetCommaSepStr returns a comma separated list of moble numbers and error
func (l *RecipientList) GetCommaSepStr() string {
	if len(l.NumberList) == 0 {
		return ""
	}
	return strings.Join(l.NumberList, ",")
}
