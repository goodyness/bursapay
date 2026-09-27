# Payments API Reference

The Payments API is the core interface for accepting card payments, bank transfers, USSD, and Apple Pay on BursaPay.

---

## 1. Initialize Payment

Generate a secure checkout session URL or access code for client-side popup checkout.

- **Endpoint:** `POST /api/v1/payments/initialize/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `amount` | `number` / `string` | **Yes** | Amount in standard currency units (e.g. `10000.00` for ₦10,000.00). |
| `email` | `string` | **Yes** | Payer's email address for receipt delivery. |
| `currency` | `string` | No | Currency code (`NGN`, `USD`, `GHS`, `KES`). Default: `NGN`. |
| `reference` | `string` | No | Unique merchant transaction ID. Auto-generated if omitted. |
| `callback_url`| `string` | No | URL to redirect customer upon payment completion. |
| `channels` | `array` | No | Supported payment channels: `["card", "bank_transfer", "ussd", "qr"]`. |
| `metadata` | `object` | No | Key-value dictionary stored with the transaction. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/payments/initialize/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 15000.00,
    "email": "sarah.ade@example.com",
    "currency": "NGN",
    "reference": "BP-ORD-90214",
    "callback_url": "https://mystore.com/checkout/success",
    "metadata": {
      "cart_id": "CART-8812",
      "customer_tier": "VIP"
    }
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "message": "Payment initialized",
  "data": {
    "reference": "BP-ORD-90214",
    "access_code": "0pylnd82pe",
    "authorization_url": "https://checkout.bursapay.com/pay/0pylnd82pe",
    "amount": "15000.00",
    "currency": "NGN"
  }
}
```

---

## 2. Verify Payment

Check the real-time status of a transaction using its reference.

- **Endpoint:** `GET /api/v1/payments/verify/`
- **Required Scope:** `payments:read`

### Query Parameters

| Parameter | Type | Required | Description |
|---|---|---|---|
| `reference` | `string` | **Yes** | The unique reference of the payment to verify. |

### Example Request

```bash
curl -X GET "https://api.bursapay.com/api/v1/payments/verify/?reference=BP-ORD-90214" \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE"
```

### Example Response (`200 OK`)

```json
{
  "success": true,
  "message": "Payment verified successfully",
  "data": {
    "reference": "BP-ORD-90214",
    "status": "success",
    "amount": "15000.00",
    "currency": "NGN",
    "channel": "card",
    "gateway_fee": "225.00",
    "net_amount": "14775.00",
    "paid_at": "2026-09-27T02:14:00Z",
    "customer": {
      "email": "sarah.ade@example.com",
      "customer_code": "CUST-99214"
    },
    "authorization": {
      "authorization_code": "AUTH_8b91fa02c",
      "card_type": "visa",
      "last4": "4081",
      "exp_month": "12",
      "exp_year": "2028",
      "bank": "Guaranty Trust Bank",
      "reusable": true
    },
    "metadata": {
      "cart_id": "CART-8812"
    }
  }
}
```

---

## 3. Charge Saved Card (Tokenized Recurring Charge)

Charge a customer's previously tokenized card authorization without requiring interactive checkout.

- **Endpoint:** `POST /api/v1/payments/charge-saved-card/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `authorization_code`| `string` | **Yes** | Reusable card token obtained from previous payment verification. |
| `email` | `string` | **Yes** | Customer email. |
| `amount` | `number` | **Yes** | Amount to charge. |
| `reference` | `string` | No | Custom transaction reference. |

---

## 4. Payment Intents (2-Step Authorize & Capture)

For marketplaces, hotel booking, and ride-hailing where funds must be held and captured only after service delivery.

- **Create Intent:** `POST /api/v1/payment-intents/`
- **Capture Intent:** `POST /api/v1/payment-intents/<intent_reference>/capture/`
- **Cancel Intent:** `POST /api/v1/payment-intents/<intent_reference>/cancel/`

### Capture Example

```bash
curl -X POST https://api.bursapay.com/api/v1/payment-intents/PI-991204/capture/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "amount_to_capture": 12500.00
  }'
```
