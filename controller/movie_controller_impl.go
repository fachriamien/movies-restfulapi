package controller

import (
	"fchramn/movies-restfulapi/helper"
	"fchramn/movies-restfulapi/model/web"
	"fchramn/movies-restfulapi/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MovieControllerImpl struct {
	MovieService service.MovieService
}

func NewMovieController(movieService service.MovieService) MovieController {
	return &MovieControllerImpl{
		MovieService: movieService,
	} //polymorphism
}

func (controller *MovieControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieResponses := controller.MovieService.FindAll(request.Context())
	webRepsonse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponses,
	}

	helper.WriteToResponseBody(writer, webRepsonse)
}

func (controller *MovieControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieId := params.ByName("id")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	movieResponse := controller.MovieService.FindById(request.Context(), id)
	webRepsonse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(writer, webRepsonse)
}

func (controller *MovieControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieCreateRequest := web.MovieCreateRequest{}
	helper.ReadFromRequestBody(request, &movieCreateRequest)

	movieResponse := controller.MovieService.Create(request.Context(), movieCreateRequest)
	webRepsonse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(writer, webRepsonse)
}

func (controller *MovieControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieUpdateRequest := web.MovieUpdateRequest{}
	helper.ReadFromRequestBody(request, &movieUpdateRequest)

	movieId := params.ByName("id")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	movieUpdateRequest.Id = id

	movieResponse := controller.MovieService.Update(request.Context(), movieUpdateRequest)
	webRepsonse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(writer, webRepsonse)
}

func (controller *MovieControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieId := params.ByName("id")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	controller.MovieService.Delete(request.Context(), id)
	webRepsonse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webRepsonse)
}
