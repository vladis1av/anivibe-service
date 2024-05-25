package telegram

import (
	"fmt"
	"strings"
	"text/template"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FeedbackForm struct {
	Name    string
	Email   string
	Subject string
	Message string
}

const feedbackTemplate = `Новая форма обратной связи:
{{ if .Name }}Имя: {{ .Name }}{{ end }}
Email: {{ .Email }}
Тема: {{ .Subject }}
Сообщение: {{ .Message }}
`

func formatFeedbackMessage(form *FeedbackForm) (string, error) {
	tmpl, err := template.New("feedback").Parse(feedbackTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse feedback template: %w", err)
	}

	var message strings.Builder
	if err := tmpl.Execute(&message, form); err != nil {
		return "", fmt.Errorf("failed to execute feedback template: %w", err)
	}

	return message.String(), nil
}

func SendFeedback(adminUserID int64, form *FeedbackForm, imageBytes [][]byte) error {
	message, err := formatFeedbackMessage(form)
	if err != nil {
		return err
	}

	if len(imageBytes) > 5 {
		return fmt.Errorf("too many images, maximum allowed is 5")
	}

	if len(imageBytes) > 0 {
		var mediaGroup []interface{}

		for i, bytes := range imageBytes {
			if i >= 5 {
				break
			}

			if len(bytes) > 0 {
				photoFileBytes := tgbotapi.FileBytes{
					Name:  "image.jpg",
					Bytes: bytes,
				}

				photo := tgbotapi.NewInputMediaPhoto(photoFileBytes)
				mediaGroup = append(mediaGroup, photo)
			}
		}

		textMessage := tgbotapi.NewMessage(adminUserID, message)
		mediaGroup = append(mediaGroup, textMessage)

		mediaGroupConfig := tgbotapi.NewMediaGroup(adminUserID, mediaGroup)
		_, err = bot.Send(mediaGroupConfig)
		if err != nil {
			return fmt.Errorf("failed to send media group: %w", err)
		}
	} else {
		msg := tgbotapi.NewMessage(adminUserID, message)
		_, err = bot.Send(msg)
		if err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}
	}

	return nil
}
