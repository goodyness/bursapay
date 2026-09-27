# Wallet & Settlements API

The Wallet & Settlements API provides real-time access to merchant available balances, ledger audit records, and automated bank payout schedules.

---

## 1. Get Wallet Balance

Retrieve real-time available funds and pending clearing balances.

- **Endpoint:** `GET /api/v1/wallet/balance`
- **Required Scope:** `transfers:read`

### Example Response (`200 OK`)

```json
{
  "success": true,
  "data": {
    "currency": "NGN",
    "available_balance": "1450200.50",
    "pending_balance": "85000.00",
    "escrow_balance": "320000.00",
    "total_withdrawn": "8900000.00",
    "last_settlement_at": "2026-09-26T23:00:00Z"
  }
}
```

---

## 2. Wallet Ledger Audit Trail

Retrieve an immutable double-entry journal of every credit, debit, fee deduction, and payout affecting your account.

- **Endpoint:** `GET /api/v1/wallet/ledger/`
- **Required Scope:** `transfers:read`

### Query Parameters

| Parameter | Type | Description |
|---|---|---|
| `from_date` | `string` | Filter entries from date (`YYYY-MM-DD`). |
| `to_date` | `string` | Filter entries to date (`YYYY-MM-DD`). |
| `type` | `string` | Filter: `credit`, `debit`, `fee`, `settlement`, `refund`. |

---

## 3. Settlements & Automated Payouts

View completed and scheduled automated batch settlements to your corporate bank account.

- **Endpoint:** `GET /api/v1/wallet/settlements/`
- **Required Scope:** `transfers:read`

### Example Settlement Object:
```json
{
  "reference": "SETTLE-20260926-01",
  "total_gross": "5400000.00",
  "total_fees": "81000.00",
  "net_payout": "5319000.00",
  "destination_bank": "Zenith Bank PLC",
  "destination_account": "1012345678",
  "status": "paid",
  "payout_time": "2026-09-26T23:00:00Z"
}
```
