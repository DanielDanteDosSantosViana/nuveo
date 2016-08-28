package robot

import (
	"errors"
	"github.com/DanielDanteDosSantosViana/nuveo/converter"
	sender "github.com/DanielDanteDosSantosViana/nuveo/http"
	"github.com/DanielDanteDosSantosViana/nuveo/models"
	"io/ioutil"
	"strings"
)

var (
	ErrorNotFoundTypeFile = errors.New("Não foi encontrado o tipo do arquivo !")
)

type Robot struct {
}

func New() *Robot {
	return &Robot{}
}

func (r *Robot) GetData(url string) ([]models.Data, error) {
	resp, err := sender.Send(url)
	if err != nil {
		return nil, err
	}
	tipo := resp.Header.Get("Content-Type")
	conv := converter.New()
	if isJson(tipo) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		data := []models.Data{}
		err = conv.ToJson(body, &data)
		if err != nil {
			return nil, err
		}
		return data, nil

	} else if isCsv(tipo) {
		data := []models.Data{}
		err := conv.ToCsv(resp.Body, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	defer resp.Body.Close()
	return []models.Data{}, ErrorNotFoundTypeFile
}

func isJson(str string) bool {
	return strings.Contains(str, "application/json")
}

func isCsv(str string) bool {
	return strings.Contains(str, "text/csv") || strings.Contains(str, "application/octet-stream")

}
