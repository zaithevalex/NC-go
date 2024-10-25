package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

const accessToken = "8098355627:AAFahiUf5Z4C_n6FHYt4xJ9GV1dDRP0m2EU"

func main() {
	ctx := context.Background()

	opts := []bot.Option{
		bot.WithDefaultHandler(mainMenuHandler),
		// bot.WithMessageTextHandler("/payment", bot.MatchTypeExact, handlerPaymentCommand),
	}

	b, err := bot.New(accessToken, opts...)
	if err != nil {
		log.Println(err.Error())
		return
	}

	b.Start(ctx)
}

func mainMenuHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Click by button",
		ReplyMarkup: models.ReplyKeyboardMarkup{
			Keyboard: [][]models.KeyboardButton{
				{
					{Text: "start button"},
				},
				{
					{Text: "finish button"},
				},
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerPaymentCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("CONTENT:", update.Message.Text)
	_, err := b.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID:          update.Message.Chat.ID,
		MessageThreadID: 0,
		Title:           "Invoice title",
		Description:     "PAYROLL",
		Payload:         "test-invoice-payload",
		ProviderToken:   "381764678:TEST:99020",
		Currency:        "rub",
		StartParameter:  "Subscription",
		Prices: []models.LabeledPrice{
			{Label: "Price 1", Amount: 12425},
			{Label: "Price 2", Amount: 32454},
		},
	})

	if err != nil {
		log.Println(err.Error())
	}
}
