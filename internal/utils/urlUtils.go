package utils

// GetFullShortURL returns the full URL by appending the given short URL to the base URL.
//
// Parameters:
// - shortURL: the short URL to be appended to the base URL.
//
// Returns:
// - string: the full URL formed by appending the short URL to the base URL.
func GetFullShortURL(shortURL string) string {
	return "http://localhost:8080/" + shortURL
}
