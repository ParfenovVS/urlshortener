package models

import "time"

type Link struct {
	ID          int
	ShortUrl    string
	OriginalUrl string
	Created     time.Time
}
