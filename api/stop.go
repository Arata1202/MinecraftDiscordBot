package api

import (
	"io"
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

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "サーバの停止を開始します",
		},
	})

	if err != nil {

		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "コマンド応答中にエラーが発生しました",
		})
		return
	}

	resp, err := http.Get(url)
	if err != nil {

		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "停止処理呼び出し失敗",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
			Content: "サーバからの応答の読み取りに失敗しました",
		})
		return
	}

	s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
		Content: "サーバの停止処理が完了しました。応答: " + string(body),
	})
}
