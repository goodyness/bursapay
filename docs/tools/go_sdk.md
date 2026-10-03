# Go SDK (`bursapay-go-sdk`)

The official Go SDK for integrating with the **BursaPay Developer Gateway API**, built with standard library `net/http` and `context.Context` support for high-throughput, concurrent backend services.

---

## Installation

```bash
go get github.com/goodyness/bursapay-go-sdk
```

Requires Go 1.18 or higher.

---

## Client Initialization

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/goodyness/bursapay-go-sdk"
)

func main() {
    // Initialize client with your BursaPay Secret Key
    client := bursapay.NewClient(os.Getenv("BURSAPAY_SECRET_KEY"))

    // Optional: override base URL (e.g., for local testing or custom proxy)
    // client.BaseURL = "http://localhost:8000/api/v1"

    ctx := context.Background()
    
    // Check health
    health, err := client.Health(ctx)
    if err != nil {
        log.Fatalf("Health check failed: %v", err)
    }
    fmt.Printf("Gateway status: %s\n", health.Status)
}
```

---

## Payments

### 1. Initialize a Payment
Generate a secure hosted checkout session:

```go
payment, err := client.Payments.Initialize(ctx, &bursapay.PaymentInitRequest{
    Amount:      15000.00,
    Email:       "customer@example.com",
    Currency:    "NGN",
    Reference:   "ORD-GO-2026-001",
    CallbackURL: "https://yourapp.com/checkout/callback",
    Metadata: map[string]interface{}{
        "order_id": "ORD-GO-2026-001",
        "user_id":  "USER-8812",
    },
})
if err != nil {
    log.Fatalf("Payment init failed: %v", err)
}

fmt.Printf("Authorization URL: %s\n", payment.AuthorizationURL)
fmt.Printf("Reference: %s\n", payment.Reference)
```

### 2. Verify a Payment
Verify transaction status after customer redirects back:

```go
verified, err := client.Payments.Verify(ctx, "ORD-GO-2026-001")
if err != nil {
    log.Fatalf("Verification failed: %v", err)
}

if verified.Status == "success" {
    fmt.Printf("Payment verified! Amount: ₦%.2f, Channel: %s\n", verified.Amount, verified.Channel)
}
```

### 3. Retrieve Payment Details
```go
payment, err := client.Payments.Retrieve(ctx, "ORD-GO-2026-001")
if err != nil {
    log.Fatalf("Retrieve failed: %v", err)
}
fmt.Printf("Payment status: %s\n", payment.Status)
```

### 4. Charge Saved Card (Tokenized Recurring)
```go
charge, err := client.Payments.ChargeSavedCard(ctx, &bursapay.ChargeSavedCardRequest{
    CustomerReference: "CUST-99214",
    AuthorizationCode: "AUTH_8b91fa02c",
    Amount:            5000.00,
})
if err != nil {
    log.Fatalf("Card charge failed: %v", err)
}
fmt.Printf("Charge reference: %s, Status: %s\n", charge.Reference, charge.Status)
```

---

## Dedicated Virtual Accounts

Generate dedicated virtual NUBAN accounts for automated bank transfer reconciliation:

```go
va, err := client.VirtualAccounts.Create(ctx, &bursapay.VirtualAccountRequest{
    CustomerEmail: "merchant@example.com",
    BVN:           "12345678901",
    PreferredBank: "Wema Bank",
})
if err != nil {
    log.Fatalf("Virtual account creation failed: %v", err)
}

fmt.Printf("Bank: %s | Account: %s | Name: %s\n", va.BankName, va.AccountNumber, va.AccountName)
```

---

## Transfers & Bulk Payouts

Disburse funds to Nigerian bank accounts:

```go
// Single Transfer
transfer, err := client.Transfers.Initiate(ctx, &bursapay.TransferRequest{
    Amount:        50000.00,
    AccountNumber: "0123456789",
    BankCode:      "058", // GTBank
    Narration:     "Vendor payout for services rendered",
    Reference:     "TRF-2026-0091",
})
if err != nil {
    log.Fatalf("Transfer failed: %v", err)
}
fmt.Printf("Transfer Ref: %s | Status: %s\n", transfer.Reference, transfer.Status)
```

---

## Subscriptions & Invoices

```go
// Create a billing plan
plan, err := client.Subscriptions.CreatePlan(ctx, &bursapay.PlanRequest{
    Name:     "Pro Monthly",
    Amount:   10000.00,
    Interval: "monthly",
    Currency: "NGN",
})

// Create an itemized invoice
invoice, err := client.Invoices.Create(ctx, &bursapay.InvoiceRequest{
    CustomerEmail: "client@example.com",
    DueDate:       "2026-10-31",
    Items: []bursapay.InvoiceItem{
        {Description: "Cloud Architecture Consulting", Quantity: 10, UnitPrice: 25000.00},
    },
})
```

---

## Wallet & Balances

```go
balance, err := client.Wallets.GetBalance(ctx)
if err != nil {
    log.Fatalf("Wallet balance check failed: %v", err)
}

fmt.Printf("Available: ₦%.2f | Ledger: ₦%.2f | Escrow: ₦%.2f\n", 
    balance.AvailableBalance, balance.LedgerBalance, balance.EscrowBalance)
```

---

## Webhook Signature Verification

Verify incoming HMAC SHA-256 signatures in your Go HTTP handler:

```go
package main

import (
    "io"
    "net/http"
    "os"

    "github.com/goodyness/bursapay-go-sdk"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    payload, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read request body", http.StatusBadRequest)
        return
    }

    signature := r.Header.Get("X-BursaPay-Signature")
    timestamp := r.Header.Get("X-BursaPay-Timestamp")
    secret := os.Getenv("BURSAPAY_WEBHOOK_SECRET")

    event, err := bursapay.ConstructEvent(payload, signature, timestamp, secret)
    if err != nil {
        http.Error(w, "Webhook verification failed: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Process typed event
    switch event.Event {
    case "payment.success":
        fmt.Printf("Payment confirmed for ref: %v\n", event.Data["reference"])
    case "virtual_account.credited":
        fmt.Printf("Virtual account credited: %v\n", event.Data["amount"])
    case "dispute.created":
        fmt.Printf("Dispute logged: %v\n", event.Data["reference"])
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"success"}`))
}
```

---

## Error Handling

The Go SDK returns typed error structures:

```go
payment, err := client.Payments.Initialize(ctx, req)
if err != nil {
    if apiErr, ok := err.(*bursapay.APIError); ok {
        fmt.Printf("API Error [%d]: %s (Code: %s)\n", apiErr.StatusCode, apiErr.Message, apiErr.Code)
    } else {
        fmt.Printf("Network error: %v\n", err)
    }
}
```
