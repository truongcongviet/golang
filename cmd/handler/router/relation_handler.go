package router

import (
	"friend-management-v1/internal/service"
	"friend-management-v1/internal/utils"
	"friend-management-v1/model"

	"net/http"

	"encoding/json"
)

type RelationHandler struct {
	service service.RelationService
}

func (h *RelationHandler) GetFriendsEmail(w http.ResponseWriter, r *http.Request) {
	var request model.GetFriendsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if !utils.IsEmailValid(request.Email) {
			response := model.NewErrorResponse("Invalid email format")
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		email, err := h.service.GetFriendsEmail(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		response := model.AddAndGetResponse{
			Success: true,
			Friends: email,
			Count:   len(email),
		}
		respondwithJSON(w, http.StatusOK, response)
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
	}

}

func (h *RelationHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	var request model.AddAndGetCommonRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if err := utils.ValidateAddComonRequest(request); err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		_, err := h.service.Addfriend(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		respondwithJSON(w, http.StatusOK, model.SuccessRespone{Success: true})
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
		return
	}

}

func (h *RelationHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var request model.AddAndGetCommonRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if err := utils.ValidateAddComonRequest(request); err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		friends, err := h.service.GetCommonFriends(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		response := model.AddAndGetResponse{
			Success: true,
			Friends: friends,
			Count:   len(friends),
		}
		respondwithJSON(w, http.StatusOK, response)
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
		return
	}
}

func (h *RelationHandler) SubcribeToEmail(w http.ResponseWriter, r *http.Request) {
	var request model.SubcribeAndBlockRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if err := utils.ValidateSubcribeAndBlockRequest(request); err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		_, err := h.service.SubcribeToEmail(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		respondwithJSON(w, http.StatusOK, model.SuccessRespone{Success: true})
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
		return
	}
}

func (h *RelationHandler) BlockEmail(w http.ResponseWriter, r *http.Request) {
	var request model.SubcribeAndBlockRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if err := utils.ValidateSubcribeAndBlockRequest(request); err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		_, err := h.service.BlockEmail(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		respondwithJSON(w, http.StatusOK, model.SuccessRespone{Success: true})
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
		return
	}
}

func (h *RelationHandler) GetRetrivableEmails(w http.ResponseWriter, r *http.Request) {
	var request model.RetrieveRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err == nil {
		if err := utils.ValidateRetrieveRequest(request); err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		recipients, err := h.service.RetrieveContactEmail(request)
		if err != nil {
			response := model.NewErrorResponse(err.Error())
			respondWithError(w, http.StatusBadRequest, response)
			return
		}
		response := model.RetrieveResponse{
			Success:    true,
			Recipients: recipients,
		}
		respondwithJSON(w, http.StatusOK, response)
	} else {
		response := model.NewErrorResponse(err.Error())
		respondWithError(w, http.StatusInternalServerError, response)
		return
	}
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	respondwithJSON(w, code, payload)
}
