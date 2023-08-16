package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/cyan-store/hook/cache"
	"github.com/cyan-store/hook/config"
	"github.com/cyan-store/hook/database"
	"github.com/cyan-store/hook/log"
	"github.com/cyan-store/hook/mail"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"github.com/stripe/stripe-go/v74/webhook"
)

func convertPaymentStatus(status stripe.CheckoutSessionPaymentStatus) (string, error) {
	switch status {
	case "paid":
		return "PAID", nil
	case "unpaid":
		return "UNPAID", nil
	case "no_payment_required":
		return "CANCELED", nil
	default:
		return "", errors.New("invalid payment status")
	}
}

func updateOrder(sessionID string) int {
	sc := client.New(config.Data.StripeKey, nil)
	sessions, err := sc.CheckoutSessions.Get(sessionID, nil)

	// Fetch stripe session
	if err != nil {
		log.Error.Println("[checkout] Could not get stripe session:", err)
		mail.ReportError(sessionID, "Could not get stripe session", err)

		return http.StatusInternalServerError
	}

	if sessions.Status == "expired" {
		log.Info.Println(sessions.ID, "has expired")
		return http.StatusOK
	} else if sessions.Status == "open" {
		log.Info.Println("Skipping", sessions.ID, "- session is open.")
		return http.StatusOK
	}

	// Fetch redis session
	cs, err := cache.GetSession(sessionID)
	id := uuid.New()

	if err != nil {
		log.Error.Println("[checkout] Could not get redis session:", err)
		mail.ReportError(sessionID, "Could not get redis session", err)

		return http.StatusInternalServerError
	}

	// Create order, clear cache
	cstatus, err := convertPaymentStatus(sessions.PaymentStatus)

	if err != nil {
		log.Error.Println("[checkout] Could not get payment status!")
		mail.ReportError(sessionID, "Could not get payment status", err)

		return http.StatusInternalServerError
	}

	if err := database.CreateOrder(
		id.String(), cs.IDS, sessions.PaymentIntent.ID,
		cs.User, cstatus, cs.Quantity,
		int(sessions.AmountTotal), cs.Email, sessions.CustomerDetails.Address.Country,
		sessions.CustomerDetails.Address.PostalCode, "PENDING",
	); err != nil {
		log.Error.Println("[checkout] Could not insert order:", err)
		mail.ReportError(sessionID, "Could not insert order", err)

		return http.StatusInternalServerError
	}

	if err := cache.DeleteSession(sessionID); err != nil {
		log.Error.Println("[checkout] Could not clear cache:", err, sessionID)
		mail.ReportError(sessionID, "Could not clear cache", err)

		return http.StatusInternalServerError
	}

	log.Info.Println("[checkout] Resolved", sessions.PaymentIntent.ID)
	mail.ReportSuccess(sessions.PaymentIntent.ID)

	return http.StatusOK
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	// Read body
	if err != nil {
		log.Error.Println("[checkout] Could not read request body:", err)
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), config.Data.StripeHook)

	// Verify webhook
	if err != nil {
		log.Error.Println("[checkout] Could not verify webhook:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Handle completed session
	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			log.Error.Println("[checkout] Could not parsing webhook JSON:", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.WriteHeader(updateOrder(session.ID))
		return
	}

	w.WriteHeader(http.StatusOK)
}
