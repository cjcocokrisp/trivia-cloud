package models

// contains structs that represent database items
type Player struct {
	Username     string `json:"username" dynamodbav:"username"`
	ConnectionId string `json:"connectionId" dynamodbav:"connectionId"`
	Connected    bool   `json:"connected" dynamodbav:"connected"`
	Score        int    `json:"score" dynamodbav:"score"`
	Submitted    bool   `json:"submitted" dynamodbav:"submitted"`
	Correct      bool   `json:"correct" dynamodbav:"correct"`
}

type Game struct {
	GameId          string     `json:"gameId" dynamodbav:"gameId"`
	Started         bool       `json:"started" dynamodbav:"started"`
	CurrentQuestion int        `json:"currentQuestion" dynamodbav:"currentQuestion"`
	NumQuestions    string     `json:"numQuestions" dynamodbav:"numQuestions"`
	Questions       []Question `json:"questions" dynamodbav:"questions"`
	Category        string     `json:"category" dynamodbav:"category"`
	Players         []Player   `json:"players" dynamodbav:"players"`
}

type Connection struct {
	ConnectionId string `json:"connectionId" dynamodbav:"connectionId"`
	GameId       string `json:"gameId" dynamodbav:"gameId"`
}

// this struct is a message struct that uses a generic type for the content so it can be anything
type Message[T any] struct {
	Type    string `json:"type"`
	Content T      `json:"content"`
}

type GameInformation struct {
	GameId       string   `json:"gameId"`
	Players      []string `json:"players"`
	NumQuestions string   `json:"numQuestions"`
	Category     string   `json:"category"`
}

type OpenTriviaDBResponse struct {
	ResponseCode string     `json:"response_code"`
	Results      []Question `json:"results"`
}

type Question struct {
	Type       string   `json:"type" dynamodbav:"type"`
	Difficulty string   `json:"difficulty" dynamodbav:"difficulty"`
	Category   string   `json:"category" dynamodbav:"category"`
	Question   string   `json:"question" dynamodbav:"question"`
	Correct    string   `json:"correct_answer" dynamodbav:"correct_answer"`
	Incorrect  []string `json:"incorrect_answers" dynamodbav:"incorrect_answers"`
}

type QuestionInformation struct {
	Difficulty string   `json:"difficulty"`
	Category   string   `json:"category"`
	Question   string   `json:"question"`
	Choices    []string `json:"choices"`
}

type PlayerAnswer struct {
	Action string `json:"action"`
	Answer string `json:"answer"`
}

type AnswerResponse struct {
	Correct      bool `json:"correct"`
	AllSubmitted bool `json:"allsubmitted"`
}

type NextQuestion struct {
	QuestionNum         int                 `json:"num"`
	QuestionInformation QuestionInformation `json:"information"`
}
