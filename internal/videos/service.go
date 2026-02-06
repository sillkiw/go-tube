package videos

type Service struct {
	storage   Storage
	fileStore FileStore
}

func New(storage Storage, fileStore FileStore) *Service {
	return &Service{storage: storage, fileStore: fileStore}
}

func (s *Service) Create(v Video) (string, error) {
	id, err := s.storage.Create(v)
	if err != nil {
		return "", err
	}
	name := "VD_" + id
	s.fileStore.CreateFolder(name)
	return id, err
}
