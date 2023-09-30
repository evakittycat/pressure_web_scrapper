package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"rocketJump5G/internal/domain"
)

const (
	getSiteType = `SELECT`
)

type GetThemeByCatResp struct {
	Themes []string `json:"themes,omitempty"`
}

type GetSiteByCatResp struct {
	Sites []string `json:"sites,omitempty"`
}

func (s *SiteTypeService) Get(ectx echo.Context) error {
	name := ectx.QueryParam("name")

	var siteType []domain.SiteType

	var statement string

	if name != "" {
		statement = fmt.Sprint(getSiteType)

		err := s.db.Select(&siteType, statement, name)
		if err != nil {
			return ectx.String(http.StatusInternalServerError, err.Error()) // think about it
		}
	} else {
		statement = fmt.Sprintf(getSiteType, "")

		err := s.db.Select(&siteType, statement)
		if err != nil {
			return ectx.String(http.StatusInternalServerError, err.Error()) // think about it
		}
	}

	return ectx.JSON(http.StatusOK, siteType)
}

type Parser interface {
	GetThemeByCategory(category string) []string
	GetSiteByCategory(category, theme string) []string
}

func (s *SiteTypeService) GetThemeByCategory(c echo.Context) error {
	category := c.Param("category")
	return c.JSON(200, GetThemeByCatResp{
		Themes: s.parser.GetThemeByCategory(category),
	})
}

func (s *SiteTypeService) GetSiteByCategory(c echo.Context) error {
	category := c.Param("category")
	theme := c.Param("theme")

	return c.JSON(200, GetSiteByCatResp{
		Sites: s.parser.GetSiteByCategory(category, theme),
	})
}
