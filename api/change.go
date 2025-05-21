package api

import (
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func ChangeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "change",
		Description: "インスタンスタイプを変更し、起動します",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "変更するインスタンスタイプ",
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

func ChangeHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var url string
	var message string

	if option, ok := optionMap["type"]; ok {
		instanceType := option.StringValue()
		switch instanceType {
		case "t3a":
			url = os.Getenv("LAMBDA_CHANGE_T3A_URL")
			message = "インスタンスタイプを t3a.medium に変更し、サーバーの起動を開始します"
		case "c7i":
			url = os.Getenv("LAMBDA_CHANGE_C7I_URL")
			message = "インスタンスタイプを c7i.xlarge に変更し、サーバーの起動を開始します"
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
			Content: "変更処理呼び出し失敗",
		})
		return
	}
	defer resp.Body.Close()
}
