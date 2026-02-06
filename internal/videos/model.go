package videos

type Status string

const (
	StatusCreated    Status = "created"
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

type Video struct {
	ID     string
	Title  string
	Size   int64
	Status Status
}
