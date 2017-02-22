package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strings"
)

var target string

const channel string = "##dankville"

func main() {
	current, _ := user.Current()
	nick := flag.String("n", current.Username, "Nickname")
	pass := flag.String("P", "", "Connection Password")
	user := flag.String("u", current.Username, "Username")
	server := flag.String("s", "chat.freenode.net", "Server to connect to")
	port := flag.Int("p", 6667, "Port to use")
	usetls := flag.Bool("z", false, "Use TLS")
	noverify := flag.Bool("v", false, "Skip TLS connection verification")
	flag.Parse()
	client, err := New(TlsCon{*usetls, *noverify},
		fmt.Sprint(*server, ":", *port), *nick, *user)
	if err != nil {
		log.Fatalln("Could not connect to IRC server; ", err.Error())
	}
	if *pass == "" {
		client.Auth()
	} else {
		client.Authpass(*pass)
	}
	client.Send(Join(channel))
	mainloop(client)
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
	case ".bots":
		client.Send(PrivMsg(channel, "Gobot reporting in! [Golang] https://github.com/japanoise/Gobot"))
	case "!comfort":
		client.Send(PrivMsg(channel, fmt.Sprintf(
			"%s %s", "Gobot loves you,", name)))
	case "!quote", "!fortune":
		printfortune(client, "login")
	}

	words := strings.Split(msg, " ")

	//Multi-word commands
	switch strings.ToLower(words[0]) {
	case "!translate":
		client.Send(PrivMsg(channel, "That phrase translates to \"My Hovercraft Is Full Of Eels\""))
	case "!quality":
		if len(words) > 1 {
			client.Send(PrivMsg(channel, fmt.Sprint("That was \x1F\x1D\x02", strings.ToUpper(strings.Join(words[1:], " ")), "\x0F / \x0304,02Quality!")))
		}
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
