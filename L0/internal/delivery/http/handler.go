package http

import (
    "L0/internal/usecase"
    "fmt"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "strings"
)

const host = "localhost"

type Handler struct {
    useCase *usecase.OrderUseCase
}

func NewHandler(uc *usecase.OrderUseCase) *Handler {
    return &Handler{useCase: uc}
}

func (h *Handler) StartServer(port string) {
    router := mux.NewRouter()
    router.HandleFunc("/orders/{id}", h.getOrderHandler).Methods("GET")
    err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
    if err != nil {
        return
    }
}

func (h *Handler) getOrderHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    orderID := vars["id"]
    if orderID == "" {
        http.Error(w, "Missing order ID", http.StatusBadRequest)
        return
    }
    order, err := h.useCase.GetOrder(orderID)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            http.Error(w, "Order not found", http.StatusNotFound)
        } else {
            log.Printf("Failed to get order: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    tmpl, err := template.ParseFiles("internal/delivery/templates/order.html")
    if err != nil {
        log.Printf("Failed to parse template: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, order); err != nil {
        log.Printf("Failed to execute template: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}
