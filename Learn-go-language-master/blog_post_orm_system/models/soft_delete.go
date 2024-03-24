package models

import "time"

type SoftDelete struct {
	CreatedAt time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
	DeletedAt time.Time `bun:"deleted_at" json:"deleted_at"`
}