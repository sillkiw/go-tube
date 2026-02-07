package videos

type Repo interface {
	Create(v Video) (string, error)
	Get(id string) (Video, error)
}
