import React from 'react';
import { connect } from 'react-redux';
import PageLayout from '../../../../components/PageLayout';
import * as API from '../../../../api';
import { State as GlobalState } from '../../../../store';
import { updateEnvVar } from '../../../../store/Configuration';
import Form from '../../components/Form';

interface Props {
  envVar: API.EnvVar;
  updateEnvVar: (envVar: API.EnvVar) => void;
  match: any;
  history: any;
}

interface State {
  envVar: API.EnvVar;
}

class EditConfigurationService extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { envVar: props.envVar };
  }

  render(): JSX.Element {
    const { envVar } = this.state;
    
    return(
      <PageLayout>
        <header>
          <h1>Edit {envVar.key}</h1>

          <button className='btn danger' onClick={this.onCancel.bind(this)}>
            <p>Cancel</p>
          </button>

          <button className='btn' onClick={this.onSave.bind(this)}>
            <p>Save</p>
          </button>
        </header>

        <Form envVar={envVar} onSubmit={this.onSave.bind(this)} onUpdate={this.onUpdate.bind(this)} />
      </PageLayout>
    );
  }

  onUpdate(envVar: API.EnvVar): void {
    this.setState({ envVar });
  }

  onSave(): void {
    this.props.updateEnvVar(this.state.envVar);
    this.props.history.push('/configuration');
  }

  onCancel(): void {
    // eslint-disable-next-line no-restricted-globals
    if (!confirm(`Are you sure you want to cancel? All your changes will be lost.`)) return;
    this.props.history.push('/configuration');
  }
}

function mapStateToProps(state: GlobalState, ownProps: Props): any {
  const { params } = ownProps.match;
  return({
    envVar: state.configuration.envVars.find(e => e.service === params.service && e.key === params.key),
  });
}

function mapDispatchToProps(dispatch: Function): any {
  return({
    updateEnvVar: (envVar: API.EnvVar) => dispatch(updateEnvVar(envVar)),
  })
}

export default connect(mapStateToProps, mapDispatchToProps)(EditConfigurationService);