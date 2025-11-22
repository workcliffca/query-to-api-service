package utils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrNotSelectQuery     = errors.New("only SELECT queries are allowed")
	ErrForbiddenKeyword   = errors.New("query contains forbidden keywords")
	ErrInvalidQueryFormat = errors.New("invalid query format")
	ErrInvalidPathFormat  = errors.New("invalid API path format")
)

// ValidateReadOnlyQuery ensures the query is SELECT-only and safe
func ValidateReadOnlyQuery(query string) error {
	// Trim and normalize
	trimmed := strings.TrimSpace(query)
	upper := strings.ToUpper(trimmed)

	// Must start with SELECT
	if !strings.HasPrefix(upper, "SELECT") {
		return ErrNotSelectQuery
	}

	// Forbidden keywords for read-only enforcement
	forbiddenKeywords := []string{
		"INSERT", "UPDATE", "DELETE", "DROP", "CREATE", "ALTER",
		"TRUNCATE", "EXEC", "EXECUTE", "PRAGMA", "VACUUM",
	}

	for _, keyword := range forbiddenKeywords {
		if strings.Contains(upper, keyword) {
			return ErrForbiddenKeyword
		}
	}

	// Check for semicolons (multiple statements)
	if strings.Contains(trimmed, ";") {
		return ErrForbiddenKeyword
	}

	return nil
}

// ValidateAPIPath validates the API endpoint path format
func ValidateAPIPath(path string) error {
	// Path must start with /api/ and contain only alphanumeric, hyphens, underscores, slashes
	pattern := `^/api/[a-zA-Z0-9_\-/]+$`
	matched, _ := regexp.MatchString(pattern, path)

	if !matched {
		return ErrInvalidPathFormat
	}

	return nil
}
