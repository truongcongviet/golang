package service

import (
	"errors"
	"friend-management-v1/internal/utils"
	"friend-management-v1/model"
	"friend-management-v1/model/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFriendsBlock(t *testing.T) {
	request := model.GetFriendsRequest{
		Email: "quan12yt@gmail.com",
	}
	expect := []string{
		"quan12yt@gmail.com",
	}

	testCases := []struct {
		name         string
		mockId       string
		mockResponse []string
		err          error
		finalErr     error
	}{
		{
			name:         "Get Friends succeed",
			mockId:       "1",
			mockResponse: expect,
			err:          nil,
		},
		{
			name:         "Get Friends email not exist",
			mockId:       "1",
			mockResponse: nil,
			err:          errors.New("email: quan12yt@gmail.com is not exist in database"),
			finalErr:     errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:         "Get Friends email failed",
			mockId:       "1",
			mockResponse: nil,
			err:          nil,
			finalErr:     errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.err)
			mockRepo.On("GetEmailByStatus", mock.Anything, mock.Anything).Return(tc.mockResponse, tc.finalErr)

			actual, err := service.GetFriendsEmail(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, tc.mockResponse, actual)
		})
	}
}

func TestAddFriendBlock(t *testing.T) {
	request := model.AddAndGetCommonRequest{
		Friends: []string{
			"quan12yt@gmail.com",
			"quang@gmail.com",
		},
	}

	testCases := []struct {
		name         string
		mockId       string
		checkExist   bool
		mockResponse bool
		getIdError   error
		finalErr     error
	}{
		{
			name:         "Add succeed",
			getIdError:   nil,
			mockId:       "1",
			checkExist:   false,
			mockResponse: true,
			finalErr:     nil,
		},
		{
			name:       "Add email not exist",
			mockId:     "1",
			getIdError: errors.New("email: quan12yt@gmail.com is not exist in database"),
			finalErr:   errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:       "Add already friend",
			getIdError: nil,
			checkExist: true,
			mockId:     "1",
			finalErr:   errors.New("2 emails are already being friend"),
		},
		{
			name:       "Add failed",
			getIdError: nil,
			checkExist: false,
			mockId:     "1",
			finalErr:   errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.getIdError)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.getIdError)
			mockRepo.On("CheckIfExist", mock.Anything, mock.Anything, mock.Anything).Return(tc.checkExist)
			mockRepo.On("AddRelation", mock.Anything, mock.Anything).Return(tc.mockResponse, tc.finalErr)

			actual, err := service.Addfriend(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, tc.mockResponse, actual)
		})
	}
}

func TestGetCommonFriendBlock(t *testing.T) {
	request := model.AddAndGetCommonRequest{
		Friends: []string{
			"quan12yt@gmail.com",
			"quang@gmail.com",
		},
	}
	slice1 := []string{
		"ad@gmail.com",
		"test@gmail.com",
	}
	slice2 := []string{
		"hau@gmail.com",
		"test@gmail.com",
	}
	expect := utils.RetainSlices(slice1, slice2)

	testCases := []struct {
		name           string
		mockId         string
		mockResponse1  []string
		mockResponse2  []string
		expectResponse []string
		getIdError     error
		finalErr       error
	}{
		{
			name:           "Get common  succeed",
			getIdError:     nil,
			mockId:         "1",
			mockResponse1:  slice1,
			mockResponse2:  slice2,
			expectResponse: expect,
			finalErr:       nil,
		},
		{
			name:           "Get common email not exist",
			mockId:         "1",
			expectResponse: nil,
			getIdError:     errors.New("email: quan12yt@gmail.com is not exist in database"),
			finalErr:       errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:           "Get common  failed",
			getIdError:     nil,
			mockId:         "1",
			expectResponse: nil,
			finalErr:       errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", request.Friends[0]).Return(tc.mockId, tc.getIdError)
			mockRepo.On("GetIdFromEmail", request.Friends[1]).Return("2", tc.getIdError)
			mockRepo.On("GetEmailByStatus", "1", mock.Anything).Return(tc.mockResponse1, tc.finalErr)
			mockRepo.On("GetEmailByStatus", "2", mock.Anything).Return(tc.mockResponse2, tc.finalErr)

			actual, err := service.GetCommonFriends(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, tc.expectResponse, actual)
		})
	}
}

func TestSubcribeBlock(t *testing.T) {
	request := model.SubcribeAndBlockRequest{
		Requestor: "quan12yt@gmail.com",
		Target:    "quang@gmail.com",
	}

	testCases := []struct {
		name           string
		mockId         string
		expectResponse bool
		isBlock        bool
		isFriend       bool
		isSubcribe     bool
		getIdError     error
		finalErr       error
	}{
		{
			name:       "Subcribe succeed",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			isFriend:   false,
			isSubcribe: false,
			finalErr:   nil,
		},
		{
			name:       "Subcribe email not exist",
			getIdError: errors.New("email: quan12yt@gmail.com is not exist in database"),
			mockId:     "1",
			finalErr:   errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:       "Subcribe blocked",
			getIdError: nil,
			mockId:     "1",
			isBlock:    true,
			isFriend:   false,
			isSubcribe: false,
			finalErr:   errors.New("target email has been blocked"),
		},
		{
			name:       "Subcribe friend",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			isFriend:   true,
			isSubcribe: false,
			finalErr:   errors.New("already being friend to the target email, no need to subcribe"),
		},
		{
			name:       "Subcribe already subcribe",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			isFriend:   false,
			isSubcribe: true,
			finalErr:   errors.New("already subcribe to the target email"),
		},
		{
			name:       "Subcribe failed",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			isFriend:   false,
			isSubcribe: false,
			finalErr:   errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.getIdError)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return("2", tc.getIdError)
			mockRepo.On("CheckIfExist", mock.Anything, mock.Anything, "BLOCK").Return(tc.isBlock)
			mockRepo.On("CheckIfExist", mock.Anything, mock.Anything, "FRIEND").Return(tc.isFriend)
			mockRepo.On("CheckIfExist", mock.Anything, mock.Anything, "SUBCRIBE").Return(tc.isSubcribe)
			mockRepo.On("AddRelation", mock.Anything, "SUBCRIBE").Return(tc.expectResponse, tc.finalErr)

			actual, err := service.SubcribeToEmail(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, tc.expectResponse, actual)
		})
	}
}

func TestBlockEmailBlock(t *testing.T) {
	request := model.SubcribeAndBlockRequest{
		Requestor: "quan12yt@gmail.com",
		Target:    "quang@gmail.com",
	}

	testCases := []struct {
		name           string
		mockId         string
		expectResponse bool
		isBlock        bool
		getIdError     error
		finalErr       error
	}{
		{
			name:       "Block succeed",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			finalErr:   nil,
		},
		{
			name:       "Block email not exist",
			getIdError: errors.New("email: quan12yt@gmail.com is not exist in database"),
			mockId:     "1",
			finalErr:   errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:       "Block already blocked",
			getIdError: nil,
			mockId:     "1",
			isBlock:    true,
			finalErr:   errors.New("target email has already being blocked"),
		},
		{
			name:       "Block failed",
			getIdError: nil,
			mockId:     "1",
			isBlock:    false,
			finalErr:   errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.getIdError)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return("2", tc.getIdError)
			mockRepo.On("CheckIfExist", mock.Anything, mock.Anything, "BLOCK").Return(tc.isBlock)
			mockRepo.On("AddRelation", mock.Anything, "BLOCK").Return(tc.expectResponse, tc.finalErr)

			actual, err := service.BlockEmail(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, tc.expectResponse, actual)
		})
	}
}

func TestRetrieveBlock(t *testing.T) {
	request := model.RetrieveRequest{
		Sender: "quan12yt@gmail.com",
		Text:   "asd hau@gmail.com",
	}

	emails := []string{
		"asd@gmail.com",
		"test@gmail.com",
	}

	expect := []string{
		"asd@gmail.com",
		"test@gmail.com",
		"hau@gmail.com",
	}
	testCases := []struct {
		name           string
		mockId         string
		mockResponse   []string
		expectResponse []string
		err            error
		finalErr       error
	}{
		{
			name:           "Retrieve succeed",
			mockId:         "1",
			mockResponse:   emails,
			expectResponse: expect,
			err:            nil,
		},
		{
			name:           "Retrieve email not exist",
			mockId:         "1",
			mockResponse:   nil,
			expectResponse: nil,
			err:            errors.New("email: quan12yt@gmail.com is not exist in database"),
			finalErr:       errors.New("email: quan12yt@gmail.com is not exist in database"),
		},
		{
			name:           "Retrieve email not exist",
			mockId:         "1",
			mockResponse:   nil,
			expectResponse: nil,
			err:            nil,
			finalErr:       errors.New(""),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.RelationRepo)
			service := NewRelationService(mockRepo)
			mockRepo.On("GetIdFromEmail", mock.Anything).Return(tc.mockId, tc.err)
			mockRepo.On("GetRetrivableEmails", mock.Anything).Return(tc.mockResponse, tc.finalErr)

			actual, err := service.RetrieveContactEmail(request)

			assert.Equal(t, tc.finalErr, err)
			assert.Equal(t, len(tc.expectResponse), len(actual))
		})
	}
}
