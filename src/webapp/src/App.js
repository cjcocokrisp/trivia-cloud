import Game from './Game';
import logo from './logo.svg';
import './App.css';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Select from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import Box from '@mui/material/Box';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import { useState, useEffect } from 'react';

function SimpleDialog({
    open,
    onClose,
    numQuestions,
    setNumQuestions,
    category,
    setCategory,
    showNumWarning,
    showCategoryWarning,
    createGame,
}) {
    const [categories, setCategories] = useState([]);

    const handleNumQuestionsChange = (event) => {
        const inputValue = event.target.value;
        if (/^\d*$/.test(inputValue)) {
            setNumQuestions(inputValue);
        }
    };

    useEffect(() => {
        // Fetch trivia categories from the API
        fetch('https://opentdb.com/api_category.php')
            .then((res) => res.json())
            .then((data) => {
                setCategories(data.trivia_categories);
            });
    }, []);

    const handleCategoryChange = (event) => {
        setCategory(categories[event.target.value]);
    };

    return (
        <Dialog
            open={open}
            onClose={onClose}
            PaperProps={{
                style: {
                    padding: '20px',
                    backgroundColor: '#f9f9f9',
                    borderRadius: '10px',
                },
            }}
        >
            <DialogTitle style={{ textAlign: 'center', fontSize: '1.5rem', fontWeight: 'bold' }}>
                Game Settings
            </DialogTitle>
            <Box padding={2} display="flex" flexDirection="column" gap="20px">
                <TextField
                    id="outlined-basic"
                    value={numQuestions}
                    onChange={handleNumQuestionsChange}
                    label="Number of Questions"
                    variant="outlined"
                    fullWidth
                />
                {showNumWarning && (
                    <div style={{ color: 'red', marginTop: '5px', fontSize: '0.9rem' }}>
                        <strong>Warning:</strong> Please enter a value between 1 and 50.
                    </div>
                )}
                <FormControl fullWidth>
                    <InputLabel id="select-label">Category</InputLabel>
                    <Select
                        labelId="select-label"
                        id="simple-select"
                        value={category}
                        onChange={handleCategoryChange}
                    >
                        {categories.map((cat, index) => (
                            <MenuItem key={cat['id']} value={index}>
                                {cat['name']}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>
                {showCategoryWarning && (
                    <div style={{ color: 'red', marginTop: '5px', fontSize: '0.9rem' }}>
                        <strong>Warning:</strong> Please choose a valid category.
                    </div>
                )}
                <Button onClick={createGame} variant="contained" fullWidth style={{ marginTop: '10px' }}>
                    Create
                </Button>
            </Box>
        </Dialog>
    );
}

function App() {
    const [username, setUsername] = useState('Player');
    const [id, setId] = useState('');
    const [numQuestions, setNumQuestions] = useState(10);
    const [category, setCategory] = useState('');
    const [connectiontype, setConnectionType] = useState('');
    const [inGame, setInGame] = useState(false);
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const [showNumWarning, setShowNumWarning] = useState(false);
    const [showCategoryWarning, setShowCategoryWarning] = useState(false);

    const validateCode = () => {
        setInGame(true);
        setConnectionType('join');
    };

    const openDialog = () => {
        setIsDialogOpen(true);
    };

    const createGame = () => {
        if (numQuestions === '' || numQuestions < 1 || numQuestions > 50) {
            setShowNumWarning(true); // Show warning if the number is invalid
        } else if (!category) {
            setShowCategoryWarning(true); // Show warning if category isn't chosen
        } else {
            setInGame(true); // Create the game
            setConnectionType('create');
            setShowNumWarning(false);
            setShowCategoryWarning(false);
        }
    };

    const closeDialog = () => {
        setIsDialogOpen(false);
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
                        style={{ marginBottom: '20px' }}
                    />
                    <TextField
                        id="outlined-basic"
                        onChange={(event) => setId(event.target.value)}
                        label="Game ID"
                        variant="outlined"
                        style={{ marginBottom: '20px' }}
                    />
                    <div className="button-container" style={{ display: 'flex', gap: '10px' }}>
                        <Button onClick={openDialog} variant="contained">
                            Create
                        </Button>
                        <Button onClick={validateCode} variant="contained">
                            Join
                        </Button>
                    </div>
                </div>
                <SimpleDialog
                    open={isDialogOpen}
                    onClose={closeDialog}
                    numQuestions={numQuestions}
                    setNumQuestions={setNumQuestions}
                    category={category}
                    setCategory={setCategory}
                    showNumWarning={showNumWarning}
                    showCategoryWarning={showCategoryWarning}
                    createGame={createGame}
                />
            </div>
        );
    else
        return (
            <Game
                connectiontype={connectiontype}
                username={username}
                numQuestions={numQuestions}
                category={category}
                id={id}
            />
        );
}

export default App;
