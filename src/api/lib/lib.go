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
