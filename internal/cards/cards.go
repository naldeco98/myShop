package cards

import (
	"log"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

// alias for CreatePaymentIntent
func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create payment CreatePaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	pI, err := paymentintent.New(params)
	if err != nil {
		var msg string
		if stripeErr, ok := err.(*stripe.Error); ok {
			log.Default().Println(err.Error())
			msg = cardErrorMsg(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pI, "", nil
}

func cardErrorMsg(code stripe.ErrorCode) string {
	var msg string
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was decliend"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect Zip Code"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Incorrect Postal Code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is to large to charge to your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is to small to charge to your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient balance"
	default:
		msg = "Your card was decliend"
	}
	return msg
}
