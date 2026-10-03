package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type PaymentInitRequest struct {
	Amount      float64 `json:"amount"`
	Email       string  `json:"email"`
	Currency    string  `json:"currency"`
	CallbackURL string  `json:"callback_url"`
}

type PaymentInitData struct {
	AuthorizationURL string `json:"authorization_url"`
	Reference        string `json:"reference"`
}

type TransferItem struct {
	Amount        float64 `json:"amount"`
	RecipientBank string  `json:"recipient_bank"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	Narration     string  `json:"narration,omitempty"`
}

type BulkTransferRequest struct {
	Transfers []TransferItem `json:"transfers"`
}

type apiEnvelope struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func doRequest(method, url string, body interface{}, secretKey string) ([]byte, int, error) {
	var reqBody io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(encoded)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+secretKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("HTTP request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("ERROR: %s %s returned %d: %s", method, url, resp.StatusCode, string(respBody))
	}

	return respBody, resp.StatusCode, nil
}

func verifyWebhookSignature(body []byte, sigHeader, tsHeader, secret string) bool {
	if secret == "" || sigHeader == "" {
		return false
	}

	var ts, expectedHex string
	if strings.HasPrefix(sigHeader, "t=") {
		parts := make(map[string]string)
		for _, seg := range strings.Split(sigHeader, ",") {
			idx := strings.Index(seg, "=")
			if idx > 0 {
				parts[seg[:idx]] = seg[idx+1:]
			}
		}
		ts = parts["t"]
		expectedHex = parts["v1"]
	} else {
		ts = tsHeader
		expectedHex = sigHeader
	}

	if ts == "" || expectedHex == "" {
		return false
	}

	var payload interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		return false
	}
	compactBody, err := json.Marshal(payload)
	if err != nil {
		return false
	}

	message := ts + "." + string(compactBody)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	computed := hex.EncodeToString(mac.Sum(nil))

	expectedBytes, err := hex.DecodeString(expectedHex)
	if err != nil {
		return false
	}
	computedBytes, _ := hex.DecodeString(computed)
	return hmac.Equal(computedBytes, expectedBytes)
}

func makeWebhookHandler(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		sigHeader := r.Header.Get("X-BursaPay-Signature")
		tsHeader := r.Header.Get("X-BursaPay-Timestamp")

		if !verifyWebhookSignature(body, sigHeader, tsHeader, secret) {
			log.Printf("webhook: signature verification failed — sig=%s ts=%s", sigHeader, tsHeader)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var payload map[string]json.RawMessage
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		var eventName string
		if raw, ok := payload["event"]; ok {
			_ = json.Unmarshal(raw, &eventName)
		}

		var data map[string]interface{}
		if raw, ok := payload["data"]; ok {
			_ = json.Unmarshal(raw, &data)
		}

		switch eventName {
		case "transfer.success":
			ref, _ := data["transfer_reference"].(string)
			log.Printf("EVENT transfer.success — reference=%s", ref)
		case "payment.success":
			ref, _ := data["reference"].(string)
			log.Printf("EVENT payment.success — reference=%s", ref)
		default:
			log.Printf("EVENT %s received", eventName)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}
}

func main() {
	secretKey := os.Getenv("BURSAPAY_SECRET_KEY")
	baseURL := os.Getenv("BURSAPAY_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.bursapay.com/api/v1"
	}

	if secretKey == "" {
		fmt.Fprintf(os.Stderr, "ERROR: BURSAPAY_SECRET_KEY is required\n")
		os.Exit(1)
	}

	reference := flag.String("reference", "", "Payment reference to disburse funds for")
	flag.Parse()

	// 1. Initialize checkout payment
	initBody := PaymentInitRequest{
		Amount:      50000.00,
		Email:       "client@example.com",
		Currency:    "NGN",
		CallbackURL: "http://localhost:8080/callback",
	}

	respBody, statusCode, err := doRequest("POST", baseURL+"/payments/initialize/", initBody, secretKey)
	if err != nil {
		log.Fatalf("Payment initialization failed: %v", err)
	}
	if statusCode >= 200 && statusCode < 300 {
		var initEnv apiEnvelope
		_ = json.Unmarshal(respBody, &initEnv)
		var initData PaymentInitData
		_ = json.Unmarshal(initEnv.Data, &initData)
		fmt.Println("✓ Payment initialized")
		fmt.Printf("  Authorization URL: %s\n", initData.AuthorizationURL)
		fmt.Printf("  Reference:         %s\n\n", initData.Reference)
	}

	// 2. Perform bulk transfer if reference given
	if *reference != "" {
		fmt.Printf("Submitting bulk vendor payout for reference: %s\n", *reference)
		bulkBody := BulkTransferRequest{
			Transfers: []TransferItem{
				{
					Amount:        20000.00,
					RecipientBank: "058",
					AccountNumber: "0123456789",
					AccountName:   "Vendor A",
					Narration:     fmt.Sprintf("Payout for ref %s", *reference),
				},
				{
					Amount:        25000.00,
					RecipientBank: "035",
					AccountNumber: "9876543210",
					AccountName:   "Vendor B",
					Narration:     fmt.Sprintf("Payout for ref %s", *reference),
				},
			},
		}

		bulkResp, _, err := doRequest("POST", baseURL+"/transfers/bulk/", bulkBody, secretKey)
		if err == nil {
			fmt.Printf("✓ Bulk transfer submitted: %s\n", string(bulkResp))
		}
	}

	webhookSecret := os.Getenv("BURSAPAY_WEBHOOK_SECRET")
	if webhookSecret == "" {
		webhookSecret = secretKey
	}

	http.HandleFunc("/webhooks/bursapay/", makeWebhookHandler(webhookSecret))
	fmt.Println("Webhook receiver running on http://localhost:8080/webhooks/bursapay/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
