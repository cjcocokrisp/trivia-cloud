import './Game.css'
import Button from '@mui/material/Button'

function QuestionResult(props) {
    const correct = props['correct'];
    const correctAnswer = props['correctAnswer'];
    const connection = props['connection'];
    const submitted = props['submitted'];

    const submitNext = () => {
        const payload = {
            'action': 'submitNext'
        }
        connection.send(JSON.stringify(payload));
    }

    function determineNextContent() {
        if (submitted) {
            return (<p className='Question-header-text'>Waiting for other players...</p>)
        } else {
            return (<Button 
                    variant='contained'
                    className='Question-button'
                    onClick={submitNext} 
                    >
                        Next
                    </Button>
            )
        }
    }

    return (
        <div className='Game-header'>
            <div className='Question-container'>
                <p className='Question-header-text'>Question Result:</p>
                <p className='Question-header-text'>{correct ? "Correct" : "Incorrect"}</p>
                <p className='Question-header-text'>Correct Answer:</p>
                <p className='Question-header-text'>{correctAnswer}</p>
                {
                    determineNextContent()
                }
            </div>

        </div>
    )

}

export default QuestionResult;