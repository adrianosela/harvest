import React from 'react'
import Main from './pages/main/Main';

import {
    BrowserRouter as Router,
    Switch,
    Route,
} from 'react-router-dom';

class Navigation extends React.Component{
    render() {
        return(
            <Router>
                <Switch>
                    <Route exact path="/" component={Main}/>
                    {/* TODO: Game page has to wrapped with a custom ProtectedRoute Component
                              Protected routes can only be accessed after establishing user session
                    */}
                    <Route path="/game"/>
                    <Route path="/login">
                        {/* Login Component (We might have this inside of Main) */}
                    </Route>
                    <Route path="/recover/password">
                    </Route>
                </Switch>   
            </Router>
        );
    }
}

export default (Navigation);
