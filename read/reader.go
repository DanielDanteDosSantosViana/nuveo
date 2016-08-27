package read

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Reader struct {
}

func New() *Reader {
	return &Reader{}
}

func (r *Reader) Read(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

func (r *Reader) isJson(body []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(body, &js) == nil
}
