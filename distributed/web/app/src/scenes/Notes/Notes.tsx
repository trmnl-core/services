import React from 'react';
import Call, { Note } from '../../api';
import NotesList from './components/NotesList';
import NotesEditor from './components/NotesEditor';
import PageLayout from '../../components/PageLayout';
import './Notes.scss';

interface Props {
  history: any;
  match: any;
}

interface State {
  notes: Note[];
}

export default class NotesScene extends React.Component<Props, State> {
  _mounted = false;

  constructor(props: Props) {
    super(props);
    this.state = { notes: [] };
  }

  componentDidMount() {
    this._mounted = true;

    // Set the default note when navigating to /notes
    if(!this.props.match.params.id) {
      this.props.history.push('/notes/new');
      return
    }

    // Fetch the notes from the API
    Call('listNotes').catch(console.warn).then(res => {
      if(!this._mounted || !res) return;
      
      const notes = (res.data.notes || []).map((n: any) => {
        return new Note(n);
      });

      this.setState({ notes });
    })
  }

  componentWillUnmount() {
    this._mounted = false;
  }

  render():JSX.Element {
    const { notes } = this.state;
    
    const activeNoteID = this.props.match.params.id;
    const activeNote = notes.find(n => n.id === activeNoteID) || new Note({});

    return(
      <PageLayout className='NotesScene'>
        <h1>Notes</h1>
        <p>There are {notes.length} notes</p>

        <div className='inner'>
          <NotesList
            notes={notes}
            activeNoteID={activeNoteID}
            onNoteClicked={this.onNoteClicked.bind(this)} />
            
          <NotesEditor key={activeNoteID} note={activeNote} />
        </div>
      </PageLayout>
    );
  }

  onNoteClicked(id: string) {
    this.props.history.push('/notes/' + id)
  }
}