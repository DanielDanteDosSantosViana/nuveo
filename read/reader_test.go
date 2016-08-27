package read

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Deve_Ler_Saida_Do_Response(t *testing.T) {
	var msg string = "Hello Teste\n"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, msg, http.StatusOK)
	}))
	defer server.Close()

	resp, _ := http.Get(server.URL)
	reader := New()
	body, err := reader.Read(resp)
	if body.String() != msg {
		t.Errorf(" Error : Deveria retornar %s , porém retornou %s)", msg, body.String())
	}
}
