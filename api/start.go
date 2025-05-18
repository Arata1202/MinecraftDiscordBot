package api

import (
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func StartCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "start",
		Description: "サーバを起動します",
	}
}

func StartHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	url := os.Getenv("LAMBDA_START_URL")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "サーバの起動を開始します",
		},
	})

	resp, err := http.Get(url)
	if err != nil {
		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "起動処理呼び出し失敗",
		})
		return
	}
	defer resp.Body.Close()
}
