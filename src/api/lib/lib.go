package lib

import (
	"sort"
	"trivia-cloud/backend/lib/models"
)

func GetConnectedUsers(game models.Game) []models.Player {
	var users []models.Player
	for _, element := range game.Players {
		if element.Connected {
			users = append(users, element)
		}
	}
	return users
}

func PrepareQuestionInfo(game models.Game) models.QuestionInformation {
	question := game.Questions[game.CurrentQuestion]

	var choices []string
	choices = append(choices, question.Correct)
	choices = append(choices, question.Incorrect...)

	return models.QuestionInformation{
		Difficulty: question.Difficulty,
		Category:   question.Category,
		Question:   question.Question,
		Choices:    choices,
	}
}

func SortPlayerByScore(players []models.Player) []models.Player {
	sort.SliceStable(players, func(x, y int) bool {
		return players[x].Score > players[y].Score
	})
	return players
}
