package dal

import "time"

type Keyword struct {
	ID        uint64
	Word      string
	CreatedAt time.Time
}

type Domain struct {
	ID        uint64
	Link      string
	CreatedAt time.Time
}

type Reference struct {
	ID        uint64
	DomainID  uint64
	Href      string
	CreatedAt time.Time
}
