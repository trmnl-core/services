import React from 'react';
import { createStore } from 'redux';
import { Provider } from 'react-redux';
import { BrowserRouter , Route } from 'react-router-dom';

// Scenes 
import NotesScene from './scenes/Notes';
import HomeScene from './scenes/Home';
import { rootReducer } from './store';

// Redux
window.store = createStore(
  rootReducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);

export default class App extends React.Component {
  render():JSX.Element {
    return(
      <Provider store={window.store}>
        <BrowserRouter>
          <div className='App'>
            <Route exact path='/' component={HomeScene}/>
            <Route exact path='/notes' component={NotesScene}/>
            <Route exact path='/notes/:id' component={NotesScene}/>
            <Route exact path='/notes/:id/:options' component={NotesScene}/>
          </div>
        </BrowserRouter>
      </Provider>
    );
  }
}