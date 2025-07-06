package get

import (
	"encoding/json"
	"net/http"

	"github.com/dprio/currency-exchange/server/internal/usecase/getdollarquote"
	"github.com/dprio/currency-exchange/server/pkg/handler"
)

type Handler struct {
	getDollarQuoteUseCase getdollarquote.UseCase
}

func New(usecase getdollarquote.UseCase) handler.HandlerInterface {
	return &Handler{
		getDollarQuoteUseCase: usecase,
	}
}

func (h *Handler) Serve(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dollarQuote, err := h.getDollarQuoteUseCase.Execute(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := newResponse(dollarQuote)

	json.NewEncoder(w).Encode(response)

}
