package app

import (
	"net/http"

	"github.com/nialdeco98/myShop/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "home", nil); err != nil {
		app.ErrorLog.Println(err)
	}
}

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {

	if err := app.renderTemplate(w, r, "terminal", nil, "stripe-js"); err != nil {
		app.ErrorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		app.ErrorLog.Println(err)
		return
	}

	data := map[string]interface{}{
		"cardhoder": r.Form.Get("cardholder-name"),
		"email":     r.Form.Get("email"),
		"pi":        r.Form.Get("payment_intent"),
		"pm":        r.Form.Get("payment_method"),
		"pa":        r.Form.Get("payment_amount"),
		"pc":        r.Form.Get("payment_currency"),
	}

	if err := app.renderTemplate(w, r, "succeeded", &templateData{Data: data}); err != nil {
		app.ErrorLog.Println(err)
		return
	}
}

func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {

	// mockit for now
	widget := models.Widget{
		ID:             1,
		Name:           "Custom Widget",
		Description:    "A very nice widget",
		InventoryLevel: 10,
		Price:          1500,
	}
	//.--.
	data := make(map[string]any)
	data["widget"] = widget
	err := app.renderTemplate(w, r,
		"buy-once",
		&templateData{
			Data: data,
		},
		"stripe-js")

	if err != nil {
		app.ErrorLog.Println(err)
	}
}
