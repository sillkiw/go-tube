package templates

import (
	"gotube/internal/utils"
)

type PageList struct {
	Files     []utils.FolderInfo
	Page      int
	PrevPage  int
	NextPage  int
	TotalPage int
	CanDelete bool
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
