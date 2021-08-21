package service

import (
	"friend-management-v1/model"
)

type RelationService interface {
	GetFriendsEmail(rq model.GetFriendsRequest) ([]string, error)
	Addfriend(rq model.AddAndGetCommonRequest) (bool, error)
	GetCommonFriends(rq model.AddAndGetCommonRequest) ([]string, error)
	SubcribeToEmail(rq model.SubcribeAndBlockRequest) (bool, error)
	BlockEmail(rq model.SubcribeAndBlockRequest) (bool, error)
	RetrieveContactEmail(rq model.RetrieveRequest) ([]string, error)
}
