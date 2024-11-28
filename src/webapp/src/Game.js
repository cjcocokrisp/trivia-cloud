import './Game.css';
import config from "./config";
import Lobby from './Lobby.js';
import Question from './Question.js'
import QuestionResult from './QuestionResult.js';

import { useEffect, useState } from "react";

function Game(props) {
    const { connectiontype, username, numQuestions, category, id } = props;
    const [ players, setPlayers ] = useState([username]);
    const [ gameId, setGameId ] = useState("");
    const [ gameState, setGameState ] = useState("Lobby");
    const [ currentQuestion, setCurrentQuestion ] = useState(null);
    const [ questionNum, setQuestionNum ] = useState(0);
    const [ websocket, setWebsocket ] = useState(null);
    const [ categoryName, setCategoryName ] = useState(category['name']);
    const [ submitted, setSubmitted ] = useState(false);
    const [ correct, setCorrect ] = useState(false);

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
                setQuestionNum(data["content"]["numQuestions"]);
                setCategoryName(data["content"]["category"]);
                break;
            case "new_connection":
                setPlayers(players => [...players, data["content"]]);
                break;
            case "disconnection":
                setPlayers(data["content"]);
                break;
            case "question_start":
                setQuestionNum(1);
                setCurrentQuestion(data["content"]);
                setGameState('Playing');
                break;
            case "submission":
                setSubmitted(true);
                setCorrect(data['content']['correct']);
                if (data['content']['allsubmitted']) {
                    const payload = {
                        action: "sendResult"
                    }
                    socket.send(JSON.stringify(payload));
                }
                break;
            case "view_result":
                setSubmitted(false);
                setGameState('QuestionResult');
                setCurrentQuestion(data["content"]);
                break;
            case "next_submission":
                setSubmitted(true);
                if (data['content']) {
                    const payload = {
                        action: "sendQuestion"
                    }
                    socket.send(JSON.stringify(payload));
                }
                break;
            case "next_question":
                setQuestionNum(data['content']['num'] + 1);
                setSubmitted(false);
                setCurrentQuestion(data['content']['information']);
                setGameState('Playing');
                break;
            case "gameover":
                console.log('GAME OVER SHOW RESULTS!');
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
                    category={categoryName} 
                    questionnum={questionNum} 
                    connection={websocket}
                />
        )
    case 'Playing':
        return (<Question
                    question={currentQuestion}
                    questionNum={questionNum}
                    connection={websocket}
                    submitted={submitted}
                />
        )
    case 'QuestionResult':
        return (<QuestionResult
                    correct={correct}
                    correctAnswer={currentQuestion}
                    connection={websocket}
                    submitted={submitted}
                />
            
        )
    }
}

export default Game;