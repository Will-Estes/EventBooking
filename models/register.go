package models

type Register struct {
	ID      int64 `json:"id"`
	EventId int64 `json:"event_id"`
	UserId  int64 `json:"user_id"`
}
