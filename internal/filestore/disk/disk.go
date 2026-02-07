package disk

type DiskStore struct {
	DashPath string
	RawPath  string
}

func New(dashPath string, rawPath string) DiskStore {
	return DiskStore{
		DashPath: dashPath,
		RawPath:  rawPath,
	}
}
