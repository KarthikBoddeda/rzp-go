// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rzp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/KarthikBoddeda/rzp-go/internal/apijson"
	"github.com/KarthikBoddeda/rzp-go/internal/requestconfig"
	"github.com/KarthikBoddeda/rzp-go/option"
	"github.com/KarthikBoddeda/rzp-go/packages/param"
	"github.com/KarthikBoddeda/rzp-go/packages/respjson"
)

// PaymentLinkService contains methods and other services that help with
// interacting with the rzp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentLinkService] method instead.
type PaymentLinkService struct {
	Options []option.RequestOption
}

// NewPaymentLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentLinkService(opts ...option.RequestOption) (r PaymentLinkService) {
	r = PaymentLinkService{}
	r.Options = opts
	return
}

// Create a new payment link.
// [Docs](https://razorpay.com/docs/api/payments/payment-links/create-standard/)
func (r *PaymentLinkService) New(ctx context.Context, body PaymentLinkNewParams, opts ...option.RequestOption) (res *PaymentLink, err error) {
	opts = append(r.Options[:], opts...)
	path := "payment_links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a payment link by its unique ID.
// [Docs](https://razorpay.com/docs/api/payments/payment-links/fetch-id-standard/)
func (r *PaymentLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentLink, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payment_links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PaymentLink struct {
	ID                    string              `json:"id"`
	AcceptPartial         bool                `json:"accept_partial"`
	Amount                int64               `json:"amount"`
	AmountPaid            int64               `json:"amount_paid"`
	CallbackMethod        string              `json:"callback_method"`
	CallbackURL           string              `json:"callback_url"`
	CancelledAt           int64               `json:"cancelled_at"`
	CreatedAt             int64               `json:"created_at"`
	Currency              string              `json:"currency"`
	Customer              PaymentLinkCustomer `json:"customer"`
	Description           string              `json:"description"`
	Entity                string              `json:"entity"`
	ExpireBy              int64               `json:"expire_by"`
	ExpiredAt             int64               `json:"expired_at"`
	FirstMinPartialAmount int64               `json:"first_min_partial_amount"`
	Notes                 map[string]string   `json:"notes"`
	Notify                PaymentLinkNotify   `json:"notify"`
	Payments              []any               `json:"payments,nullable"`
	ReferenceID           string              `json:"reference_id"`
	ReminderEnable        bool                `json:"reminder_enable"`
	Reminders             []any               `json:"reminders"`
	ShortURL              string              `json:"short_url"`
	Status                string              `json:"status"`
	UpdatedAt             int64               `json:"updated_at"`
	UserID                string              `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AcceptPartial         respjson.Field
		Amount                respjson.Field
		AmountPaid            respjson.Field
		CallbackMethod        respjson.Field
		CallbackURL           respjson.Field
		CancelledAt           respjson.Field
		CreatedAt             respjson.Field
		Currency              respjson.Field
		Customer              respjson.Field
		Description           respjson.Field
		Entity                respjson.Field
		ExpireBy              respjson.Field
		ExpiredAt             respjson.Field
		FirstMinPartialAmount respjson.Field
		Notes                 respjson.Field
		Notify                respjson.Field
		Payments              respjson.Field
		ReferenceID           respjson.Field
		ReminderEnable        respjson.Field
		Reminders             respjson.Field
		ShortURL              respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		UserID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLink) RawJSON() string { return r.JSON.raw }
func (r *PaymentLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkCustomer struct {
	Contact string `json:"contact"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contact     respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkCustomer) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkNotify struct {
	Email bool `json:"email"`
	SMS   bool `json:"sms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkNotify) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkNotify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkNewParams struct {
	// Amount in the smallest currency unit (e.g., paise)
	Amount int64 `json:"amount,required"`
	// ISO currency code
	Currency string `json:"currency,required"`
	// Allow partial payments
	AcceptPartial param.Opt[bool] `json:"accept_partial,omitzero"`
	// Redirect URL after payment
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// Description of the payment link
	Description param.Opt[string] `json:"description,omitzero"`
	// Unix timestamp when the link expires
	ExpireBy param.Opt[int64] `json:"expire_by,omitzero"`
	// Minimum amount for the first partial payment
	FirstMinPartialAmount param.Opt[int64] `json:"first_min_partial_amount,omitzero"`
	// Unique reference for the payment link
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	// Enable reminders for the payment link
	ReminderEnable param.Opt[bool] `json:"reminder_enable,omitzero"`
	// HTTP method for callback
	//
	// Any of "get", "post".
	CallbackMethod PaymentLinkNewParamsCallbackMethod `json:"callback_method,omitzero"`
	Customer       PaymentLinkNewParamsCustomer       `json:"customer,omitzero"`
	Notes          map[string]string                  `json:"notes,omitzero"`
	Notify         PaymentLinkNewParamsNotify         `json:"notify,omitzero"`
	paramObj
}

func (r PaymentLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP method for callback
type PaymentLinkNewParamsCallbackMethod string

const (
	PaymentLinkNewParamsCallbackMethodGet  PaymentLinkNewParamsCallbackMethod = "get"
	PaymentLinkNewParamsCallbackMethodPost PaymentLinkNewParamsCallbackMethod = "post"
)

type PaymentLinkNewParamsCustomer struct {
	Contact param.Opt[string] `json:"contact,omitzero"`
	Email   param.Opt[string] `json:"email,omitzero"`
	Name    param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PaymentLinkNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParamsCustomer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParamsCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkNewParamsNotify struct {
	Email param.Opt[bool] `json:"email,omitzero"`
	SMS   param.Opt[bool] `json:"sms,omitzero"`
	paramObj
}

func (r PaymentLinkNewParamsNotify) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParamsNotify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParamsNotify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
