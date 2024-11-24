import Game from './Game'
import logo from './logo.svg';
import './App.css';
import TextField from '@mui/material/TextField'
import Button from '@mui/material/Button'
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import FormControl from '@mui/material/FormControl';
import Box from '@mui/material/Box';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import InputLabel from '@mui/material/InputLabel';
import { useState } from 'react';

function SimpleDialog({open, onClose, numQuestions, setNumQuestions, category, setCategory, showNumWarning, showCategoryWarning, createGame}) {
    const handleClose = () => {
        console.log();
    }
    const handleNumQuestionsChange = (event) => {
        const inputValue = event.target.value
        if (/^\d*$/.test(inputValue) ) {
            console.log("valid input");
            setNumQuestions(inputValue);
        } else {
            console.log('invalid input "${inputValue}"')
        }
    }
    const handleCategoryChange = (event) => {
        const selectedCategory = event.target.value
        setCategory(selectedCategory)
    }
    return (
        <div className='simple-dialog'>
            <Dialog open={open} onClose={onClose}>
                <DialogTitle className="dialog-title">Game Settings</DialogTitle>
                <TextField 
                id="outlined-basic" 
                value={numQuestions} 
                onChange={handleNumQuestionsChange}
                slotProps={{input: {
                    inputMode: 'numeric',
                    pattern: '[0-9]*', 
                },
            }}
                label="Number of Questions" variant='outlined'/>
            {showNumWarning && (
                <div style={{ color: 'red', marginTop: '10px' }}>
                    <strong>Warning:</strong> Please enter a value between 1-50
                </div>
            )}
                <Box sx={{ minWidth: 120 }}>
                    <FormControl fullWidth>
                        <InputLabel id="select-label">Category</InputLabel>
                            <Select
                                labelId="select-label"
                                id="simple-select"
                                value={category}
                                label="Category"
                                onChange={handleCategoryChange}
                            >
                                <MenuItem value={9}>General Knowledge</MenuItem>
                                <MenuItem value={10}>Entertainment: Books</MenuItem>
                                <MenuItem value={11}>Entertainment: Film</MenuItem>
                                <MenuItem value={12}>Entertainment: Music</MenuItem>
                                <MenuItem value={13}>Entertainment: Musicals & Theatres</MenuItem>
                                <MenuItem value={14}>Entertainment: Television</MenuItem>
                                <MenuItem value={15}>Entertainment: Video Games</MenuItem>
                                <MenuItem value={16}>Entertainment: Board Games</MenuItem>
                                <MenuItem value={17}>Science & Nature</MenuItem>
                                <MenuItem value={18}>Science: Computers</MenuItem>
                                <MenuItem value={19}>Science: Mathematics</MenuItem>
                                <MenuItem value={20}>Mythology</MenuItem>
                                <MenuItem value={21}>Sports</MenuItem>
                                <MenuItem value={22}>Geography</MenuItem>
                                <MenuItem value={23}>History</MenuItem>
                                <MenuItem value={24}>Politics</MenuItem>
                                <MenuItem value={25}>Art</MenuItem>
                                <MenuItem value={26}>Celebrities</MenuItem>
                                <MenuItem value={27}>Animals</MenuItem>
                                <MenuItem value={28}>Vehicles</MenuItem>
                                <MenuItem value={29}>Entertainment: Comics</MenuItem>
                                <MenuItem value={30}>Science: Gadgets</MenuItem>
                                <MenuItem value={31}>Entertainment: Japanese Anime & Manga</MenuItem>
                                <MenuItem value={32}>Entertainment: Cartoon & Animations</MenuItem>
          
                            </Select>
                    </FormControl>
                </Box>
                {showCategoryWarning && (
                    <div style={{ color: 'red', marginTop: '10px' }}>
                        <strong>Warning:</strong> Please choose a category
                    </div>
                )}
                <Button onClick={createGame} variant="contained">Create</Button>
            </Dialog>
        </div>
        
    )
}
function App() {
    const [ username, setUsername ] = useState("Player");
    const [ id, setId] = useState("");
    const [ numQuestions, setNumQuestions] = useState(10);
    const [ category, setCategory] = useState(0);
    const [ connectiontype, setConnectionType ] = useState("");
    const [ inGame, setInGame ] = useState(false);
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const [showNumWarning, setShowNumWarning] = useState(false);
    const [showCategoryWarning, setShowCategoryWarning] = useState(false);

    const validateCode = () => {
        setInGame(true);
        setConnectionType("join");
    }
    const openDialog = () => {
        setIsDialogOpen(true);
    }
    const createGame = () => {
        if (numQuestions === "" || numQuestions < 1 || numQuestions > 50) {
            setShowNumWarning(true); // Show warning if the number is invalid
        } else if (category === undefined || category < 9  || category > 32){
            setShowCategoryWarning(true);  // Show warning if category isnt chosen
        } else {
            //console.log(`num questions: "${numQuestions}" category id: "${category}"`);
            setInGame(true); // Create the game
            setConnectionType("create");
            setShowNumWarning(false);
            setShowCategoryWarning(false);
        }
    }
    const closeDialog = () => {
        setIsDialogOpen(false);
    }
    
    if (!inGame)
    return (
        <div className="App">
        <div className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <TextField id="outlined-basic" value={username} onChange={(event) => setUsername(event.target.value)} label="Username" variant='outlined'/>
            <TextField id="outlined-basic" onChange={(event) => setId(event.target.value)} label="Game ID" variant='outlined'/>
            <div className="button-container">
                <Button onClick={openDialog} variant="contained">Create</Button>
                <Button onClick={validateCode} variant="contained">Join</Button>
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
            createGame={createGame} />
        </div>
    );
    else
    return (
        <Game connectiontype={connectiontype} username={username} numQuestions={numQuestions} category={category} id={id}/>
    );

}

export default App;
