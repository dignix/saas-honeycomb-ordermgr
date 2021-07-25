package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dignix/saas-honeycomb-ordermgr/internal/ordermgr"
)

type OrderService interface {
	Get(ctx context.Context, id string) (*ordermgr.Order, error)
	Create(ctx context.Context, order *ordermgr.Order) (*ordermgr.Order, error)
}

type orderRestEndpoint struct {
	orderService OrderService
}

func NewOrderRestEndpoint(orderService OrderService) *orderRestEndpoint {
	return &orderRestEndpoint{
		orderService: orderService,
	}
}

func (e *orderRestEndpoint) BuildRoutes(r *mux.Router) {
	r.Path("/v1/orders/{id}").Methods(http.MethodGet).HandlerFunc(e.Get)
	r.Path("/v1/orders").Methods(http.MethodPost).HandlerFunc(e.Store)
}

func (e *orderRestEndpoint) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	order, err := e.orderService.Get(ctx, mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	resp, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	w.Write(resp)
}

func (e *orderRestEndpoint) Store(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var order *ordermgr.Order
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	orders, err := e.orderService.Create(ctx, order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	resp, err := json.Marshal(orders)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("HTTP Error: Bad Request"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
