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
| `charge_at` | `string` | No | ISO 8601 timestamp for scheduled future charge. |
| `splits` | `array` | No | Multi-party split payment rules with subaccounts and percentages. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/payments/initialize/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: 7b31e9a2-4a5f-4a3d-a4e9-9d0a1b2c3d4e" \
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

## 3. List & Retrieve Payments

### List Payments
- **Endpoint:** `GET /api/v1/payments/`
- **Required Scope:** `payments:read`
- **Query Params:** `status` (`success`, `failed`, `pending`), `page`, `page_size`, `from_date`, `to_date`.

### Retrieve Payment Detail
- **Endpoint:** `GET /api/v1/payments/<reference>/`
- **Required Scope:** `payments:read`

---

## 4. Charge Saved Card (Tokenized 1-Click Recurring Charge)

Charge a customer's previously saved card authorization without requiring an interactive checkout redirect.

- **Endpoint:** `POST /api/v1/payments/charge-saved-card/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `customer_reference`| `string` | **Yes** | BursaPay customer reference (`CUST-xxx`). |
| `authorization_code`| `string` | **Yes** | Tokenized card authorization (`AUTH_xxx`). |
| `amount` | `number` | **Yes** | Amount to charge in NGN. |
| `reference` | `string` | No | Custom merchant reference. |

---

## 5. Direct Authorization Charge

- **Endpoint:** `POST /api/v1/payments/charge/`
- **Required Scope:** `payments:write`

Charge an authorization code directly with email and amount.

---

## 6. Payment Intents (2-Step Authorize & Delayed Capture)

For marketplaces, car rentals, and escrow workflows where funds must be authorized, held, and captured later.

| Action | Method | Endpoint | Description |
|---|---|---|---|
| **Create Intent** | `POST` | `/api/v1/payment-intents/` | Create authorization hold intent. |
| **List Intents** | `GET` | `/api/v1/payment-intents/` | List merchant payment intents. |
| **Retrieve Intent**| `GET` | `/api/v1/payment-intents/<intent_ref>/` | Retrieve hold status and details. |
| **Capture Intent** | `POST` | `/api/v1/payment-intents/<intent_ref>/capture/` | Capture held funds to wallet. |
| **Cancel Intent**  | `POST` | `/api/v1/payment-intents/<intent_ref>/cancel/`  | Release/void held funds. |

### Capture Example Request
```bash
curl -X POST https://api.bursapay.com/api/v1/payment-intents/PI-991204/capture/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "amount_to_capture": 12500.00
  }'
```

---

## 7. Bulk Payment Initialization

Initialize up to 100 payments simultaneously in a single API call.

- **Endpoint:** `POST /api/v1/payments/bulk/`
- **Required Scope:** `payments:write`
- **Status Endpoint:** `GET /api/v1/payments/bulk/<batch_reference>/`

### Request Body
```json
{
  "payments": [
    {
      "amount": 5000.00,
      "email": "student1@unilag.edu.ng",
      "reference": "DUES-2026-001"
    },
    {
      "amount": 5000.00,
      "email": "student2@unilag.edu.ng",
      "reference": "DUES-2026-002"
    }
  ]
}
```

---

## 8. Cancel Scheduled Payment & Mark Order Fulfillment

### Cancel Scheduled Charge
- **Endpoint:** `DELETE /api/v1/payments/<reference>/schedule/`
- **Required Scope:** `payments:write`

### Mark Fulfillment
- **Endpoint:** `POST /api/v1/payments/<reference>/fulfillment/`
- **Required Scope:** `payments:write`
- **Payload:** `{"fulfillment_status": "fulfilled", "tracking_number": "DHL-981203"}`
