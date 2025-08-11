package utils

import (
	"regexp"
	"time"
)

type FolderInfo struct {
	Name    string
	ModTime time.Time
}

type FolderInfos []FolderInfo

var safeFileName = regexp.MustCompile("^[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_]+)*$")

func IsSafeFileName(fileName string) bool {
	return safeFileName.MatchString(fileName)
}
