import React from 'react';
import { BrowserRouter , Route } from 'react-router-dom';

// Scenes 
import NotesScene from './scenes/Notes';
import HomeScene from './scenes/Home';

export default class App extends React.Component {
  render():JSX.Element {
    return(
      <BrowserRouter>
        <div className='App'>
          <Route exact path='/' component={HomeScene}/>
          <Route exact path='/notes' component={NotesScene}/>
          <Route path='/notes/:id' component={NotesScene}/>
        </div>
      </BrowserRouter>
    );
  }
}