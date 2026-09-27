# Transfers & Payouts API

The Transfers API enables businesses, universities, and platforms to disburse funds to any commercial bank account, microfinance bank, or mobile money operator across Nigeria.

---

## 1. Initiate Single Transfer

- **Endpoint:** `POST /api/v1/transfers/`
- **Required Scope:** `transfers:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `amount` | `number` | **Yes** | Transfer amount in NGN. |
| `bank_code` | `string` | **Yes** | 3-digit CBN bank code (e.g., `058` for GTBank, `011` for First Bank). |
| `account_number` | `string` | **Yes** | 10-digit NUBAN account number. |
| `account_name` | `string` | **Yes** | Verified recipient account name. |
| `narration` | `string` | No | Transfer remark / description. |
| `reference` | `string` | No | Custom merchant reference. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/transfers/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: b7189d20-1a2b-4c3d-8e4f-998877665544" \
  -d '{
    "amount": 75000.00,
    "bank_code": "058",
    "account_number": "0123456789",
    "account_name": "TUNDE BAKARE",
    "narration": "Departmental Equipment Honorarium"
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "message": "Transfer queued for processing",
  "data": {
    "reference": "TRF-2026-90214",
    "status": "processing",
    "amount": "75000.00",
    "fee": "25.00",
    "bank_name": "Guaranty Trust Bank",
    "account_number": "0123456789",
    "recipient_name": "TUNDE BAKARE",
    "created_at": "2026-09-27T02:16:00Z"
  }
}
```

---

## 2. Bulk Transfers / Payroll Payouts

Disburse payments to hundreds of recipients simultaneously in a single API call (ideal for salary disbursements, vendor settlements, and refund batches).

- **Endpoint:** `POST /api/v1/transfers/bulk/`
- **Required Scope:** `transfers:write`

### Request Body Example

```json
{
  "currency": "NGN",
  "transfers": [
    {
      "amount": 120000.00,
      "bank_code": "058",
      "account_number": "0123456789",
      "account_name": "JOHN DOE",
      "reference": "SAL-001"
    },
    {
      "amount": 95000.00,
      "bank_code": "033",
      "account_number": "9876543210",
      "account_name": "JANE SMITH",
      "reference": "SAL-002"
    }
  ]
}
```

---

## 3. Retrieve Transfer Status

- **Endpoint:** `GET /api/v1/transfers/<reference>/`
- **Required Scope:** `transfers:read`
