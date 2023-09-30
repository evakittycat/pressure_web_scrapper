package domain

type SiteType struct {
	Category string `json:"category" db:"categories"`
	Theme    string `json:"theme" db:"themes"`
}

type Responses interface {
	// вернут
}
