package handler

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	pb "notes/proto/notes"

	"github.com/google/uuid"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2/store"
)

// ServiceName is the identifier of the service
const ServiceName = "go.micro.srv.distributed.notes"

// NewHandler returns an initialized Handler
func NewHandler() *Handler {
	s := store.DefaultStore

	// store namespace can only contain letters
	// todo: move this to cockroach.configure() method
	namespace := strings.ReplaceAll(ServiceName, ".", "")
	s.Init(store.Namespace(namespace))

	return &Handler{store: s}
}

// Handler imlements the notes proto definition
type Handler struct {
	store store.Store
}

// Create inserts a new note in the store
func (h *Handler) Create(ctx context.Context, req *pb.CreateNoteRequest, rsp *pb.CreateNoteResponse) error {
	// generate a key (uuid v4)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	// set the generated fields on the note
	note := req.Note
	note.Id = id.String()
	note.Created = time.Now().Unix()

	// encode the message as json
	bytes, err := json.Marshal(req.Note)
	if err != nil {
		return err
	}

	// write to the store
	err = h.store.Write(&store.Record{Key: note.Id, Value: bytes})
	if err != nil {
		return err
	}

	// return the note in the response
	rsp.Note = note
	return nil
}

// Update is a client streaming RPC which streams update events from the client which are used
// to update the note in the store
func (h *Handler) Update(ctx context.Context, stream pb.Notes_UpdateStream) error {
	for {
		// Get a request from the stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		// Validate the request
		if len(req.Note.Id) == 0 {
			return errors.BadRequest(ServiceName, "Missing Note ID")
		}

		// Lookup the note from the store
		recs, err := h.store.Read(req.Note.Id)
		if err != nil {
			return errors.InternalServerError(ServiceName, "Error reading from store: %v", err.Error())
		}
		if len(recs) == 0 {
			return errors.NotFound(ServiceName, "Note not found")
		}

		// Decode the note
		var note *pb.Note
		if err := json.Unmarshal(recs[0].Value, &note); err != nil {
			return errors.InternalServerError(ServiceName, "Error unmarshaling JSON: %v", err.Error())
		}

		// Update the notes title and text
		note.Title = req.Note.Title
		note.Text = req.Note.Text

		// Remarshal the note into bytes
		bytes, err := json.Marshal(note)
		if err != nil {
			return errors.InternalServerError(ServiceName, "Error marshaling JSON: %v", err.Error())
		}

		// Write the updated note to the store
		err = h.store.Write(&store.Record{Key: note.Id, Value: bytes})
		if err != nil {
			return errors.InternalServerError(ServiceName, "Error writing to store: %v", err.Error())
		}
	}

	return nil
}

// Delete removes the note from the store, looking up using ID
func (h *Handler) Delete(ctx context.Context, req *pb.DeleteNoteRequest, rsp *pb.DeleteNoteResponse) error {
	// Validate the request
	if len(req.Note.Id) == 0 {
		return errors.BadRequest(ServiceName, "Missing Note ID")
	}

	// Delete the note using ID and return the error
	return h.store.Delete(req.Note.Id)
}

// List returns all of the notes in the store
func (h *Handler) List(ctx context.Context, req *pb.ListNotesRequest, rsp *pb.ListNotesResponse) error {
	// Retrieve all of the records in the store
	recs, err := h.store.List()
	if err != nil {
		return errors.InternalServerError(ServiceName, "Error reading from store: %v", err.Error())
	}

	// Initialize the response notes slice
	rsp.Notes = make([]*pb.Note, len(recs))

	// Unmarshal the notes into the response
	for i, r := range recs {
		if err := json.Unmarshal(r.Value, &rsp.Notes[i]); err != nil {
			return errors.InternalServerError(ServiceName, "Error unmarshaling json: %v", err.Error())
		}
	}

	return nil
}
