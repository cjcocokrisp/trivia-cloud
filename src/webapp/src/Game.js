import './Game.css';
import config from "./config";
import Lobby from './Lobby.js';
import Question from './Question.js'

import { useEffect, useState } from "react";

function Game(props) {
    const { connectiontype, username, numQuestions, category, id } = props;
    const [ players, setPlayers ] = useState([username]);
    const [ gameId, setGameId ] = useState("");
    const [ gameState, setGameState ] = useState("Lobby");
    const [ currentQuestion, setCurrentQuestion ] = useState(null);
    const [ questionNum, setQuestionNum ] = useState(0);
    const [ websocket, setWebsocket ] = useState(null);

    let url = new URL(config.API_ENDPOINT);
    url.searchParams.append("connectiontype", connectiontype);
    url.searchParams.append("username", username);
    url.searchParams.append("id", id);
    url.searchParams.append("questionnum", numQuestions);
    url.searchParams.append("categoryname", category['name']);
    url.searchParams.append("categorynum", category['id'])

    useEffect(() => {
        const socket = new WebSocket(url.toString());
        socket.addEventListener("open", (event) => {
            const payload = {
                action: "broadcastConnect"
            };
            socket.send(JSON.stringify(payload));
        })

        socket.addEventListener("message", (event) => {
            let data = JSON.parse(event.data)
            switch (data["type"]) {
            case "connected":
                setGameId(data["content"]["gameId"]);
                setPlayers(data["content"]["players"]);
                break;
            case "new_connection":
                setPlayers(players => [...players, data["content"]]);
                break;
            case "disconnection":
                setPlayers(data["content"]);
                break;
            case "question_start":
                setQuestionNum(questionNum + 1);
                setCurrentQuestion(data["content"]);
                setGameState('Playing');
                break;
            }
        });

        setWebsocket(socket);
    }, [])

    switch (gameState) {
    case 'Lobby':
        return (<Lobby 
                    GameId={gameId} 
                    players={players} 
                    category={category['name']} 
                    questionnum={numQuestions} 
                    connection={websocket}
                />
        )
    case 'Playing':
        return (<Question
                    question={currentQuestion}
                    questionNum={questionNum}
                    connection={websocket}
                />
        )
    }
}

export default Game;