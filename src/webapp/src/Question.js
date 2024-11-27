import './Game.css';
import he from 'he';
import Button from '@mui/material/Button'

function Question(props) {
    const question = props['question'];
    const questionNum = props['questionNum']
    const connection = props['connection'];
    const submitted = props['submitted'];

    const submitAnswer = (choice) => {
        return () => {
            const payload = {
                'action': 'submitAnswer',
                'answer': choice
            }
            connection.send(JSON.stringify(payload));
        }
    }

    function determineQuestionContent() {
        if (submitted) {
            return (<p className='Question-header-text'>Waiting for other players to answer...</p>)
        } else {
            return question['choices'].map(choice => {
                return (<Button
                    variant='contained'
                    className='Question-button'
                    onClick={submitAnswer(choice)}
                >
                    {choice}
                </Button>
                )
            })
        }
    }

    return (
        <div className='Game-header'>
            <div className='Question-container'>
                <p className='Question-header-text'>Question {questionNum}</p>
                <p className='Question-header-text'>Category: {he.decode(question['category'])}</p>
                <p className='Question-header-text'>{he.decode(question['question'])}</p>
                <div className='Question-buttons'>
                    { 
                        determineQuestionContent()
                    }
                </div>
            </div>
        </div>
    )
}

export default Question;