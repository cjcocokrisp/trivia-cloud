import Game from './Game';
import logo from './logo.svg';
import './App.css';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { useState } from 'react';

function App() {
    const [username, setUsername] = useState("Player");
    const [id, setId] = useState("");
    const [connectiontype, setConnectionType] = useState("");
    const [inGame, setInGame] = useState(false);

    // New state for dialog box
    const [openDialog, setOpenDialog] = useState(false);
    const [numQuestions, setNumQuestions] = useState(10);
    const [category, setCategory] = useState("History");

    const validateCode = () => {
        setInGame(true);
        setConnectionType("join");
    };

    if (!inGame)
        return (
            <div className="App">
                <div className="App-header">
                    <img src={logo} className="App-logo" alt="logo" />
                    <TextField
                        id="outlined-basic"
                        value={username}
                        onChange={(event) => setUsername(event.target.value)}
                        label="Username"
                        variant="outlined"
                    />
                    <TextField
                        id="outlined-basic"
                        onChange={(event) => setId(event.target.value)}
                        label="Game ID"
                        variant="outlined"
                    />
                    <div className="button-container">
                        {/* Create Button Opens Dialog */}
                        <Button onClick={() => setOpenDialog(true)} variant="contained">
                            Create
                        </Button>
                        <Button onClick={validateCode} variant="contained">
                            Join
                        </Button>
                    </div>
                </div>

                {/* Dialog Box */}
                <Dialog open={openDialog} onClose={() => setOpenDialog(false)}>
                    <DialogTitle>Game Settings</DialogTitle>
                    <DialogContent>
                        <TextField
                            id="number-of-questions"
                            label="Number of Questions"
                            type="number"
                            value={numQuestions}
                            onChange={(e) => setNumQuestions(e.target.value)}
                            fullWidth
                            variant="outlined"
                            margin="dense"
                        />
                        <TextField
                            id="category"
                            label="Category"
                            value={category}
                            onChange={(e) => setCategory(e.target.value)}
                            fullWidth
                            variant="outlined"
                            margin="dense"
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
                        <Button
                            onClick={() => {
                                setOpenDialog(false);
                                setInGame(true); 
                                setConnectionType("create"); 
                            }}
                            variant="contained"
                        >
                            Save
                        </Button>
                    </DialogActions>
                </Dialog>
            </div>
        );
    else
        return (
            <Game
                connectiontype={connectiontype}
                username={username}
                id={id}
                numQuestions={numQuestions}
                category={category} 
            />
        );
}

export default App;

