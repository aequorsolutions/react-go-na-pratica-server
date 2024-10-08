// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const deleteRoom = `-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1
`

func (q *Queries) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteRoom, id)
	return err
}

const deleteRoomMessages = `-- name: DeleteRoomMessages :exec
DELETE FROM messages
WHERE room_id = $1
`

func (q *Queries) DeleteRoomMessages(ctx context.Context, roomID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteRoomMessages, roomID)
	return err
}

const getMessage = `-- name: GetMessage :one
SELECT
    "id", "room_id", "message", "reaction_count", "answered", "moderated"
FROM messages
WHERE
    id = $1
`

func (q *Queries) GetMessage(ctx context.Context, id uuid.UUID) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Message,
		&i.ReactionCount,
		&i.Answered,
		&i.Moderated,
	)
	return i, err
}

const getRoom = `-- name: GetRoom :one
SELECT
    "id", "theme", "secret"
FROM rooms
WHERE id = $1
`

func (q *Queries) GetRoom(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(&i.ID, &i.Theme, &i.Secret)
	return i, err
}

const getRoomMessages = `-- name: GetRoomMessages :many
SELECT
    "id", "room_id", "message", "reaction_count", "answered", "moderated"
FROM messages
WHERE
    room_id = $1
`

func (q *Queries) GetRoomMessages(ctx context.Context, roomID uuid.UUID) ([]Message, error) {
	rows, err := q.db.Query(ctx, getRoomMessages, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Message,
			&i.ReactionCount,
			&i.Answered,
			&i.Moderated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoomSecret = `-- name: GetRoomSecret :one
SELECT
    "secret"
FROM rooms
WHERE id = $1
`

func (q *Queries) GetRoomSecret(ctx context.Context, id uuid.UUID) (string, error) {
	row := q.db.QueryRow(ctx, getRoomSecret, id)
	var secret string
	err := row.Scan(&secret)
	return secret, err
}

const getRooms = `-- name: GetRooms :many
SELECT
    "id", "theme"
FROM rooms
`

type GetRoomsRow struct {
	ID    uuid.UUID
	Theme string
}

func (q *Queries) GetRooms(ctx context.Context) ([]GetRoomsRow, error) {
	rows, err := q.db.Query(ctx, getRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRoomsRow
	for rows.Next() {
		var i GetRoomsRow
		if err := rows.Scan(&i.ID, &i.Theme); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMessage = `-- name: InsertMessage :one
INSERT INTO messages
    ( "room_id", "message" ) VALUES
    ( $1, $2 )
RETURNING "id"
`

type InsertMessageParams struct {
	RoomID  uuid.UUID
	Message string
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertMessage, arg.RoomID, arg.Message)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertRoom = `-- name: InsertRoom :one
INSERT INTO rooms
    ( "theme", "secret" ) VALUES
    ( $1, $2 )
RETURNING "id"
`

type InsertRoomParams struct {
	Theme  string
	Secret string
}

func (q *Queries) InsertRoom(ctx context.Context, arg InsertRoomParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRoom, arg.Theme, arg.Secret)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const markMessageAsAnswered = `-- name: MarkMessageAsAnswered :exec
UPDATE messages
SET
    answered = true
WHERE
    id = $1
`

func (q *Queries) MarkMessageAsAnswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAsAnswered, id)
	return err
}

const markMessageAsModerated = `-- name: MarkMessageAsModerated :exec
UPDATE messages
SET
    moderated = true
WHERE
    id = $1
`

func (q *Queries) MarkMessageAsModerated(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAsModerated, id)
	return err
}

const reactToMessage = `-- name: ReactToMessage :one
UPDATE messages
SET
    reaction_count = reaction_count + 1
WHERE
    id = $1
RETURNING reaction_count
`

func (q *Queries) ReactToMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reactToMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}

const removeMessageAsModerated = `-- name: RemoveMessageAsModerated :exec
UPDATE messages
SET
    moderated = false
WHERE
    id = $1
`

func (q *Queries) RemoveMessageAsModerated(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeMessageAsModerated, id)
	return err
}

const removeReactionFromMessage = `-- name: RemoveReactionFromMessage :one
UPDATE messages
SET
    reaction_count = reaction_count - 1
WHERE
    id = $1
RETURNING reaction_count
`

func (q *Queries) RemoveReactionFromMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, removeReactionFromMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}
