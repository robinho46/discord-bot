package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"math/rand" // For random number generation
	"os"        // For reading the quotes.txt file
	"strings"
	"time"
)

var (
	token     string
	channelID string
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Add this
	loadEnv()

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		return
	}

	go sendDailyQuote(dg)

	select {}
}

func loadEnv() {
	err := godotenv.Load() // Loads .env by default
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token = os.Getenv("TOKEN") // Assign token from .env
	channelID = os.Getenv("CHANNEL_ID")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "!quote") {
		quote, err := getRandomQuote("quotes.txt")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Något gick fel ^_^")
			return
		}
		s.ChannelMessageSend(m.ChannelID, quote)
	}
}

func getRandomQuote(filename string) (string, error) {
	// 1. Read entire file
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("kunde inte läsa filen: %v", err)
	}

	// 2. Split into lines and clean
	lines := strings.Split(string(data), "\n")
	var quotes []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			quotes = append(quotes, trimmed)
		}
	}

	// 3. Check for empty quotes
	if len(quotes) == 0 {
		return "", fmt.Errorf("inga quotes finns")
	}

	// 4. Pick random quote
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex], nil
}

func sendDailyQuote(dg *discordgo.Session) {
	ticker := time.NewTicker(24 * time.Hour)
	// ticker := time.NewTicker(5 * time.Second) -- used for testing

	defer ticker.Stop()

	for range ticker.C {
		quote, err := getRandomQuote("quotes.txt")
		if err != nil {
			log.Println("Daily quote error:", err)
			continue
		}
		_, err = dg.ChannelMessageSend(channelID, "@everyone\nDaily Quote:\n"+quote)
		if err != nil {
			log.Println("Failed to send daily quote:", err)
		}
	}
}
