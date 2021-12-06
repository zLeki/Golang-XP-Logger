package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"main/functions"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + "ez")
	if err != nil {
		log.Println("error created while making a bot")
		return
	}
	dg.AddHandler(help)
	dg.AddHandler(on_ready)
	dg.AddHandler(ping)
	dg.AddHandler(xpCheck)
	dg.AddHandler(source)
	dg.AddHandler(giveXp)
	err = dg.Open()
	if err != nil {
		log.Println("Error created while opening the bot", err)
		return
	}
	log.Println("Bot is up and running :sunglasses:")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
func on_ready(s *discordgo.Session, e *discordgo.Ready) {
	log.Println("Bot online real")
	s.UpdateGameStatus(1, "c!help | real clan labs")
}
func GenerateEmbed(embed chan *discordgo.MessageEmbed, typeE string, title string, description string, thumbnail string) {
	embedMake := functions.EmbedCreate(typeE, title, description, thumbnail)
	embed <- embedMake
}
func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		if m.Content == "c!help" {
			embed := make(chan *discordgo.MessageEmbed)
			defer close(embed)
			go GenerateEmbed(embed, "thumbnail", "Help", "c!xp\nc!givexp\nc!format\nc!removexp\nc!source\nc!ping", "https://i.imgur.com/NldSwaZ.png")
			realEmbed := <-embed
			s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
		}}
}
func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "c!ping" {
		if m.Author.ID != s.State.User.ID {
			embed := make(chan *discordgo.MessageEmbed)
			defer close(embed)
			go GenerateEmbed(embed, "thumbnail", "Ping", "Ping -> "+s.HeartbeatLatency().String(), "https://i.imgur.com/v2n7qPs.png")
			realEmbed := <-embed
			s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
		}
	}
}
func source(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "c!source" {
		embed := make(chan *discordgo.MessageEmbed)
		defer close(embed)
		go GenerateEmbed(embed,"image", "Source code!", "Visit the repository here, https://github.com/zLeki/Golang-XP-Logger", "https://opengraph.githubassets.com/c2b6b57b2921bba36ddb0026699de77f5c86beae11e7055dc94dff42f3988a14/zLeki/Golang-XP-Logger")
		realEmbed := <-embed
		s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
	}
}
func giveXp(s *discordgo.Session, m *discordgo.MessageCreate) {


	if strings.HasPrefix(m.Content, "c!givexp ") {
		var requiredPermission = map[int]string{
			discordgo.PermissionAdministrator: "Administrator",
		}
		perm, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			log.Fatalf("error", err)
		}
		for k, v := range requiredPermission {
			if int(perm) & k != k {
				embed := make(chan *discordgo.MessageEmbed)
				defer close(embed)
				go GenerateEmbed(embed, "thumbnail", "Error", "You are missing the permission; "+v, "https://i.imgur.com/qs4QOjF.png")
				realEmbed := <-embed
				s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
				return
			}
		}
		User := strings.Split(m.Content, " ")[1]
		Amount := strings.Split(m.Content, " ")[2]
		f, err := os.OpenFile(m.GuildID+".txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(User+":"+Amount+"\n"); err != nil {
			log.Println(err)
		}
	}else if strings.HasPrefix(m.Content, "c!givexp")  {
		embed := make(chan *discordgo.MessageEmbed)
		defer close(embed)
		go GenerateEmbed(embed, "thumbnail", "Error", "Your missing a parameter", "https://i.imgur.com/qs4QOjF.png")
		realEmbed := <-embed
		s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
	}else if strings.HasPrefix(m.Content, "c!removexp")  {
		var requiredPermission = map[int]string{
			discordgo.PermissionAdministrator: "Administrator",
		}
		perm, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			log.Fatalf("error", err)
		}
		for k, v := range requiredPermission {
			if int(perm) & k != k {
				embed := make(chan *discordgo.MessageEmbed)
				defer close(embed)
				go GenerateEmbed(embed, "thumbnail", "Error", "You are missing the permission; "+v, "https://i.imgur.com/qs4QOjF.png")
				realEmbed := <-embed
				s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
				return
			}
		}
		User := strings.Split(m.Content, " ")[1]
		Amount := strings.Split(m.Content, " ")[2]
		f, err := os.OpenFile(m.GuildID+".txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(User+":-"+Amount+"\n"); err != nil {
			log.Println(err)
		}
		embed := make(chan *discordgo.MessageEmbed)
		defer close(embed)
		go GenerateEmbed(embed, "thumbnail", "Success", "Successfully removed "+Amount+" from the user "+User, "https://i.pinimg.com/originals/70/a5/52/70a552e8e955049c8587b2d7606cd6a6.gif")
		realEmbed := <-embed
		s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
	}}
	func xpCheck(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, "c!xp ") {
			embed := make(chan *discordgo.MessageEmbed)
			defer close(embed)
			go GenerateEmbed(embed, "thumbnail", "Error", "Please re-run the command I have auto-fixed the error", "https://i.imgur.com/qs4QOjF.png")
			realEmbed := <-embed
			if _, err := os.Stat(m.GuildID+".txt"); errors.Is(err, os.ErrNotExist) {
				s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
				err := ioutil.WriteFile(m.GuildID+".txt", []byte("//NEW FILE\n"), 0755)
				if err != nil {
					log.Fatal("Error occured", err)
					return
				}
				return
			}
			var User = strings.Split(m.Content, " ")[1]
			dat, err := os.ReadFile(m.GuildID+".txt")
			if err != nil {
				log.Fatalf("Error occured", err)
			}
			slice1 := strings.Split(string(dat), "\n")
			var summary = 0
			for i := range slice1 {
				if strings.HasPrefix(slice1[i], User) {
					log.Println("Counting..")
					idk := strings.Split(slice1[i], ":")[1]
					intver, err := strconv.Atoi(idk)
					if err != nil {
						log.Fatalf("Error", err)
					}
					summary+=intver
				}
			}
			log.Println(summary)
			for i := range slice1 {
				if strings.HasPrefix(slice1[i], User) {
					log.Println("Found user")
					resp, err := http.Get("https://api.roblox.com/users/get-by-username?username="+User)
					if err != nil {
						log.Fatalf("Error", err)
					}
					defer func(Body io.ReadCloser) {
						err := Body.Close()
						if err != nil {
						}
					}(resp.Body)
					b, err := httputil.DumpResponse(resp, true)
					if err != nil {
						log.Fatalln(err)
					}
					id := strings.Split(string(b), `"Id":`)[1]
					idReal := strings.Split(id, `,"Username`)[0]
					log.Println(idReal)
					embed := make(chan *discordgo.MessageEmbed)
					defer close(embed)
					go GenerateEmbed(embed, "thumbnail", "Your xp", "You have -> "+strconv.Itoa(summary), "https://www.roblox.com/headshot-thumbnail/image?userId="+idReal+"&width=420&height=420&format=png")
					realEmbed := <-embed
					s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
					return
				}
			}
			//if !found {
			log.Println("User doesnt exist.. Creating..")
			s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
			f, err := os.OpenFile(m.GuildID+".txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(User+":0\n"); err != nil {
				log.Println(err)
			}
		}else if strings.HasPrefix(m.Content, "c!xp")  {
			embed := make(chan *discordgo.MessageEmbed)
			defer close(embed)
			go GenerateEmbed(embed, "thumbnail", "Error", "Your missing a parameter", "https://i.imgur.com/qs4QOjF.png")
			realEmbed := <-embed
			s.ChannelMessageSendEmbed(m.ChannelID, realEmbed)
		}
	}
