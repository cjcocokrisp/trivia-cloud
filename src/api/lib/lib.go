package lib

import (
	"trivia-cloud/backend/lib/models"
)

func GetConnectedUsers(game models.Game) []string {
	var users []string
	for _, element := range game.Players {
		if element.Connected {
			users = append(users, element.Username)
		}
	}
	return users
}
