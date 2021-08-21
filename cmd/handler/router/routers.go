package router

import (
	"friend-management-v1/internal/repos"
	"friend-management-v1/internal/service"
	"friend-management-v1/internal/utils"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetUpRouter() *chi.Mux {
	db := utils.DBConnection()
	relation_repo := repos.NewRelationRepo(db)
	relation_service := service.NewRelationService(relation_repo)
	relation_handler := RelationHandler{
		service: relation_service,
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/api", func(r chi.Router) {
		r.Post("/friends", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.GetFriendsEmail(w, r)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.AddFriend(w, r)
		})
		r.Post("/common", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.GetCommonFriends(w, r)
		})
		r.Post("/subcribe", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.SubcribeToEmail(w, r)
		})
		r.Post("/block", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.BlockEmail(w, r)
		})
		r.Post("/retrieve", func(w http.ResponseWriter, r *http.Request) {
			relation_handler.GetRetrivableEmails(w, r)
		})
	})
	return r
}
