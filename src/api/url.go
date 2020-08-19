package api

import (
	netUrl "net/url"
	"strings"
)

type Url struct {
	*netUrl.URL
}

func NewUrl(baseUrl string) (*Url, error) {
	var err error
	newUrl := &Url{}
	newUrl.URL, err = netUrl.ParseRequestURI(strings.TrimRight(baseUrl, "/"))

	return newUrl, err
}

func (u *Url) Join(path string) (err error) {
	u.URL, err = netUrl.ParseRequestURI(u.String() + "/" + strings.TrimLeft(path, "/"))
	return err
}
