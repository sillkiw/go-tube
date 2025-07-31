package utils

import "regexp"

var safeFileName = regexp.MustCompile("^[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_]+)*$")

func IsSafeFileName(fileName string) bool {
	return safeFileName.MatchString(fileName)
}
