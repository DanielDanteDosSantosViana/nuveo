package models

type (
	Data struct {
		Nome   string `json:"nome"`
		Email  string `json:"email"`
		Sexo   string `json:"sexo"`
		Idade  string `json:"idade"`
		Outros string `json:"outros"`
	}
)
