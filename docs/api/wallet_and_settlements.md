# Wallet, Settlements, Reconciliation & Analytics API

The Wallet & Ledger API provides financial transparency, programmatic settlement history, 3-way reconciliation audit export, and operational metrics.

---

## 1. Query Wallet Balance

Check real-time available, pending, and escrow balances.

- **Endpoint:** `GET /api/v1/wallet/balance`
- **Required Scope:** `transfers:read`

### Response Example:
```json
{
  "success": true,
  "data": {
    "currency": "NGN",
    "available_balance": "482950.00",
    "ledger_balance": "512450.00",
    "escrow_balance": "29500.00",
    "pending_payout": "0.00",
    "updated_at": "2026-09-27T03:10:00Z"
  }
}
```

---

## 2. Wallet Ledger (Audit Trail)

Retrieve the double-entry accounting journal entries for the merchant wallet.

- **Endpoint:** `GET /api/v1/wallet/ledger/`
- **Required Scope:** `transfers:read`
- **Query Params:** `entry_type` (`credit`, `debit`), `page`, `page_size`, `from_date`, `to_date`

---

## 3. Settlements

### List Automated Bank Payouts
- **Endpoint:** `GET /api/v1/wallet/settlements/`
- **Required Scope:** `transfers:read`

### Retrieve Settlement Batch Details
- **Endpoint:** `GET /api/v1/wallet/settlements/<reference>/`
- **Required Scope:** `transfers:read`

---

## 4. Reconciliation Engine

Automated 3-way reconciliation comparing Gateway transactions, internal ledger records, and bank statements.

### Reconciliation Summary
- **Endpoint:** `GET /api/v1/reconciliation/`
- **Required Scope:** `transfers:read`
- **Query Params:** `start_date`, `end_date`

### Reconciliation Export (CSV)
- **Endpoint:** `GET /api/v1/reconciliation/export/`
- **Required Scope:** `transfers:read`
- **Response:** CSV download with columns: `Transaction Date, Reference, Channel, Gross Amount, Fee, Net Amount, Bank Match Status, Ledger Journal ID`.

---

## 5. Analytics Endpoint

Query high-level operational performance metrics:

- **Endpoint:** `GET /api/v1/analytics/`
- **Required Scope:** `payments:read`
- **Query Params:** `period` (`7d`, `30d`, `90d`, `1y`, `custom`)

### Response Example:
```json
{
  "success": true,
  "data": {
    "total_volume": "14250000.00",
    "total_transactions": 842,
    "successful_transactions": 819,
    "failed_transactions": 23,
    "success_rate": "97.27%",
    "average_order_value": "17399.27",
    "top_channel": "card"
  }
}
```

---

## 6. IP Check & Allowlist Status

Verify caller's IP address and whether it satisfies the merchant's configured CIDR IP allowlist.

- **Endpoint:** `GET /api/v1/ip-check/`
- **Required Scope:** None

### Response Example:
```json
{
  "success": true,
  "client_ip": "102.89.23.11",
  "is_allowed": true,
  "allowlist_enabled": true
}
```
