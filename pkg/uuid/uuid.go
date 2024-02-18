package uuid

import (
	"regexp"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.New().String()
}

func InvalidUUID(uuid string) bool {
	uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)

	return uuidRegex.MatchString(uuid)
}
