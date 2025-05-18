package api

import (
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func StopCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "stop",
		Description: "サーバを停止します",
	}
}

func StopHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	url := os.Getenv("LAMBDA_STOP_URL")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "サーバの停止を開始します",
		},
	})

	resp, err := http.Get(url)
	if err != nil {
		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "停止処理呼び出し失敗",
		})
		return
	}
	defer resp.Body.Close()
}
