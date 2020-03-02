import React from 'react';
import Call, { Note } from '../../../../api';
import './NotesEditor.scss';

interface Props {
  note: Note;
}

interface State {
  note: Note;
  typingTimer?: NodeJS.Timeout;
}

export default class NotesEditor extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { note: props.note };
  }

  componentDidUpdate(prevProps: Props, prevState: State) {
    if(this.props.note !== prevProps.note) return;
    if(this.state.note === prevState.note) return;

    // If there was already a save scheduled, cancel it
    if (this.state.typingTimer) clearTimeout(this.state.typingTimer);

    // Schedule a save for 500ms, enough time for a user to continue
    // typing, extending by another 500ms.
    this.setState({
      typingTimer: setTimeout(this.saveChanges.bind(this), 500),
    });
  }

  saveChanges() {
    const note = {
      id: this.state.note.id,
      title: this.state.note.title,
      text: this.state.note.text,
    }

    Call('updateNote', { note })
      .then(() => console.log("note saved"))
      .catch(console.warn)
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
          autoFocus={!id}
          placeholder={id ? 'Note title' : 'Create a new note'}
          onChange={this.onTitleChanged.bind(this)} />

        <textarea
          value={text}
          placeholder='Note text'
          onChange={this.onTextChanged.bind(this)} />
      </form>
    );
  }
}