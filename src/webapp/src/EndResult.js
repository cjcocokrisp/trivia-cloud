import './Game.css'

function EndResult(props) {
    const rankings = props['rankings']

    return (
        <div className='Game-header'>
            <div className='Question-container'>
                <p className='Question-header-text'>Rankings:</p>
                {
                    rankings.map((player, index) => {
                        return (<p className='Question-header-text'>{index + 1}. {player['username']} ({player['score']})</p>)
                    })
                }
                <p className='Question-header-text'>Refresh the tab to play again!</p>
            </div>
        </div>
    )
}

export default EndResult;