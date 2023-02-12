package mods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура респонса
type InfoResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Drinks  []DrinkInfo `json:"result"`
}

// Структура коктейля
type DrinkInfo struct {
	Name           string `json:"name"`
	Price          int    `json:"price"`
	Flavour        string `json:"flavour"`
	Primary_Type   string `json:"primary_type"`
	Secondary_Type string `json:"secondary_type"`
	Recipe         string `json:"recipe"`
	Shortcut       string `json:"shortcut"`
	Description    string `json:"description"`
}

// Функция отправки рецептов
func SendDrinkInfo(botUrl string, update Update, parameters []string) error {

	// Rest запрос для получения апдейтов
	resp, err := http.Get("https://vall-halla-api.vercel.app/api/info?shortcut=5xT")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Запись и обработка полученных данных
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var response InfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	for _, d := range response.Drinks {
		SendDrink(botUrl, update, d)
	}

	return nil

}

// Функция отправки рецепта
func SendDrink(botUrl string, update Update, drink DrinkInfo) {
	SendMsg(botUrl, update, fmt.Sprintf(
		"%s\nIt's a <strong>%s</strong>, <strong>%s</strong> and <strong>%s</strong> drink coasting <strong>$%d</strong>\n"+
			"<b>Recipe</b> - %s\n<b>Shortcut</b> - %s\n\n<i>\"%s\"</i>",
		drink.Name, drink.Flavour, drink.Primary_Type, drink.Secondary_Type, drink.Price, drink.Recipe, drink.Shortcut, drink.Description))
}