import React from 'react';
import logo from './logo.svg';
import Navigation from './components/Navigation';
import './App.css';

class App extends React.Component {
    async componentDidMount() {
        try {
            // TODO: Fetch user info from api using JWT if stored in `localStorage`
            //       store info in redux
          } catch {
            // If token is not valid/not present redirect to landing/login page
          }
    }
  
    render() {
      return (
        <div className="AppContainer">
          <Navigation />
        </div>
      );
    }
}

export default App;
