import React from 'react'
import PlayerCards from './PlayerCards'

import '../styles/PlayerContainer.sass'

class PlayerContainer extends React.Component {
    render() {
        const {data} = this.props
        return(
            <div className="mr-auto ml-auto">
                <PlayerCards data={data.hand} />
            </div>
        )
    }
}

export default (PlayerContainer)
