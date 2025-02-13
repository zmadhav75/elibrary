package models

import "time"

type BorrowRecord struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	DueDate    time.Time `json:"due_date"`
	ReturnedAt time.Time `json:"returned_at"`
}
