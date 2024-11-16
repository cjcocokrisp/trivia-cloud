import config from "./config";
import { useEffect, useState } from "react";

function Game(props) {
    const { connectiontype, username, id } = props;
    const [ players, setPlayers ] = useState([username]);

    let url = new URL(config.API_ENDPOINT);
    url.searchParams.append("connectiontype", connectiontype);
    url.searchParams.append("username", username);
    url.searchParams.append("id", id)

    useEffect(() => {
        const socket = new WebSocket(url.toString());
        socket.addEventListener("open", (event) => {
            const payload = {
                action: "broadcastConnect"
            }
            socket.send(JSON.stringify(payload));
        })

        socket.addEventListener("message", (event) => {
            setPlayers(players => [...players, event.data]);
        })
    }, [])

    return (
        <div>
            <p>Players</p>
            {
                players.map((value) => {
                    return (<p>{value}</p>)
                })
            }
        </div>
    )
}

export default Game;