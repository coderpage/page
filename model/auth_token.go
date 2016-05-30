package model

import (
	"time"
)

type AuthToken struct {
	Value  string    `json:"value"`
	Expire time.Time `json:"expire"`
}
