package playlists

import (
	"fmt"
	"myuseek/business/playlists"
	controllers "myuseek/controller"
	"myuseek/controller/playlists/requests"
	"myuseek/controller/playlists/responses"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type PlaylistController struct {
	PlaylistUseCase playlists.Usecase
}

func NewPlaylistController(playlistUseCase playlists.Usecase) *PlaylistController {
	return &PlaylistController{
		PlaylistUseCase: playlistUseCase}
}

func (playlistController PlaylistController) Create(c echo.Context) error {
	fmt.Println("Create Playlist")
	playlistCreate := requests.PlaylistCreate{}
	c.Bind(&playlistCreate)

	ctx := c.Request().Context()
	playlist, error := playlistController.PlaylistUseCase.Create(ctx, playlistCreate.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(playlist))
}

func (playlistController PlaylistController) AddSong(c echo.Context) error {
	fmt.Println("AddSong to playlist")
	//specify playlist id
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	//binding song
	songtoadd := requests.NewSong{}
	c.Bind(&songtoadd)

	songtoadd.Id = id
	ctx := c.Request().Context()
	playlistdomain, error := playlistController.PlaylistUseCase.AddSong(ctx, songtoadd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(playlistdomain))
}

func (playlistController PlaylistController) GetbyID(c echo.Context) error {
	fmt.Println("GetPlaylists by ID")
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	playlistByID := requests.PlaylistByID{}
	playlistByID.Id = id
	ctx := c.Request().Context()
	playlistdomain, error := playlistController.PlaylistUseCase.GetbyID(ctx, playlistByID.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(playlistdomain))
}

func (playlistController PlaylistController) GetPlaylists(c echo.Context) error {
	fmt.Println("GetPlaylists")

	ctx := c.Request().Context()
	playlistlistdomain, error := playlistController.PlaylistUseCase.GetPlaylists(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(playlistlistdomain))
}
