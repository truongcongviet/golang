package router

import (
	"bytes"
	"errors"
	"fmt"
	"friend-management-v1/model/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetFriendBlock(t *testing.T) {
	jsonStr := []byte(`{"email" : "quan@gmail.com"}`)
	jsonStr2 := []byte(`{"email" : "quangmail.com"}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  []string
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:       "Get friend email succeed",
			statusCode: http.StatusOK,
			mockResponse: []string{
				"quan@gmail.com",
				"hau@gmail.com",
			},
			requestBody: bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{
								"success": true,
								"friends": [
									"quan@gmail.com",
									"hau@gmail.com"
								],
								"count": 2
							}`),
			err: nil,
		},
		{
			name:         "Get friend email not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: nil,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Get friends invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonStr2),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "Invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Get friends invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("GetFriendsEmail", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/friends", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/friends", func(w http.ResponseWriter, r *http.Request) {
				handler.GetFriendsEmail(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func TestAddFriendBlock(t *testing.T) {
	jsonStr := []byte(`{
		"friends" : [
			"quan@gmail.com",
			"quang@gmail.com"
		]
	}`)
	jsonStr2 := []byte(`{
		"friends" : [
			"quan12ytgmail.com",
			"quang@gmail.com"
		]
	}`)
	jsonLackEmail := []byte(`{
		"friends" : [
			"quang@gmail.com"
		]
	}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  bool
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:         "Add friend succeed",
			statusCode:   http.StatusOK,
			mockResponse: true,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{ "success": true }`),
			err:          nil,
		},
		{
			name:         "Add friend email not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: false,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Add friend invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonStr2),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Add friend lack email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonLackEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "must contain 2 emails",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Add friend invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("Addfriend", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/add", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/add", func(w http.ResponseWriter, r *http.Request) {
				handler.AddFriend(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func TestGetCommonFriendBlock(t *testing.T) {
	jsonStr := []byte(`{
		"friends" : [
			"quan@gmail.com",
			"quang@gmail.com"
		]
	}`)
	jsonInvalidEmail := []byte(`{
		"friends" : [
			"quan12ytgmail.com",
			"quang@gmail.com"
		]
	}`)
	jsonLackEmail := []byte(`{
		"friends" : [
			"quang@gmail.com"
		]
	}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  []string
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:       "Get common friends succeed",
			statusCode: http.StatusOK,
			mockResponse: []string{
				"quan@gmail.com",
				"hau@gmail.com",
			},
			requestBody: bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{
								"success": true,
								"friends": [
									"quan@gmail.com",
									"hau@gmail.com"
								],
								"count": 2
							}`),
			err: nil,
		},
		{
			name:         "Get common friends not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: nil,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Get common friends invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonInvalidEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Get common friends lack email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonLackEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "must contain 2 emails",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Get common friends invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("GetCommonFriends", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/common", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/common", func(w http.ResponseWriter, r *http.Request) {
				handler.GetCommonFriends(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func TestSubcribeToEmailBlock(t *testing.T) {
	jsonStr := []byte(`{
						"requestor": "quan@gmail.com",
						"target": "quang@gmail.com"
				}`)
	jsonStrInvalidEmail := []byte(`{
						"requestor": "quangmail.com",
						"target": "quang@gmail.com"
				}`)
	jsonEmptyRequest := []byte(`{
					"requestor": "",
					"target": "quang@gmail.com"
			}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  bool
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:         "Subcribe succeed",
			statusCode:   http.StatusOK,
			mockResponse: true,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{ "success": true }`),
			err:          nil,
		},
		{
			name:         "Subcribe email not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: false,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Subcribe invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonStrInvalidEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Subcribe empty request",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonEmptyRequest),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "requestor and target must not be empty",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Subcribe invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("SubcribeToEmail", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/subcribe", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/subcribe", func(w http.ResponseWriter, r *http.Request) {
				handler.SubcribeToEmail(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func TestBlockEmailBlock(t *testing.T) {
	jsonStr := []byte(`{
						"requestor": "quan@gmail.com",
						"target": "quang@gmail.com"
				}`)
	jsonStrInvalidEmail := []byte(`{
						"requestor": "quangmail.com",
						"target": "quang@gmail.com"
				}`)
	jsonEmptyRequest := []byte(`{
					"requestor": "",
					"target": "quang@gmail.com"
			}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  bool
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:         "Block succeed",
			statusCode:   http.StatusOK,
			mockResponse: true,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{ "success": true }`),
			err:          nil,
		},
		{
			name:         "Block email not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: false,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Block invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonStrInvalidEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Block empty request",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonEmptyRequest),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "requestor and target must not be empty",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Block invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("BlockEmail", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/block", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/block", func(w http.ResponseWriter, r *http.Request) {
				handler.BlockEmail(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func TestRetrieveFriendBlock(t *testing.T) {
	jsonStr := []byte(`{
		"sender": "quan12yt@gmail.com",
		"text" : "ahdad la@gmail.com ahdad la@gmail.com ahdad la@gmail.com"
	}`)
	jsonInvalidEmail := []byte(`{
		"sender": "quan12ytgmail.com",
		"text" : "ahdad la@gmail.com ahdad la@gmail.com ahdad la@gmail.com"
	}}`)
	jsonEmptyRequest := []byte(`{
		"sender": "",
		"text" : "ahdad la@gmail.com ahdad la@gmail.com ahdad la@gmail.com"
	}}`)
	current := time.Now().Format("2006-01-02 15:04:05")

	testCases := []struct {
		name          string
		statusCode    int
		mockResponse  []string
		requestBody   *bytes.Buffer
		err           error
		jsonResponse  string
		errorResponse string
	}{
		{
			name:       "Retrieve contact succeed",
			statusCode: http.StatusOK,
			mockResponse: []string{
				"quan@gmail.com",
				"hau@gmail.com",
			},
			requestBody: bytes.NewBuffer(jsonStr),
			jsonResponse: string(`{
								"success": true,
								"recipients": [
									"quan@gmail.com",
									"hau@gmail.com"
								]
							}`),
			err: nil,
		},
		{
			name:         "Retrieve email not exist",
			statusCode:   http.StatusBadRequest,
			mockResponse: nil,
			requestBody:  bytes.NewBuffer(jsonStr),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "email: quan@gmail.com is not exist in database",
									"timestamp": "%s"
								}`, current),
			err: errors.New("email: quan@gmail.com is not exist in database"),
		},
		{
			name:        "Retrieve invalid email",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonInvalidEmail),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "invalid email format",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Retrieve invalid request",
			statusCode:  http.StatusInternalServerError,
			requestBody: bytes.NewBuffer(nil),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "EOF",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
		{
			name:        "Retrieve empty request",
			statusCode:  http.StatusBadRequest,
			requestBody: bytes.NewBuffer(jsonEmptyRequest),
			jsonResponse: fmt.Sprintf(`{
									"success": false,
									"text": "sender must not empty",
									"timestamp": "%s"
								}`, current),
			err: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(mocks.RelationService)
			handler := RelationHandler{
				service: mockService,
			}
			mockService.On("RetrieveContactEmail", mock.Anything).Return(tc.mockResponse, tc.err)
			request, er := http.NewRequest("POST", "/api/retrieve", tc.requestBody)
			checkError(er, t)

			chi := chi.NewRouter()
			rr := httptest.NewRecorder()
			chi.Post("/api/retrieve", func(w http.ResponseWriter, r *http.Request) {
				handler.GetRetrivableEmails(w, r)
			})
			chi.ServeHTTP(rr, request)
			assert.Equal(t, tc.statusCode, rr.Code)
			assert.JSONEq(t, tc.jsonResponse, rr.Body.String())
		})
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}
