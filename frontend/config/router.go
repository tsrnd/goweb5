package config

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/tsrnd/goweb5/frontend/services/cache"
	threadDeliver "github.com/tsrnd/goweb5/frontend/thread/delivery/http"
	threadRepo "github.com/tsrnd/goweb5/frontend/thread/repository/psql"
	threadCase "github.com/tsrnd/goweb5/frontend/thread/usecase"
	userDeliver "github.com/tsrnd/goweb5/frontend/user/delivery/http"
	userRepo "github.com/tsrnd/goweb5/frontend/user/repository/psql"
	userCase "github.com/tsrnd/goweb5/frontend/user/usecase"
)

// Router func
func Router(db *sql.DB, c cache.Cache) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	addUserRoutes(r, db, c)
	addThreadRoutes(r, db, c)
	return r
}
func addUserRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

func addThreadRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repoThread := threadRepo.NewThreadRepository(db)
	repoUser := userRepo.NewUserRepository(db)
	threadUC := threadCase.NewThreadUsecase(repoThread)
	userUC := userCase.NewUserUsecase(repoUser)
	threadDeliver.NewThreadController(r, threadUC, userUC, c)
}
