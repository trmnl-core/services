// Frameworks
import React from 'react';
import { connect } from 'react-redux';
import { NavLink } from 'react-router-dom';

// Utils
import { State as GlobalState } from '../../store';
import * as API from '../../api';

// Components
import ProjectSwitcher from './components/ProjectSwitcher';

// Styling
import Logo from './assets/logo.png';
import NavDashboard from './assets/nav-dashboard.png';
import NavProjects from './assets/nav-projects.png';
import NavBilling from './assets/nav-billing.png';
import NavSettings from './assets/nav-settings.png';
import './style.scss';

interface Props {
  user: API.User;
  className?: string;
}

class PageLayout extends React.Component<Props> {
  render(): JSX.Element {
    const { profile_picture_url, first_name, last_name } = this.props.user;

    return(
      <div className='PageLayout'>
        <div className='sidebar'>
          <img src={Logo} alt='M3O Logo' className='logo' />

          <nav>
            <a href='https://web.micro.mu' target='blank'>
              <img src={NavDashboard} alt='Dashboard' />
              <p>Dashboard</p>
            </a>

            <NavLink to='/projects'>
              <img src={NavProjects} alt='Projects' />
              <p>Projects</p>
            </NavLink>
            
            <NavLink exact to='/billing'>
              <img src={NavBilling} alt='Billing' />
              <p>Billing</p>
            </NavLink>

            <NavLink exact to='/settings'>
              <img src={NavSettings} alt='Settings' />
              <p>Settings</p>
            </NavLink>
          </nav>

          <div className='lower'>
            <ProjectSwitcher />

            <a href='https://account.micro.mu' target='blank'>
              { profile_picture_url ? <img className='account' src={profile_picture_url} alt='Your account' /> : <div className='initials'>
                <p>{first_name.slice(0,1)}{last_name.slice(0,1)}</p>
              </div> }
            </a>
          </div>
        </div>

        <div className={`main ${this.props.className}`}>
          { this.props.children }
        </div>
      </div>
    );
  }
}

function mapStateToProps(state: GlobalState): any {
  return({
    user: state.account.user,
  });
}

export default connect(mapStateToProps)(PageLayout);
