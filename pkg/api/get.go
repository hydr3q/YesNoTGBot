package api

import (
	"encoding/json"
	"net/http"
)

func GetAnswer() Answer {
	resp, err := http.Get("https://yesno.wtf/api")

	if err != nil {
		return Answer{}
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return Answer{}
	}

	var answer Answer

	switch result["answer"].(string) {
	case "yes":
		answer.Value = "Да"
	case "no":
		answer.Value = "Нет"
	case "maybe":
		answer.Value = "Возможно"
	}

	answer.Image = result["image"].(string)

	return answer
}
