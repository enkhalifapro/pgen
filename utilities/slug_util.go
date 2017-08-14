package utilities

import (
	"strings"
)

type SlugUtil struct {
}

// TODO: rewrite this...
func (s *SlugUtil) GetSlug(name string) string {
	slug := strings.TrimSpace(name)
	slug = strings.Replace(slug, "!", "", -1)
	slug = strings.Replace(slug, "#", "", -1)
	slug = strings.Replace(slug, "#", "", -1)
	slug = strings.Replace(slug, "$", "", -1)
	slug = strings.Replace(slug, "%", "", -1)
	slug = strings.Replace(slug, "&", "", -1)
	slug = strings.Replace(slug, "’", "", -1)
	slug = strings.Replace(slug, "(", "", -1)
	slug = strings.Replace(slug, ")", "", -1)
	slug = strings.Replace(slug, "*", "", -1)
	slug = strings.Replace(slug, "+", "", -1)
	slug = strings.Replace(slug, ",", "", -1)
	slug = strings.Replace(slug, ".", "", -1)
	slug = strings.Replace(slug, "/", "", -1)
	slug = strings.Replace(slug, ":", "", -1)
	slug = strings.Replace(slug, `\`, "", -1)
	slug = strings.Replace(slug, ";", "", -1)
	slug = strings.Replace(slug, "<", "", -1)
	slug = strings.Replace(slug, "=", "", -1)
	slug = strings.Replace(slug, ">", "", -1)
	slug = strings.Replace(slug, "?", "", -1)
	slug = strings.Replace(slug, "@", "", -1)
	slug = strings.Replace(slug, "[", "", -1)
	slug = strings.Replace(slug, "]", "", -1)
	slug = strings.Replace(slug, "^", "", -1)
	slug = strings.Replace(slug, "`", "", -1)
	slug = strings.Replace(slug, "~", "", -1)
	slug = strings.Replace(slug, "|", "", -1)
	slug = strings.Replace(slug, "{", "", -1)
	slug = strings.Replace(slug, "}", "", -1)
	slug = strings.Replace(slug, " ", "-", -1)
	return slug
}

func (s *SlugUtil) GetName(slug string) string {
	return strings.Replace(slug, "-", " ", -1)
}
