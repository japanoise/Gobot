package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

var target string

var waifu map[string]string

const channel string = "##dankville"

var nick *string = flag.String("n", "gobot", "Nickname")

func main() {
	pass := flag.String("P", "", "Connection Password")
	server := flag.String("s", "chat.freenode.net", "Server to connect to")
	port := flag.Int("p", 6667, "Port to use")
	usetls := flag.Bool("z", false, "Use TLS")
	noverify := flag.Bool("v", false, "Skip TLS connection verification")
	flag.Parse()
	client, err := New(TlsCon{*usetls, *noverify},
		fmt.Sprint(*server, ":", *port), *nick, *nick)
	if err != nil {
		log.Fatalln("Could not connect to IRC server; ", err.Error())
	}
	if *pass == "" {
		client.Auth()
	} else {
		client.Authpass(*pass)
	}
	f, err := os.Open("waifus.json")
	if err == nil {
		loadwaifus(f)
		f.Close()
	} else {
		fmt.Println(err.Error(), ", using a blank db for now.")
		waifu = make(map[string]string)
	}
	client.Send(Join(channel))
	go sighandle(client)
	mainloop(client)
}

func loadwaifus(fi *os.File) {
	dec := json.NewDecoder(fi)
	if err := dec.Decode(&waifu); err != nil {
		fmt.Println(err.Error(), ", using a blank db for now.")
		waifu = make(map[string]string)
	}
}

func savewaifus(fi *os.File) {
	enc := json.NewEncoder(fi)
	if err := enc.Encode(&waifu); err != nil {
		fmt.Println(err.Error())
	}
}

func sighandle(client *Client) {
	sigchan := make(chan os.Signal, 10)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan
	fmt.Println("Recieved interrupt; exiting gracefully.")
	cleanup(client)
	os.Exit(0)
}

func cleanup(client *Client) {
	client.Send(Quit("Gobot terminated gracefully, bye!"))
	f, err := os.Create("waifus.json")
	if err == nil {
		savewaifus(f)
		f.Close()
	} else {
		fmt.Println(err)
	}
}

func printfortune(client *Client, page string) {
	c := exec.Command("fortune", page)
	out, err := c.Output()
	if err == nil {
		strout := strings.Replace(string(out[:]), "\n", " ", -1)
		client.Send(PrivMsg(channel, strout))
	} else {
		fmt.Println(err.Error())
	}
}

func handlemsg(client *Client, msg, name string) {
	//One-word commands
	switch strings.ToLower(msg) {
	case "ping", "!ping":
		client.Send(PrivMsg(channel, "pong"))
	case "!opening":
		client.Send(PrivMsg(channel, Opening()))
	case ".bots":
		client.Send(PrivMsg(channel, "Gobot reporting in! [Golang] https://github.com/japanoise/Gobot"))
	case "!comfort", "i need a hug":
		if waifu[name] == "" {
			client.Send(CTCP("ACTION", channel, fmt.Sprintf("hugs %s",
				name)))
		} else {
			client.Send(PrivMsg(channel, fmt.Sprintf("%s hugs %s",
				waifu[name], name)))
		}
	case "!quote", "!fortune":
		printfortune(client, "login")
	case "!waifu":
		client.Send(PrivMsg(channel, announcewaifu(name)))
	}

	words := strings.Split(msg, " ")

	//Multi-word commands
	switch strings.ToLower(words[0]) {
	case strings.ToLower(*nick) + ":":
		var answer string
		if Yes() {
			answer = "Yes"
		} else {
			answer = "No"
		}
		client.Send(PrivMsg(channel, fmt.Sprintf("%s: %s.", name, answer)))
	case "!translate":
		client.Send(PrivMsg(channel, "That phrase translates to \"My Hovercraft Is Full Of Eels\""))
	case "!quality":
		if len(words) > 1 {
			client.Send(PrivMsg(channel, fmt.Sprint("That was \x1F\x1D\x02", strings.ToUpper(strings.Join(words[1:], " ")), "\x0F / \x0304,02Quality!")))
		}
	case "!setwaifu", "!waifuset", "!waifureg", "!regwaifu":
		if len(words) > 1 {
			waifu[name] = strings.Join(words[1:], " ")
			client.Send(PrivMsg(channel, announcewaifu(name)))
		}
	case "!waifu":
		if len(words) > 1 {
			client.Send(PrivMsg(channel, announcewaifu(words[1])))
		}
	case "!comfort", "!hug":
		if len(words) > 1 {
			targ := words[1]
			if waifu[targ] == "" {
				client.Send(CTCP("ACTION", channel, fmt.Sprintf("hugs %s",
					targ)))
			} else {
				client.Send(PrivMsg(channel, fmt.Sprintf("%s hugs %s",
					waifu[targ], targ)))
			}
		}
	}
}

func announcewaifu(name string) string {
	retval := waifu[name]
	if retval == "" {
		return fmt.Sprintf("Awww, %s doesn't have a waifu :(", name)
	} else {
		return fmt.Sprintf("%s's waifu is the lovely %s.", name, retval)
	}
}

func mainloop(client *Client) {
	for {
		msg, err := client.Receive()
		if err != nil {
			fmt.Println("Output loop closing:", err)
			return
		}
		if msg != nil {
			if msg.Command == "PRIVMSG" && msg.Params[0] == channel && len(msg.Params) > 1 {
				go handlemsg(client, msg.Params[1], msg.Prefix.Name)
			}
			if msg.Command != "PING" {
				fmt.Println(msg)
			}
			if msg.Command == "ERROR" {
				return
			}
		}
	}
}

func Yes() bool {
	return rand.Intn(2) == 1
}
