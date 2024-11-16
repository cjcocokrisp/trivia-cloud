import Game from './Game'
import logo from './logo.svg';
import './App.css';
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import { useState } from 'react';

function App() {
    const [ username, setUsername ] = useState("Player");
    const [ id, setId] = useState("");
    const [ connectiontype, setConnectionType ] = useState("");
    const [ inGame, setInGame ] = useState(false);

    const validateCode = () => {
        setInGame(true);
        setConnectionType("join");
    }

    if (!inGame)
    return (
        <div className="App">
        <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <TextField id="outlined-basic" value={username} onChange={(event) => setUsername(event.target.value)} label="Username" variant='outlined'/>
            <TextField id="outlined-basic" onChange={(event) => setId(event.target.value)} label="Game ID" variant='outlined'/>
            <div className="button-container">
                <Button onClick={() => {setInGame(true); setConnectionType("create")}} variant="contained">Create</Button>
                <Button onClick={validateCode} variant="contained">Join</Button>
            </div>
        </header>
        </div>
    );
    else 
    return (
        <Game connectiontype={connectiontype} username={username} id={id}/>
    );
}

export default App;
