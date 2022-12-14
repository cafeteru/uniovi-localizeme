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

type GroupControllerImpl struct {
	groupService service.GroupService
	userService  service.UserService
}

func CreateGroupController() *GroupControllerImpl {
	return &GroupControllerImpl{impl.CreateGroupService(), impl.CreateUserService()}
}

// swagger:route POST /groups Groups CreateGroup
// Create a new group.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (g GroupControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, g.userService); user == nil {
		return
	}
	var groupDto dto.GroupDto
	if err := json.NewDecoder(r.Body).Decode(&groupDto); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	group, err := g.groupService.Create(groupDto)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, group)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /groups/{id} Groups DeleteGroup
// Delete a group by id.
//
// Responses:
// - 200: description:bool
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (g GroupControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, g.userService)
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
	result, err := g.groupService.Delete(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, result)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PATCH /groups/{id} Groups DisableGroup
// Disable of a group.
//
// Responses:
// - 200: Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (g GroupControllerImpl) Disable(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, g.userService)
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
	stage, err := g.groupService.Disable(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /groups Groups FindAllGroups
// Return all groups.
//
// Responses:
// - 200: []Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (g GroupControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, g.userService)
	if user == nil {
		return
	}
	var groups *[]domain.Group
	var err error
	if user.Admin {
		groups, err = g.groupService.FindAll()
	} else {
		groups, err = g.groupService.FindByPermissions(user.ID)
	}
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, groups)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /groups Groups FindAllGroups
// Return all groups.
//
// Responses:
// - 200: []Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (g GroupControllerImpl) FindCanWrite(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, g.userService)
	if user == nil {
		return
	}
	var groups *[]domain.Group
	var err error
	if user.Admin {
		groups, err = g.groupService.FindAll()
	} else {
		groups, err = g.groupService.FindCanWrite(user.ID)
	}
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, groups)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /groups Groups UpdateGroup
// Update the information of a group.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (g GroupControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, g.userService)
	if user == nil {
		return
	}
	var request domain.Group
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	group, err := g.groupService.Update(request, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, group)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
