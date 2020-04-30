import React from 'react';
import { connect } from 'react-redux';
import PageLayout from '../../../../components/PageLayout';
import * as API from '../../../../api';
import { State as GlobalState } from '../../../../store';
import { updateProject } from '../../../../store/Project';

interface Props {
  project: API.Project;
  updateProject: (project: API.Project) => void;
  match: any;
  history: any;
}

interface State {
  project: API.Project;
}

class EditProjectScene extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { project: props.project };
  }

  render(): JSX.Element {
    const { project } = this.state;
    
    return(
      <PageLayout>
        <header>
          <h1>Edit {project.name}</h1>

          <button className='btn danger' onClick={this.onCancel.bind(this)}>
            <p>Cancel</p>
          </button>

          <button className='btn' onClick={this.onSave.bind(this)}>
            <p>Save</p>
          </button>
        </header>

        <form onSubmit={(e: any) => {e.preventDefault(); this.onSave()}}>
          <label>Name</label>
          <input
            required
            type='text' 
            name='name'
            value={project.name}
            onChange={this.onChange.bind(this)} />
            
          <label>Web Domain</label>
          <input
            type='text' 
            name='web_domain'
            value={project.web_domain}
            onChange={this.onChange.bind(this)} />
          
          <label>API Domain</label>
          <input
            type='text'
            name='api_domain' 
            value={project.api_domain}
            onChange={this.onChange.bind(this)} />
        </form>
      </PageLayout>
    );
  }

  onChange(e: any): void {
    this.setState({
      project: {
        ...this.state.project,
        [e.target.name]: e.target.value,
      },
    });
  }

  onSave(): void {
    this.props.updateProject(this.state.project);
    this.props.history.push('/projects');
  }

  onCancel(): void {
    // eslint-disable-next-line no-restricted-globals
    if (!confirm(`Are you sure you want to cancel? All your changes will be lost.`)) return;
    this.props.history.push('/projects');
  }
}

function mapStateToProps(state: GlobalState, ownProps: Props): any {
  return({
    project: state.project.projects.find(u => u.id === ownProps.match.params.id),
  });
}

function mapDispatchToProps(dispatch: Function): any {
  return({
    updateProject: (project: API.Project) => dispatch(updateProject(project)),
  })
}

export default connect(mapStateToProps, mapDispatchToProps)(EditProjectScene);