import './Game.css';
import logo from './logo.svg'
import config from "./config";
import Button from '@mui/material/Button'
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
                setPlayers(players => [...players, data["content"]])
                break;   
            }
        })
    }, [])

    return (
        <div className="Game-header">
            <img src={logo} className="Game-logo" alt="logo" />
            <div className="Game-lobby-screen">
                <div className='Game-lobby-information'>
                    <p className="Game-id">Game ID: {gameId}</p>
                    <div>
                        <p className='Game-players-header'>Players</p>
                        {
                            players.map((value) => {
                                return (<p className='Game-player'>{value}</p>)
                            })
                        }
                    </div>
                </div>
                <Button className="Game-start-button" variant="contained">Start</Button>
            </div>
        </div>
    )
}

export default Game;