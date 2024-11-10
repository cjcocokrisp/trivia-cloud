package models

type Player struct {
	Username     string `json:"username" dynamodbav:"username"`
	ConnectionID string `json:"connectionID" dynamodbav:"connectionID"`
	Connected    bool   `json:"connected" dynamodbav:"connected"`
}

type Game struct {
	GameID  int      `json:"gameID" dynamodbav:"gameID"`
	Players []Player `json:"players" dynamodbav:"players"`
}
