package main

/*#pp

#endpp*/

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"twitchbot/database"

	"github.com/gempir/go-twitch-irc/v2"
)

var channel string
var username string
var token string
var isDebug bool

func init() {
	flag.StringVar(&channel, "channel", "", "channel to join")
	flag.StringVar(&username, "username", "", "username for bot")
	flag.StringVar(&token, "token", "", "token for login")
	flag.BoolVar(&isDebug, "debug", false, "routes output to console instead of chat")
	flag.Parse()

	if isDebug {
		fmt.Println("No Output to Chat")
	}
}

func main() {
	database.Start()

	client := twitch.NewClient(username, token)
	client.SetRateLimiter(twitch.CreateDefaultRateLimiter())

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message)

		database.WriteUser(message.User.ID, message.User.Name, message.User.DisplayName)

		if strings.HasPrefix(message.Message, "!coin") {
			if rand.Intn(100-0)+0 > 50 {
				if isDebug {
					fmt.Println(message.Channel, "=>", fmt.Sprintf("@%s you flipped head", message.User.Name))
					return
				}
				client.Say(message.Channel, fmt.Sprintf("@%s you flipped head", message.User.Name))
			} else {
				if isDebug {
					fmt.Println(message.Channel, "=>", fmt.Sprintf("@%s you flipped number", message.User.Name))
					return
				}
				client.Say(message.Channel, fmt.Sprintf("@%s you flipped number", message.User.Name))
			}
		}

		if strings.HasPrefix(message.Message, "!love") {
			var param []string = strings.Split(message.Message, " ")
			if len(param) == 1 {
				return
			}
			var userA float64 = 0
			var userB float64 = 0
			for _, v := range []byte(message.User.Name) {
				userA += float64(v)
			}
			for _, v := range []byte(param[1]) {
				userB += float64(v)
			}
			var match float64 = ((math.Min(userA, userB) / math.Max(userA, userB)) * 100)

			if isDebug {
				fmt.Println(message.Channel, "=>", fmt.Sprintf("You and %s have a love value of %.2f%s ", param[1], match, "%"))
				return
			}
			client.Say(message.Channel, fmt.Sprintf("You and %s have a love value of %.2f%s ", param[1], match, "%"))
		}

		if strings.HasPrefix(message.Message, "!miesmuschel") {
			var quotesYes []string = []string{
				"Ja",
				"Ja definitiv",
				"Positiv",
				"Hättest du etwas anderes erwartet?",
				"Ja ich will 💍",
				"Nur mit guter bezahlung",
			}
			var quotesNo []string = []string{
				"Nein",
				"Weils du bist NEIN",
				"Nö",
				"Kener hat die Absicht hier Ja zu sagen.",
				"Gegenfrage würdest du nackt und mit Fleisch behängt vor einem hungrigen Tiger tanzen?",
				"Deswegen wird er auch nicht größer also nein!",
				"Ich musste dich jetzt einfach darauf Hinweisen. Du bist so hüpsch wie ein Badewannenstöpsel deswegen muss ich deine Anfrage leider ablehnen.",
				"Nein du stinkst geh dich erstmal waschen!",
				"Sprich mit meiner Hand.",
				"Ihre Bestellung wurde erfolgreich aufgenommen es werden 2502,35€ von ihrem Konto abgebucht. Danke",
				"Nein ich bin tot. Leg den Kranz hin und lass mich in Frieden ruhen",
				"Nein, das ist flüssiger Sonnenschein.",
				"Nein, ich lüge",
				"Nein. Ich bin gerade damit beschäftigt Menschen zu beobachten wie sie sich zum Affen machen.",
				"Diese Sache finde ich genauso positiv wie Durchfall!",
				"NEIN und wenn du nochmal so dämliches zeug frägst werfe ich dich ins Feuer und opfere dich der Göttin Brutzla",
			}
			var quotes []string
			if rand.Intn(100-0)+0 > 50 {
				quotes = quotesYes
			} else {
				quotes = quotesNo
			}

			if isDebug {
				fmt.Println(message.Channel, "=>", fmt.Sprintf("@%s %s", message.User.Name, quotes[rand.Intn(len(quotes)-0)+0]));
				return;
			}
			client.Say(message.Channel, fmt.Sprintf("@%s %s", message.User.Name, quotes[rand.Intn(len(quotes)-0)+0]));
			return;
		}

		// in future own function message parser
/* 		func(message string) string {
			if isDebug {
				fmt.Println(message.Channel, "=>", fmt.Sprintf("@%s %s", message.User.Name, quotes[rand.Intn(len(quotes)-0)+0]))
				return ""
			}
			client.Say(message.Channel, fmt.Sprintf("@%s %s", message.User.Name, quotes[rand.Intn(len(quotes)-0)+0]))
			return ""
		}(message.Message) */
	})

	client.Join(channel)

	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
