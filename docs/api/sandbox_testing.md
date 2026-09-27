# Sandbox Testing & Data Simulator

The BursaPay Sandbox is a fully isolated staging environment that lets developers test end-to-end payment lifecycles, virtual account deposits, recurring subscription charges, webhook deliveries, and edge cases without touching real money.

---

## 💳 Test Card Numbers

Use the following test card numbers in sandbox mode (`sk_test_...` or `pk_test_...`):

| Test Scenario | Card Number | Expiry (MM/YY) | CVV | Expected Outcome |
|---|---|---|---|---|
| **Successful Charge** | `4084 0841 0841 0841` | Any future date | `408` | `status: "success"` (`payment.success` webhook) |
| **Insufficient Funds**| `4084 0841 0841 0842` | Any future date | `408` | `error_code: "INSUFFICIENT_FUNDS"` |
| **3D Secure (OTP)** | `4084 0841 0841 0843` | Any future date | `408` | Triggers OTP modal (Test PIN: `123456`) |
| **Expired Card** | `4084 0841 0841 0844` | Past date (e.g. `01/20`) | `408` | `error_code: "CARD_EXPIRED"` |
| **Declined / Fraud** | `4084 0841 0841 0845` | Any future date | `408` | `error_code: "TRANSACTION_DECLINED"` |

---

## 🏦 Simulated Virtual Accounts & Bank Transfers

In sandbox mode:
- Inbound transfers to any test virtual account (`0123456789`) simulate instant settlement.
- Webhooks (`virtual_account.credited`) dispatch within **2 seconds** of request creation.

---

## 🧪 Synthetic Data Seeding API

Developers can instantly populate their sandbox account with realistic test customers, transaction histories, disputes, and subscription plans using the **Sandbox Seeding API**.

### 1. Seed Test Data

- **Endpoint:** `POST /api/v1/sandbox/seed/`
- **Environment:** Sandbox / Test Mode Only
- **Authentication:** Bearer `bp_sec_test_DEMO_KEY_HERE`

#### Supported Scenarios:
| Scenario Code | Generated Test Records |
|---|---|
| `basic` | 10 Payments (Success/Failed), 5 Customers. |
| `high_volume` | 100 Transactions across Visa/Mastercard/Verve, 20 Customers. |
| `disputes` | 5 Payments, 5 Customers, 2 Open Chargeback Disputes. |
| `subscriptions` | 3 Billing Plans (Monthly/Weekly/Annual), 10 Active Subscriptions. |

#### Request Example:

```bash
curl -X POST https://api.bursapay.com/api/v1/sandbox/seed/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "scenario": "high_volume"
  }'
```

#### Response Example (`201 Created`):

```json
{
  "success": true,
  "message": "Seeded scenario 'high_volume' successfully",
  "data": {
    "scenario": "high_volume",
    "seeded_counts": {
      "payments": 100,
      "customers": 20
    },
    "seed_token": "a1b2c3d4"
  }
}
```

---

### 2. Reset / Wipe Sandbox Data

Remove all synthetic seeded records and return your test account to a clean slate:

- **Endpoint:** `DELETE /api/v1/sandbox/seed/`
- **Authentication:** Bearer `bp_sec_test_DEMO_KEY_HERE`

```bash
curl -X DELETE https://api.bursapay.com/api/v1/sandbox/seed/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE"
```

```json
{
  "success": true,
  "message": "All seeded sandbox records have been deleted"
}
```
