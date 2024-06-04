package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func NewPrefixedGuid(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, uuid.New().String())
}

func NewSuffixedGuid(suffix string) string {
	return fmt.Sprintf("%s%s", uuid.New().String(), suffix)
}
