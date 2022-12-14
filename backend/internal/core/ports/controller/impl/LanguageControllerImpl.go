package impl

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/internal/core/ports/utils"
	"uniovi-localizeme/internal/core/service"
	"uniovi-localizeme/internal/core/service/impl"
	"uniovi-localizeme/tools"
)

type LanguageControllerImpl struct {
	languageService service.LanguageService
	userService     service.UserService
}

func CreateLanguageController() *LanguageControllerImpl {
	return &LanguageControllerImpl{impl.CreateLanguageService(), impl.CreateUserService()}
}

// swagger:route POST /languages Languages CreateLanguage
// Create a new language.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Language
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
// - 422: ErrorDto
func (l LanguageControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, l.userService); user == nil {
		return
	}
	var request dto.LanguageDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	language, err := l.languageService.Create(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, language)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /languages/{id} Languages DeleteLanguage
// Return a language by id.
//
// Responses:
// - 200: description:bool
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (l LanguageControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsAdmin(w, r, l.userService)
	if user == nil {
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		err := errors.New(constants.IdNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	result, err := l.languageService.Delete(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, result)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PATCH /languages/{id} Languages DisableLanguage
// Disable of a language.
//
// Responses:
// - 200: Language
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (l LanguageControllerImpl) Disable(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, l.userService)
	if isAdmin == nil {
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		err := errors.New(constants.IdNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	stage, err := l.languageService.Disable(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /languages Languages FindAllLanguages
// Return all languages.
//
// Responses:
// - 200: []Language
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (l LanguageControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, l.userService); user == nil {
		return
	}
	stages, err := l.languageService.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stages)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /language Languages UpdateLanguage
// Update the information of a language.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Language
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
// - 422: ErrorDto
func (l LanguageControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, l.userService)
	if isAdmin == nil {
		return
	}
	var request domain.Language
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := l.languageService.Update(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
