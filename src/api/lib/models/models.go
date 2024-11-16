package models

// contains structs that represent database items
type Player struct {
	Username     string `json:"username" dynamodbav:"username"`
	ConnectionId string `json:"connectionId" dynamodbav:"connectionId"`
	Connected    bool   `json:"connected" dynamodbav:"connected"`
}

type Game struct {
	GameId  string   `json:"gameId" dynamodbav:"gameId"`
	Players []Player `json:"players" dynamodbav:"players"`
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
	GameId  string   `json:"gameId"`
	Players []string `json:"players"`
}
