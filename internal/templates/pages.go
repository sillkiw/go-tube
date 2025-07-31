package templates

import "time"

type folderInfo struct {
	Name    string
	ModTime time.Time
}
type PageList struct {
	Files     []folderInfo
	PrevPage  int
	NextPage  int
	TotalPage int
	CanDelete int
}
type PageQueque struct {
	QuequeSize int
}
type PageUploaded struct {
	FileName      string
	FileNameNoExt string
	QuequeSize    int
}
type PageVPEMB struct {
	VidNm string
}
type PageVP struct {
	VidNm string
	Embed bool
}
type PageVPNoJS struct {
	VidNm string
}
type PageErr struct {
	ErrMsg string
}
type PageSndFile struct {
	UseAuth bool
}
