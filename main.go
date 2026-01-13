package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken  string
	serverURL string
)

func init() {
	// Récupère les variables d'environnement
	botToken = os.Getenv("DISCORD_BOT_TOKEN")
	serverURL = os.Getenv("SERVER_URL")

	if botToken == "" {
		log.Fatal("La variable d'environnement DISCORD_BOT_TOKEN n'est pas définie.")
	}
	if serverURL == "" {
		log.Fatal("La variable d'environnement SERVER_URL n'est pas définie.")
	}
}

func main() {
	// Crée une nouvelle session Discord
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Erreur lors de la création de la session Discord :", err)
		return
	}

	// Ajoute un handler pour les messages
	dg.AddHandler(messageCreate)

	// Ouvre une connexion WebSocket avec Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la connexion :", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot démarré avec succès !")

	// Attend un signal pour fermer le bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// messageCreate est appelé chaque fois qu'un message est créé
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore les messages envoyés par le bot lui-même
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Vérifie si le message commence par "!"
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	// Parse la commande
	command := strings.TrimPrefix(m.Content, "!")
	args := strings.Fields(command)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "update":
		sendCommandToServer("update", s, m)
	case "restart":
		sendCommandToServer("restart", s, m)
	case "whitelist":
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Usage: !whitelist <joueur>")
			return
		}
		player := args[1]
		sendCommandToServer(fmt.Sprintf("whitelist&player=%s", player), s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "Commande inconnue. Utilisez !update, !restart ou !whitelist <joueur>.")
	}
}

// sendCommandToServer envoie une commande au serveur local
func sendCommandToServer(command string, s *discordgo.Session, m *discordgo.MessageCreate) {
	url := fmt.Sprintf("%s?command=%s", serverURL, command)
	resp, err := http.Get(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Erreur lors de l'envoi de la commande : %v", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Commande `%s` envoyée avec succès !", command))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Erreur : le serveur a retourné un code %d", resp.StatusCode))
	}
}
