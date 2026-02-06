package videos

type videoDataSaver interface {
	Create(v Video) (string, error)
}
