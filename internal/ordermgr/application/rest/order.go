package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr/domain/service"
)

type orderRestEndpoint struct {
	orderService service.OrderService
}

func NewOrderRestEndpoint(orderService service.OrderService) *orderRestEndpoint {
	return &orderRestEndpoint{
		orderService: orderService,
	}
}

func (o *orderRestEndpoint) BuildRoutes(r mux.Router) {
	r.Path("/v1/").Methods(http.MethodGet).HandlerFunc(o.List)
}

func (o *orderRestEndpoint) List(w http.ResponseWriter, r *http.Request) {
	orders, err := o.orderService.List(r.Context())
	if err != nil {
		log.Printf("Failed to get order list: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	reqBody, err := json.Marshal(orders)
	if err != nil {
		log.Printf("Failed to marshal orders into JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	w.Write(reqBody)
	return
}
