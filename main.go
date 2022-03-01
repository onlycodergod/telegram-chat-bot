package main

import (
	"go/format"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type wallet map[string]

var db = map[int64]wallet{}

func main() {
	bot, err := tgbotapi.NewBotAPI("qwertqwewqryuwwtryuqtryequtyyequyruiqyruieyteiquty")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			continue
		}

		command := strings.Split(update.Message.Text, sep:" ")

		switch command[0] {
		case "ADD":
			if len(command) != 3 {
				bot.NewMessage(tgbotapi.NewMessage(update.Message.Chat.ID, text:"unknown command"))
			}
			amount, err := strconv.ParseFLoat(command[2], bitSize:64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error)))
			}
			db[update.Message.Chat.ID][command[1]] += amount
         
			case "SUB":
				if len(command) != 3 {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text:"unknown command")
				}
			amount, err := strconv.ParseFloat(command[2], bitSize: 64)

				case "DEL":
					if len(command) != 2 {
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text:"unknown command"))
					}
					delete(db[update.Message.Chat.ID], command[1])
					case "SHOW":
						msg := ""
						var sum float64 = 0
						for key, value := range db[update.Message.Chat.ID] {
							price, _ := getPrice(key)
							sum += value*price
							msg += fmt.Sprintf(format:"Total: %.2f\n", sum)
						}
						msg += fmt.Sprintf(format:"Total: %.2f\n", sum)
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
					default:
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))

		}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			bot.Send(msg)
		}
	}

	func getPRice(symbol string) (price float64, err error) {
		http.Get(url:)
	}