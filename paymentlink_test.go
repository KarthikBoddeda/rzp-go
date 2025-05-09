// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rzp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/rzp-go"
	"github.com/stainless-sdks/rzp-go/internal/testutil"
	"github.com/stainless-sdks/rzp-go/option"
)

func TestPaymentLinkNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := rzp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.PaymentLinks.New(context.TODO(), rzp.PaymentLinkNewParams{
		Amount:         1000,
		Currency:       "INR",
		AcceptPartial:  rzp.Bool(true),
		CallbackMethod: rzp.PaymentLinkNewParamsCallbackMethodGet,
		CallbackURL:    rzp.String("https://example-callback-url.com/"),
		Customer: rzp.PaymentLinkNewParamsCustomer{
			Contact: rzp.String("+919000090000"),
			Email:   rzp.String("gaurav.kumar@example.com"),
			Name:    rzp.String("Gaurav Kumar"),
		},
		Description:           rzp.String("Payment for policy no"),
		ExpireBy:              rzp.Int(1691097057),
		FirstMinPartialAmount: rzp.Int(100),
		Notes: map[string]string{
			"policy_name": "Jeevan Bima",
		},
		Notify: rzp.PaymentLinkNewParamsNotify{
			Email: rzp.Bool(true),
			SMS:   rzp.Bool(true),
		},
		ReferenceID:    rzp.String("TS1989"),
		ReminderEnable: rzp.Bool(true),
	})
	if err != nil {
		var apierr *rzp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := rzp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *rzp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
