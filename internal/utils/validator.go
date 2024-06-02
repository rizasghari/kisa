package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateUrl(URL string) error {
	if URL == "" {
		return fmt.Errorf("url cannot be empty")
	}

	parsedUrl, err := url.ParseRequestURI(URL)
	if err != nil {
		return fmt.Errorf("invalid url: %s", err)
	}

	if !strings.HasPrefix(parsedUrl.Scheme, "http") {
		return fmt.Errorf("url should start with http:// or https://")
	}

	return nil
}
