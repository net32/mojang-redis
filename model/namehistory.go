package model

type NameHistoryEntry struct {
	Name        string `json:"name"`
	ChangedToAt int64  `json:"changedToAt"`
}
