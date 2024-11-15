package auradomain

import "time"

type Post struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uint      `json:"user_id"`
}
