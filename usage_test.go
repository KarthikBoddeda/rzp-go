// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rzp_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/rzp-go"
	"github.com/stainless-sdks/rzp-go/internal/testutil"
	"github.com/stainless-sdks/rzp-go/option"
)

func TestUsage(t *testing.T) {
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
	paymentLink, err := client.PaymentLinks.New(context.TODO(), rzp.PaymentLinkNewParams{
		Amount:   1000,
		Currency: "INR",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", paymentLink.ID)
}
