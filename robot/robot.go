package robot

import (
	"fmt"
	"github.com/DanielDanteDosSantosViana/nuveo/http"
	"github.com/DanielDanteDosSantosViana/nuveo/read"
	"net/http"
)

type Robot struct {
}

func New() *Robot {
	return &Robot{}
}

func (r *Robot) getData(url string) (Data, error) {
	resp, err := http.Send(url)
	if err != nil {
		return nil, err
	}
	reader := read.New()
	body, err := reader.Read(resp)

	if err != nil {
		return nil, err
	}

	return nil, err

}

func isJson() {

}
