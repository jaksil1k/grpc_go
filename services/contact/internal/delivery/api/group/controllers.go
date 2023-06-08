package group

import (
	"github.com/julienschmidt/httprouter"
	"grpc/services/contact/internal/delivery/api"
	"grpc/services/contact/internal/usecase"
	"net/http"
)

const (
	groupURL = "/groups:group_id"
)

type handler struct {
	groupService usecase.Group
}

func NewHandler(service usecase.Group) api.Handler {
	return &handler{groupService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(groupURL, h.GetById)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//group := h.groupService.GetGroupById(context.Background(), -1)
	w.Write([]byte("group"))
	w.WriteHeader(http.StatusOK)
}
