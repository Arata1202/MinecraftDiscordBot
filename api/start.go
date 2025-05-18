package api

import (
	"io"
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

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "サーバの起動を開始します",
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
			Content: "起動処理呼び出し失敗",
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
		Content: "サーバの起動処理が完了しました。応答: " + string(body),
	})
}
