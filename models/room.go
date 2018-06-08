package models

type Room struct {
	Name        string    `json:"name"`
	Messages    []Message `json:"messages"` // oldest messages come first
	Authors     []Author  `json:"authors"`
	HistorySize int       `json:"historySize"`
}

func (r *Room) AddMessage(message Message) {
	r.Messages = append(r.Messages, message)
	if len(r.Messages) > r.HistorySize {
		r.Messages = r.Messages[1:]
	}
}

func (r *Room) AddAuthor(author Author) {
	r.RemoveAuthor(author)
	r.Authors = append(r.Authors, author)
}

func (r *Room) RemoveAuthor(author Author) {
	for i, a := range r.Authors {
		if a.Id == author.Id {
			r.Authors = append(r.Authors[:i], r.Authors[i+1:]...)
			return
		}
	}
}
