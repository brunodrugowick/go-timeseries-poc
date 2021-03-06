// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Event struct {
	Uuid        uuid.UUID
	CreatedDate sql.NullInt64
	Description sql.NullString
	Username    sql.NullString
}

type Measurement struct {
	Uuid        uuid.UUID
	CreatedDate sql.NullInt64
	HeartRate   sql.NullInt32
	High        sql.NullInt32
	Low         sql.NullInt32
	Username    sql.NullString
}

type Snapshot struct {
	Uuid        uuid.UUID
	CreatedDate sql.NullInt64
	Description sql.NullString
	EndDate     sql.NullInt64
	IsPublic    bool
	StartDate   sql.NullInt64
	Username    sql.NullString
}

type User struct {
	Email      string
	Enabled    bool
	FullName   sql.NullString
	Provider   sql.NullString
	ProviderID sql.NullString
}
