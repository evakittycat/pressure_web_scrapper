package service

import (
	sql "github.com/jmoiron/sqlx"
)

type SiteTypeService struct {
	db     *sql.DB
	parser Parser
}

//func New(db *sql.DB) *SiteTypeService {
//	return &SiteTypeService{db: db}
//}

func New(parser Parser) *SiteTypeService {
	return &SiteTypeService{parser: parser}
}
