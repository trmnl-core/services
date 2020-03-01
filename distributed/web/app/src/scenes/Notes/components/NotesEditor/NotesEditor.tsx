import React from 'react';
import { Note } from '../../../../api';
import './NotesEditor.scss';

interface Props {
  note: Note;
}

interface State {
  note: Note;
}

export default class NotesEditor extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { note: props.note };
  }

  onTitleChanged(e: any) {
    this.setState({
      note: {
        ...this.state.note,
        title: e.target.value,
      }
    })
  }

  onTextChanged(e: any) {
    this.setState({
      note: {
        ...this.state.note,
        text: e.target.value,
      }
    })
  }

  render(): JSX.Element {
    const { title, text, id } = this.state.note;

    return(
      <form className='NotesEditor'>
        <input
          type='text'
          value={title}
          placeholder={id ? 'Note title' : 'Create a new note'}
          autoFocus={!id}
          onChange={this.onTitleChanged.bind(this)} />

        <textarea
          value={text}
          placeholder='Note text'
          onChange={this.onTextChanged.bind(this)} />
      </form>
    );
  }
}