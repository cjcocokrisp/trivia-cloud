import './Game.css';
import logo from './logo.svg';
import Button from '@mui/material/Button'

function Lobby(props) {
    const gameId = props['GameId'];
    const players = props['players'];
    const category = props['category'];
    const questionnum = props['questionnum'];
    const connection = props['connection'];

    const startGame = () => {
        const payload = {
            action: 'startGame',
        }
        connection.send(JSON.stringify(payload));
    }

    return (
        <div className="Game-header">
            <img src={logo} className="Game-logo" alt="logo" />
            <div className="Game-lobby-screen">
                <div className='Game-lobby-information'>
                    <div>
                        <p className="Game-id">Game ID: {gameId}</p>
                        <p className="Game-id">Category: {category}</p>
                        <p className="Game-id">Number of Questions: {questionnum}</p>
                    </div>
                    <div>
                        <p className='Game-players-header'>Players</p>
                        {
                            players.map((value) => {
                                return (<p className='Game-player'>{value}</p>)
                            })
                        }
                    </div>
                </div>
                <Button 
                    className="Game-start-button" 
                    variant="contained" 
                    onClick={startGame}
                >
                        Start
                </Button>
            </div>
        </div>
    )
}

export default Lobby;