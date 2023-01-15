package handcash

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

/*
{
  "transactionId": "05d7df52a1c58cabada16709469e6940342cb13e8cfa3c7e1438d7ea84765787",
  "note": "Thanks dude!",
  "type": "send",
  "time": 1608226019,
  "satoshiFees": 127,
  "satoshiAmount": 5372,
  "fiatExchangeRate": 186.15198556884275,
  "fiatCurrencyCode": "USD",
  "participants": [
    {
      "type": "user",
      "alias": "mrz@moneybutton.com",
      "displayName": "MrZ",
      "profilePictureUrl": "https://www.gravatar.com/avatar/372bc0ab9b8a8930d4a86b2c5b11f11e?d=identicon",
      "responseNote": ""
    }
  ],
  "attachments": [
    {
      "value": {
        "some": "data"
      },
      "format": "json"
    }
  ],
  "appAction": "like",
  "rawTransactionHex": "01000000018598fbea559e4a59772361994f800adb63bab592e276de7ebd5805ecc639b3b8010000006a47304402200fc98489e2bbba5cb7f8cea970c0037585d42618ef60d172179307b4446854a802206be468ffd31f97c6e01a6549be50241d42633e32ba4e06ff4b2565ec897232a2412103c1fbc71737d3820890535112ac99b2471d6bacbd8a7e7825c65863a67b1d0c7effffffff03000000000000000012006a0f7b22736f6d65223a2264617461227dfc140000000000001976a914b7ce7a4c1350f1cb9dcaecca10d48f064be9197f88ac57020000000000001976a9145233794b8bdf2fd7f809b11da081189d2e79000c88ac00000000"
}
*/

// Pay makes a new payment request to the HandCash Connect API
//
// Specs: https://github.com/HandCash/handcash-connect-sdk-js/blob/master/src/api/http_request_factory.js
func (c *Client) RequestEmailCode(ctx context.Context,
	params *EmailRequestParameters) (*RequestIDBlob, error) {
	// Make sure we have payment params
	if params == nil || len(params.Email) == 0 {
		return nil, fmt.Errorf("invalid payment parameters")
	}

	// Get the signed request
	signed, err := c.getSignedRequest(
		http.MethodPost,
		endpointRequestEmail,
		"",
		params,
		currentISOTimestamp(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating signed request: %w", err)
	}

	// Convert into bytes
	var payParamsBytes []byte
	if payParamsBytes, err = json.Marshal(params); err != nil {
		return nil, err
	}

	// Make the HTTP request
	response := httpRequest(
		ctx,
		c,
		&httpPayload{
			Data:           payParamsBytes,
			ExpectedStatus: http.StatusOK,
			Method:         signed.Method,
			URL:            signed.URI,
		},
		signed,
	)

	// Error in request?
	if response.Error != nil {
		return nil, response.Error
	}

	// Unmarshal pay response
	resp := new(EmailCodeResponse)
	if err = json.Unmarshal(response.BodyContents, &resp); err != nil {
		return nil, fmt.Errorf("failed unmarshal: %w", err)
	} else if resp == nil || resp.RequestIDBlob.RequestId == "" {
		return nil, fmt.Errorf("failed to make payment")
	}
	return &resp.RequestIDBlob, nil
}

func (c *Client) VerifyCode(ctx context.Context,
	verifyCodeRequest *VerifyCodeRequest) (bool, error) {

	if val := os.Getenv("APP_ID"); val == "" {
		return false, errors.New("missing app_id secret")
	}

	// Make sure we have payment params
	if verifyCodeRequest == nil || len(verifyCodeRequest.RequestId) == 0 || len(verifyCodeRequest.VerificationCode) == 0 || len(verifyCodeRequest.PublicKey) == 0 {
		return false, fmt.Errorf("invalid verification parameters")
	}

	// Get the signed request
	signed, err := c.getSignedTrustRequest(
		http.MethodPost,
		endpointVerifyCode,
		"",
		verifyCodeRequest,
		currentISOTimestamp(),
	)
	if err != nil {
		return false, fmt.Errorf("error creating signed request: %w", err)
	}

	// Convert into bytes
	var paramsBytes []byte
	if paramsBytes, err = json.Marshal(verifyCodeRequest); err != nil {
		return false, err
	}

	// Make the HTTP request
	response := httpRequest(
		ctx,
		c,
		&httpPayload{
			Data:           paramsBytes,
			ExpectedStatus: http.StatusOK,
			Method:         signed.Method,
			URL:            signed.URI,
		},
		signed,
	)

	// Error in request?
	if response.Error != nil {
		return false, response.Error
	}

	// Unmarshal pay response
	resp := new(VerifyCodeResponse)
	if err = json.Unmarshal(response.BodyContents, &resp); err != nil {
		return false, fmt.Errorf("failed unmarshal: %w", err)
	}
	return true, nil
}

func (c *Client) CreateNewAccount(ctx context.Context, authToken string,
	params *CreateNewAccountParameters) (*CreateNewAccountResponse, error) {
	// Make sure we have payment params
	if params == nil || len(params.Email) == 0 || len(params.AccessPublicKey) == 0 {
		return nil, fmt.Errorf("invalid account parameters")
	}

	// Get the signed request
	signed, err := c.getSignedRequest(
		http.MethodPost,
		endpointAccount,
		authToken,
		params,
		currentISOTimestamp(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating signed request: %w", err)
	}

	// Convert into bytes
	var paramsBytes []byte
	if paramsBytes, err = json.Marshal(params); err != nil {
		return nil, err
	}

	// Make the HTTP request
	response := httpRequest(
		ctx,
		c,
		&httpPayload{
			Data:           paramsBytes,
			ExpectedStatus: http.StatusOK,
			Method:         signed.Method,
			URL:            signed.URI,
		},
		signed,
	)

	// Error in request?
	if response.Error != nil {
		return nil, response.Error
	}

	// Unmarshal pay response
	resp := new(CreateNewAccountResponse)
	if err = json.Unmarshal(response.BodyContents, &resp); err != nil {
		return nil, fmt.Errorf("failed unmarshal: %w", err)
	} else if resp == nil || resp.ID == "" {
		return nil, fmt.Errorf("failed to make payment")
	}
	return resp, nil
}
