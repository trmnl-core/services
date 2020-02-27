package handler

import (
	"context"
	"io"

	"github.com/micro/go-micro/v2/errors"
	pb "github.com/micro/services/distributed/api/proto"
	notes "github.com/micro/services/notes/proto"
)

// CreateNote creates a new note in the notes service
func (h *Handler) CreateNote(ctx context.Context, req *pb.CreateNoteRequest, rsp *pb.CreateNoteResponse) error {
	if req.Note == nil {
		return errors.BadRequest("go.micro.api.distributed", "Note Required")
	}

	resp, err := h.notes.Create(ctx, &notes.CreateNoteRequest{Note: deserializeNote(req.Note)})
	if err != nil {
		return err
	}

	rsp.Note = serializeNote(resp.Note)
	return nil
}

// UpdateNote streams updates to the notes service
func (h *Handler) UpdateNote(ctx context.Context, stream pb.Distributed_UpdateNoteStream) error {
	client, err := h.notes.Update(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if req.Note == nil {
			return errors.BadRequest("go.micro.api.distributed", "Note Required")
		}

		err = client.Send(&notes.UpdateNoteRequest{Note: deserializeNote(req.Note)})
		if err != nil {
			return err
		}
	}
}

// DeleteNote note deleted a note in the notes service
func (h *Handler) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest, rsp *pb.DeleteNoteResponse) error {
	if req.Note == nil {
		return errors.BadRequest("go.micro.api.distributed", "Note Required")
	}

	note := &notes.Note{Id: req.Note.Id}
	_, err := h.notes.Delete(ctx, &notes.DeleteNoteRequest{Note: note})
	return err
}

// ListNotes returns all the notes from the notes service
func (h *Handler) ListNotes(ctx context.Context, req *pb.ListNotesRequest, rsp *pb.ListNotesResponse) error {
	resp, err := h.notes.List(ctx, &notes.ListNotesRequest{})
	if err != nil {
		return err
	}

	rsp.Notes = make([]*pb.Note, len(resp.Notes))
	for i, n := range resp.Notes {
		rsp.Notes[i] = serializeNote(n)
	}

	return nil
}

func serializeNote(n *notes.Note) *pb.Note {
	return &pb.Note{
		Id:      n.Id,
		Title:   n.Title,
		Text:    n.Text,
		Created: n.Created,
	}
}

func deserializeNote(n *pb.Note) *notes.Note {
	return &notes.Note{
		Id:    n.Id,
		Title: n.Title,
		Text:  n.Text,
	}
}
