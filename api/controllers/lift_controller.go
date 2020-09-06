package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skwongg/jackt/api/auth"
	"github.com/skwongg/jackt/api/models"
	"github.com/skwongg/jackt/api/responses"
	"github.com/skwongg/jackt/api/utils/formaterror"
)

//CreateLift creates a lift
func (server *Server) CreateLift(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	lift := models.Lift{}
	err = json.Unmarshal(body, &lift)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	lift.Prepare()
	err = lift.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if uid != lift.AuthorID {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }
	liftCreated, err := lift.SaveLift(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, liftCreated.ID))
	responses.JSON(w, http.StatusCreated, liftCreated)
}

//GetLifts fetches a list of lifts
func (server *Server) GetLifts(w http.ResponseWriter, r *http.Request) {

	lift := models.Lift{}

	lifts, err := lift.FindAllLifts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, lifts)
}

//GetLift retrieves a specific lift
func (server *Server) GetLift(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	lift := models.Lift{}

	liftReceived, err := lift.FindLiftByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, liftReceived)
}

//UpdateLift updates a lift
func (server *Server) UpdateLift(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the lift id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the lift exist
	lift := models.Lift{}
	err = server.DB.Debug().Model(models.Lift{}).Where("id = ?", pid).Take(&lift).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Lift not found"))
		return
	}

	// Read the data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	liftUpdate := models.Lift{}
	err = json.Unmarshal(body, &liftUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	liftUpdate.Prepare()
	err = liftUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	liftUpdate.ID = lift.ID //this is important to tell the model the lift id to update, the other update field are set above

	liftUpdated, err := liftUpdate.UpdateALift(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, liftUpdated)
}

//DeleteLift is to delete lifts that involve crossfit
func (server *Server) DeleteLift(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid lift id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the lift exist
	lift := models.Lift{}
	err = server.DB.Debug().Model(models.Lift{}).Where("id = ?", pid).Take(&lift).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = lift.DeleteALift(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
