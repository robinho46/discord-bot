package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
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
	// Försök ladda .env för lokal utveckling
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using environment variables instead.")
	}

	// Läser miljövariabler
	token = os.Getenv("TOKEN")
	channelID = os.Getenv("CHANNEL_ID")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	// Logga hela meddelandet och innehållet
	log.Printf("Meddelande mottaget: '%s' i kanal: '%s'", m.Content, m.ChannelID)

	quote, err := getRandomQuote("quotes.txt")
	if err != nil {
		println(err)
		return
	}
	if m.Content == "!quote" || m.Content == "!quote random" || m.Content == "!quote Random" {
		s.ChannelMessageSend(channelID, "## Your quote is:\n"+quote)
	}

	if m.Content == "!quote motivation" || m.Content == "!quote Motivation" {
		motivationQuote, err := getMotivationQuote("quotes.txt")
		if err != nil {
			println(err)
			return
		}
		s.ChannelMessageSend(channelID, "## Your funny quote is:\n"+motivationQuote)
	}

}

func getRandomQuote(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("kunde inte läsa filen: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	var quotes []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			quotes = append(quotes, trimmed)
		}
	}

	if len(quotes) == 0 {
		return "", fmt.Errorf("inga quotes finns")
	}

	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex], nil
}
func getMotivationQuote(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("kunde inte läsa filen: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	var quotes []string
	readingMotivation := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if we reach the '# motivation' section
		if trimmed == "# Motivation" {
			readingMotivation = true
			continue // Skip the '# motivation' line itself
		}

		// Stop reading if we reach the '# funny' section
		if readingMotivation && strings.HasPrefix(trimmed, "# Funny") {
			break
		}

		// Stop reading if another section starts after '# motivation'
		if readingMotivation && (strings.HasPrefix(trimmed, "#") && trimmed != "# motivation") {
			break
		}

		// Add the quote if we're reading the motivation section
		if readingMotivation && trimmed != "" {
			quotes = append(quotes, trimmed)
		}
	}

	if len(quotes) == 0 {
		return "", fmt.Errorf("inga motivation quotes finns")
	}

	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex], nil
}

func sendDailyQuote(dg *discordgo.Session) {
	for {
		now := time.Now()

		nextRun := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())

		// Om klockan redan har passerat hoppa över
		if now.After(nextRun) {
			nextRun = nextRun.Add(24 * time.Hour)
		}

		waitDuration := time.Until(nextRun)
		log.Printf("Nästa dagliga quote skickas om: %v\n", waitDuration)
		time.Sleep(waitDuration)

		quote, err := getRandomQuote("quotes.txt")
		if err != nil {
			log.Println("Daily quote error:", err)
			continue
		}
		_, err = dg.ChannelMessageSend(channelID, "@everyone\n## Daily Quote:\n"+quote)
		if err != nil {
			log.Println("Failed to send daily quote:", err)
		}
	}
}
