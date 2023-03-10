package mods

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Структуры для работы с Telegram API

type TelegramResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat    Chat    `json:"chat"`
	Text    string  `json:"text"`
	Sticker Sticker `json:"sticker"`
}

type Sticker struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type SendMessage struct {
	ChatId    int    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type SendSticker struct {
	ChatId     int    `json:"chat_id"`
	StickerUrl string `json:"sticker"`
}

// Функция отправки сообщения
func SendMsg(botUrl string, update Update, msg string) error {

	// Формирование сообщения
	botMessage := SendMessage{
		ChatId:    update.Message.Chat.ChatId,
		Text:      msg,
		ParseMode: "HTML",
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		log.Printf("json.Marshal error: %s", err)
		return err
	}

	// Отправка сообщения
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("sendMessage error: %s", err)
		return err
	}
	return nil
}

// Функция отправки стикера
func SendStck(botUrl string, update Update, url string) error {

	// Формирование стикера
	botStickerMessage := SendSticker{
		ChatId:     update.Message.Chat.ChatId,
		StickerUrl: url,
	}
	buf, err := json.Marshal(botStickerMessage)
	if err != nil {
		log.Printf("json.Marshal error: %s", err)
		return err
	}
	// Отправка стикера
	_, err = http.Post(botUrl+"/sendSticker", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("sendSticker error: %s", err)
		return err
	}
	return nil
}
