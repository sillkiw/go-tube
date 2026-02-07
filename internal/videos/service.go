package videos

type Service struct {
	repo      Repo
	fileStore FileStore
}

func New(repo Repo, fileStore FileStore) *Service {
	return &Service{repo: repo, fileStore: fileStore}
}

func (s *Service) Create(v Video) (string, error) {
	id, err := s.repo.Create(v)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) Get(videoID string) (Video, error) {
	var videoRec Video
	videoRec, err = s.repo.Get(videoID)
	if err != nil {
		return Video{}, err
	}
	return videoRec, nil
}
