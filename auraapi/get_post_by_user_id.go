package auraapi

import "aura/auradomain"

type GetPostsByUserIDRes struct {
	Posts []*auradomain.Post `json:"posts"`
}
