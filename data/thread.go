package data

import "time"

// Thread struct
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
}

// Threads - get all thread
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.ID, &th.UUID, &th.Topic, &th.UserID, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

// NumReplies - Count total of post by thread_id
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(); err != nil {
			return
		}
	}
	rows.Close()
	return
}
