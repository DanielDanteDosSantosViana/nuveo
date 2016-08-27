package http

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const URL string = `^((ftp|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]([a-zA-Z0-9-]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`

var rxURL = regexp.MustCompile(URL)
var (
	ErrorBadRequest = errors.New("Error ao tentar acessar a url informada")
	ErrorInvalidUrl = errors.New("Formato inválido para a URL")
)

func Send(path string) (*http.Response, error) {
	valid := isUrl(path)
	if !valid {
		return nil, ErrorInvalidUrl
	}
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp, ErrorBadRequest
	}
	return resp, nil
}

func isUrl(path string) bool {
	if path == "" || len(path) >= 2083 || len(path) <= 3 || strings.HasPrefix(path, ".") {
		return false
	}
	u, err := url.Parse(path)
	if err != nil {
		return false
	}

	if strings.HasPrefix(u.Host, ".") {
		return false
	}

	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return rxURL.MatchString(path)
}
