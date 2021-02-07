package domain

import (
	"fmt"
	"regexp"
)

type videoID struct {
	Value string
}

func newVideoID(uuid string) (*videoID, error) {

	if !isValidUUID(uuid) {
		return nil, fmt.Errorf("UUID (%s)its not valid", uuid)
	}

	vid := &videoID{
		Value: uuid,
	}

	return vid, nil
}

func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
