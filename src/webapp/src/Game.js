import './Game.css';
import config from "./config";
import Lobby from './Lobby.js';

import { useEffect, useState } from "react";

function Game(props) {
    const { connectiontype, username, id } = props;
    const [ players, setPlayers ] = useState([username]);
    const [ gameId, setGameId ] = useState("");

    let url = new URL(config.API_ENDPOINT);
    url.searchParams.append("connectiontype", connectiontype);
    url.searchParams.append("username", username);
    url.searchParams.append("id", id)

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
                console.log(data["content"]);
                setPlayers(data["content"]);
                break;
            }
        })
    }, [])

    console.log(gameId);
    
    return (<Lobby GameId={gameId} players={players} />)
}

export default Game;