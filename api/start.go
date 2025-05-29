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
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "インスタンスタイプ",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "t3a.medium （低スペック）",
						Value: "t3a",
					},
					{
						Name:  "c7i.xlarge （高スペック）",
						Value: "c7i",
					},
				},
			},
		},
	}
}

func StartHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var url string
	var message string

	if option, ok := optionMap["type"]; ok {
		instanceType := option.StringValue()
		baseURL := os.Getenv("LAMBDA_START_URL")
		url = baseURL + "/?type=" + instanceType

		switch instanceType {
		case "t3a":
			message = "サーバーの起動を開始します\nインスタンスタイプ：t3a.medium （低スペック）"
		case "c7i":
			message = "サーバーの起動を開始します\nインスタンスタイプ：c7i.xlarge （高スペック）"
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
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
