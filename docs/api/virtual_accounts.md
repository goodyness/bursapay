# Dedicated Virtual Accounts (NUBAN) API

The Virtual Accounts API allows software platforms to provision dynamic or permanent Nigerian Uniform Bank Account Numbers (NUBAN) for their customers. When funds are transferred to this account via mobile banking or USSD, BursaPay automatically detects the deposit and dispatches a webhook event.

---

## 1. Create a Dedicated Virtual Account

Provision a permanent bank account assigned to a specific customer.

- **Endpoint:** `POST /api/v1/virtual-accounts/`
- **Required Scope:** `virtual_accounts:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `customer_reference` | `string` | **Yes** | BursaPay Customer Reference (e.g., `CUST-99214`). |
| `preferred_bank` | `string` | No | Preferred partner bank: `wema-bank`, `titan-paystack`, `access-bank`. |
| `account_name` | `string` | No | Custom prefix for the account display name. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/virtual-accounts/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_reference": "CUST-99214",
    "preferred_bank": "wema-bank",
    "account_name": "ACME - Sarah Ade"
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "message": "Virtual account generated successfully",
  "data": {
    "account_number": "0123456789",
    "account_name": "ACME / Sarah Ade",
    "bank_name": "Wema Bank",
    "bank_code": "035",
    "currency": "NGN",
    "status": "active",
    "customer_reference": "CUST-99214",
    "created_at": "2026-09-27T02:15:00Z"
  }
}
```

---

## 2. Inbound Deposit Notification Lifecycle

When a payer makes a bank transfer to the virtual account number:

```
Customer Bank App ──── Transfer ───► Central NIP Switch ────► BursaPay Engine
                                                                   │
                                                                   ▼
Merchant Webhook Server ◄──── `virtual_account.credited` ──────────┘
```

### Webhook Payload Example:

```json
{
  "event": "virtual_account.credited",
  "data": {
    "account_number": "0123456789",
    "bank_name": "Wema Bank",
    "amount": "25000.00",
    "currency": "NGN",
    "fee": "150.00",
    "net_amount": "24850.00",
    "payer_name": "JOHN DOE",
    "payer_bank": "Zenith Bank",
    "session_id": "999001240927021500123456789012",
    "customer_reference": "CUST-99214",
    "transaction_reference": "VA-DEP-883104",
    "settled_at": "2026-09-27T02:15:22Z"
  }
}
```

---

## 3. Retrieve Virtual Account Details

- **Endpoint:** `GET /api/v1/virtual-accounts/<id>/`
- **Required Scope:** `payments:read`
