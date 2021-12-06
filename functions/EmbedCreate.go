package functions

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func EmbedCreate(typeE string, title string, description string, thumbnail string) *discordgo.MessageEmbed {
	switch typeE {
	case "thumbnail":
		embed := &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
				Name:   "Clan labs real",
				Value:  description,
				Inline: true,
			},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: thumbnail,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Made by Leki#6796",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     title,
		}
		return embed
	case "image":
		embed := &discordgo.MessageEmbed{
			Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
				Name:   "Clan labs real",
				Value:  description,
				Inline: true,
			},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: thumbnail,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Made by Leki#6796",
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Title:     title,
		}
		return embed
	}
	embed := &discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
			Name:   "Missing information",
			Value:  "The retard cough developer of this bot is missing some information for this function, hopefully his dumbass can fix it soon",
			Inline: true,
		},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.imgur.com/NldSwaZ.png",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Leki development stuff",
	}
	return embed
}