package api

import (
	"testing"
)

func TestGetSignedRequest(t *testing.T) {

	// Test known values for a login
	t.Run("valid auth", func(t *testing.T) {
		// These values are used
		authToken := "68d8fadc95324afa853f00923e0b86f06a76ceb7a6afbb1784e0dde8f43989a0"
		method := "GET"

		// To generate these values
		pubKey := "0275e7081e5b6e73c94998098e075c0ed888d1eb33c721ee38ee741648b108c90d"
		// privateKey := "5JcTmjpJkfkcRnf3W2qTvauC4mczNsnUY3SLm6EcKQDS3Gj2wGh" // 68d8fadc95324afa853f00923e0b86f06a76ceb7a6afbb1784e0dde8f439
		timestamp := "2020-11-22T16:31:23.304Z"
		signature := "304402205b08177d369e7cc112fc4de651f10fcc42cb1d65309a6851156b04e9abd38c0002206ae51434825bf24276d759fb391700db883dbfdee0fa68256d063c7413bdc133"
		endpoint := "/v1/connect/profile/currentUserProfile"
		// body := nil
		signedRequest, err := GetSignedRequest(method, endpoint, authToken, nil, timestamp)
		if err != nil {
			t.Errorf("couldnt get signed request %s", err)
		}

		if signedRequest.Headers.OauthPublicKey != pubKey {
			t.Errorf("Pubkey mismatch: %s : %s", signedRequest.Headers.OauthPublicKey, pubKey)
		}

		if signedRequest.Headers.OauthSignature != signature {
			t.Errorf("Signature mismatch: %s : %s", signedRequest.Headers.OauthSignature, signature)
		}
	})
}