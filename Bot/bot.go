package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"sort"
	"strings"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {

	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)

	// open session
	discord.Open()
	defer discord.Close() // close session, after function termination

	// keep bot running until there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

//func getBeforeUser(discord *discordgo.Session) discordgo.User {
//	msg := discord.ChannelMessages()
//}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	/* prevent bot responding to its own message
	this is achieved by looking into the message author id
	if message.author.id is same as bot.author.id then just return
	*/
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if strings.Contains(message.Content, "||") {
		return
	}

	if message.Content == "#help" {
		str := ""

		str += "> **<명령어만 쳐야 합니다>**\n"
		str += helpString(cmd_list_only) + "\n"

		str += "> **<명령어 그대로여야 하지만 그외 문자가 있어도 됩니다>**\n"
		str += helpString(cmd_list_white_space) + "\n"

		str += "> **<명령어를 띄어 써도 됩니다>**\n"
		str += helpString(cmd_list_no_white_space)

		discord.ChannelMessageSend(message.ChannelID, str)
		return
	}

	if val, ok := cmd_list_only[message.Content]; ok {
		discord.ChannelMessageSend(message.ChannelID, val)
		return
	}

	for key, val := range cmd_list_no_white_space {
		if strings.Contains(message.Content, key) {
			discord.ChannelMessageSend(message.ChannelID, val)
			return
		}
	}

	message.Content = strings.ReplaceAll(message.Content, " ", "")
	message.Content = strings.ReplaceAll(message.Content, "\n", "")

	for key, val := range cmd_list_white_space {
		if strings.Contains(message.Content, key) {
			discord.ChannelMessageSend(message.ChannelID, val)
			return
		}
	}
}

func helpString(cmd_list map[string]string) string {
	str := ""

	var key []string
	for k := range cmd_list {
		key = append(key, k)
	}
	sort.Strings(key)

	for k := range key {
		val := cmd_list[key[k]]
		str += "> " + "`" + key[k] + "`" + ": "

		if strings.Contains(val, "\n") {
			val = strings.ReplaceAll(val, "\n", "\n> ")
		}

		if strings.Contains(val, "https") {
			val = strings.ReplaceAll(val, "https", "`https")

			if strings.Contains(val, "\n") {
				val = strings.ReplaceAll(val, "\n", "`\n")
			} else {
				val = val + "`"
			}
		}

		str += val
		str += "\n"
	}

	return str
}
