package model

import "time"

type Email struct {
	EmailId int8   `json:"email_id" binding:"required"`
	Email   string `json:"name" binding:"required"`
}

type Friend_Relation struct {
	RelationId int8   `json:"relation_id" binding:"required"`
	YourId     int8   `json:"your_id" binding:"required"`
	FriendId   int8   `json:"friend_id" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type AddAndGetCommonRequest struct {
	Friends []string `json:"friends" binding:"required"`
}

type SuccessRespone struct {
	Success bool `json:"success" binding:"required"`
}

type GetFriendsRequest struct {
	Email string `json:"email" binding:"required"`
}

type AddAndGetResponse struct {
	Success bool     `json:"success" binding:"required"`
	Friends []string `json:"friends" binding:"required"`
	Count   int      `json:"count" binding:"required"`
}

type SubcribeAndBlockRequest struct {
	Requestor string `json:"requestor" binding:"required"`
	Target    string `json:"target" binding:"required"`
}

type RetrieveRequest struct {
	Sender string `json:"sender" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type RetrieveResponse struct {
	Success    bool     `json:"success" binding:"required"`
	Recipients []string `json:"recipients" binding:"required"`
}

type ErrorResponse struct {
	Success   bool   `json:"success" binding:"required"`
	Error     string `json:"text" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
}

func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Success:   false,
		Error:     err,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
}
