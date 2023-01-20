package repository

// creating `Note` struct with relevant fields
type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r Repo) SendNote(note *Note) error {
	return nil
}
