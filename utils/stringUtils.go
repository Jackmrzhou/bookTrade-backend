package utils

import "strings"

// trim white space and remove '-'
func FormatISBN(isbn string) string {
	strings.Trim(isbn, " ")
	strings.ReplaceAll(isbn, "-", "")
	return isbn
}
