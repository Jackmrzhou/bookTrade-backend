package utils

import "strings"

func TranslateExt(ext string) string {
	s := strings.TrimSpace(ext)
	if s == "jpg" || s == "jpeg" {
		return "image/jpeg"
	} else if s == "png" {
		return "image/png"
	} else {
		return "application/octet-stream"
	}
}

func GetExt(fileName string) string {
	res := strings.Split(fileName, ".")
	if len(res) == 1 {
		return ""
	} else {
		return res[len(res)-1]
	}
}
