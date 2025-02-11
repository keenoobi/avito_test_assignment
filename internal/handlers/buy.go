package handlers

import (
	"avito-merch/internal/usecase"
	"avito-merch/pkg/context"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type BuyHandler struct {
	buyUseCase *usecase.BuyUseCase
}

func NewBuyHandler(buyUseCase *usecase.BuyUseCase) *BuyHandler {
	return &BuyHandler{buyUseCase: buyUseCase}
}

func (h *BuyHandler) BuyItem(w http.ResponseWriter, r *http.Request) {
	// Получаем userID из контекста
	userID, ok := context.GetUserID(r.Context())
	if !ok {
		slog.Error("User ID not found in context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Получаем название товара из URL
	itemName := mux.Vars(r)["item"]

	// Выполняем покупку
	if err := h.buyUseCase.BuyItem(userID, itemName); err != nil {
		slog.Error("Failed to buy item", "userID", userID, "item", itemName, "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Item purchased successfully"))
}
