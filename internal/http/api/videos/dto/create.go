package dto

type CreateRequest struct {
	Title       string `json:"title"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}

type CreateResponse struct {
	ID     string        `json:"id"`
	Status string        `json:"status"`
	Upload UploadDetails `json:"upload"`
	Links  Links         `json:"links,omitempty"`
}

type UploadDetails struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	MaxSize int64             `json:"max_size,omitempty"`
}

type Links struct {
	Self string `json:"self,omitempty"`
}

func (c CreateRequest) GetTitle() string {
	return c.Title
}

func (c CreateRequest) GetContentType() string {
	return c.ContentType
}

func (c CreateRequest) GetSize() int64 {
	return c.Size
}
