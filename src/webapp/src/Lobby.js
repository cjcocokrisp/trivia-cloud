import './Game.css';
import logo from './logo.svg';
import Button from '@mui/material/Button'

function Lobby(props) {
    const gameId = props['GameId'];
    const players = props['players'];
    console.log(props);

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

export default Lobby;