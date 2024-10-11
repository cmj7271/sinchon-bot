package main

import bot "example.com/hello_world_bot/Bot"

func main() {
	bot.BotToken = ""
	bot.Run() // call the run function of bot/bot.go
}
