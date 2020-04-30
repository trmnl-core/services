// Libraries
import React from 'react';
import { connect } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';

// Utils
import { State as GlobalState } from './store';
import { setUser } from './store/Account';
import * as API from './api';

// Scenes
import GettingStartedScene from './scenes/GettingStarted';
import ProjectsScene from './scenes/Projects';
import EditProjectScene from './scenes/Projects/scenes/EditProject';
// import InviteTeamMembersScene from './scenes/Team/scenes/InviteTeamMembers';
import ConfigurationScene from './scenes/Configuration';
import EditConfigurationScene from './scenes/Configuration/scenes/EditConfiguration';
import AddConfigurationScene from './scenes/Configuration/scenes/AddConfiguration';

// Styling
import Logo from './components/PageLayout/assets/logo.png';
import './App.scss';

interface Props {
  user?: API.User;
  setUser: (user: API.User) => void;
}

class App extends React.Component<Props> {
  render(): JSX.Element {
    if(this.props.user) return this.renderLoggedIn();
    return this.renderLoading();
  }

  componentDidMount() {
    API.Call("AccountService/Read").then((res) => {
      this.props.setUser(res.data.user);
    });
  }

  renderLoading(): JSX.Element {
    return <div className='loading'>
      <img src={Logo} alt='M3O' />
    </div>
  }

  renderLoggedIn(): JSX.Element {
    return (
      <BrowserRouter>
        <Route key='getting-started' exact path='/' component={GettingStartedScene} />
        <Route key='projects' exact path='/projects' component={ProjectsScene} />
        <Route key='edit-project' exact path='/projects/:id/edit' component={EditProjectScene} />
        <Route key='configuration' exact path='/configuration' component={ConfigurationScene} />
        <Route key='edit-configuration' path='/configuration/:service/:key/edit' component={EditConfigurationScene} />
        <Route key='add-configuration' path='/configuration/add' component={AddConfigurationScene} />
      </BrowserRouter>
    );  
  }
}

function mapStateToProps(state: GlobalState): any {
  return({
    user: state.account.user,
  });
}

function mapDispatchToProps(dispatch: Function): any {
  return({
    setUser: (user: API.User) => dispatch(setUser(user)),
  });
}

export default connect(mapStateToProps, mapDispatchToProps)(App);
