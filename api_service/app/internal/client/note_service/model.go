package note_service

type CreateNoteDTO struct {
	Header       string `json:"header"`
	Body         string `json:"body"`
	ShortBody    string `json:"short_body,omitempty"`
	CategoryUUID string `json:"category_uuid"`
}

type UpdateNoteDTO struct {
	Header       string `json:"header,omitempty"`
	Body         string `json:"body,omitempty"`
	CategoryUUID string `json:"category_uuid,omitempty"`
}
