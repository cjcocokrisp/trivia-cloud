import './Game.css';
import he from 'he';
import Button from '@mui/material/Button'

function Question(props) {
    const question = props['question'];
    const questionNum = props['questionNum']
    const connection = props['connection'];

    return (
        <div className='Game-header'>
            <div className='Question-container'>
                <p className='Question-header-text'>Question {questionNum}</p>
                <p className='Question-header-text'>Category: {he.decode(question['category'])}</p>
                <p className='Question-header-text'>{he.decode(question['question'])}</p>
                <div className='Question-buttons'>
                    {
                        question['choices'].map(choice => {
                            return (<Button
                                variant='contained'
                                className='Question-button'
                            >
                                {choice}
                            </Button>
                            )
                        })
                    }
                </div>
            </div>
        </div>
    )
}

export default Question;