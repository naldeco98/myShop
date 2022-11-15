package app

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"publishable_key": app.Config.Stripe.Key,
	}
	if err := app.renderTemplate(w, r, "terminal", &templateData{StringMap: stringMap}); err != nil {
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

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.ErrorLog.Println(err)
		return
	}
}
