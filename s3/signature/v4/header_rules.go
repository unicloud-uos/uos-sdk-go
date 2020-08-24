package v4

import (
	"net/http"
	"strings"
)

type Rule interface {
	IsValid(value string) bool
}

type Rules []Rule

func (r Rules) IsValid(value string) bool {
	for _, rule := range r {
		if rule.IsValid(value) {
			return true
		}
	}
	return false
}

// MapRule generic Rule for maps
type MapRule map[string]struct{}

// IsValid for the map Rule satisfies whether it exists in the map
func (m MapRule) IsValid(value string) bool {
	_, ok := m[value]
	return ok
}

type List struct {
	Rule
}

func (b List) IsValid(value string) bool {
	return !b.Rule.IsValid(value)
}

// Patterns is a list of strings to match against
type Prefix []string

func (p Prefix) IsValid(value string) bool {
	for _, pattern := range p {
		if strings.HasPrefix(http.CanonicalHeaderKey(value), pattern) {
			return true
		}
	}
	return false
}
