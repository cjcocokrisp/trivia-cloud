package models

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
