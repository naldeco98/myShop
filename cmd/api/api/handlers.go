package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nialdeco98/myShop/internal/cards"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
}

func (app *application) PaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		app.ErrorLog.Println(err)
		return
	}
	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		w.WriteHeader(400)
		app.ErrorLog.Println(err)
		return
	}

	card := cards.Card{
		Secret:   app.Config.Stripe.Secret,
		Key:      app.Config.Stripe.Key,
		Currency: payload.Currency,
	}

	okay := true
	pi, msg, err := card.Charge(payload.Currency, amount)
	if err != nil {
		okay = false
	}
	if okay {
		out, err := json.MarshalIndent(pi, "", "\t")
		if err != nil {
			w.WriteHeader(500)
			app.ErrorLog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application-json")
		w.Write(out)
	} else {
		j := jsonResponse{
			OK:      false,
			Message: msg,
		}
		out, err := json.MarshalIndent(j, "", "\t")
		if err != nil {
			w.WriteHeader(500)
			app.ErrorLog.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application-json")
		w.Write(out)
	}

}
