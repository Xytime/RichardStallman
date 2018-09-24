package main

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"os/signal"
)
var discord *discordgo.Session

const richard = "I'd just like to interject for a moment. What you're referring to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself, but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, shell utilities and vital system components comprising a full OS as defined by POSIX.\n\nMany computer users run a modified version of the GNU system every day, without realizing it. Through a peculiar turn of events, the version of GNU which is widely used today is often called \"Linux\", and many of its users are not aware that it is basically the GNU system, developed by the GNU Project.\n\nThere really is a Linux, and these people are using it, but it is just a part of the system they use. Linux is the kernel: the program in the system that allocates the machine's resources to the other programs that you run. The kernel is an essential part of an operating system, but useless by itself; it can only function in the context of a complete operating system. Linux is normally used in combination with the GNU operating system: the whole system is basically GNU with Linux added, or GNU/Linux. All the so-called \"Linux\" distributions are really distributions of GNU/Linux."

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(strings.ToLower(m.Content), "linux") && !strings.Contains(strings.ToLower(m.Content), "gnu") {
		s.ChannelMessageSend(m.ChannelID, richard)
	}
}


func onReady(s *discordgo.Session, event *discordgo.Ready) {
	print("Bot is ready!")
	s.UpdateStatus(0, "GNU/Linux")
}

func main() {
	var err error
	Token := os.Getenv("DISCORD_BOT_TOKEN")
	discord, err= discordgo.New("Bot " + Token)
	if err != nil {
		print(err)
		os.Exit(1)
	}
	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onReady)
	err = discord.Open()
	if err != nil {
		print("Failed to connect to discord.")
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	discord.Close()
}