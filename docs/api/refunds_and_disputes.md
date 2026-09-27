# Refunds & Disputes API

The Refunds and Disputes API allows merchants to programmatically issue full or partial charge reversals and submit evidence to defend against customer chargebacks.

---

## 1. Issue a Refund

Reverse a settled transaction back to the customer's funding card or bank account.

- **Endpoint:** `POST /api/v1/refunds/`
- **Required Scope:** `refunds:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `transaction_reference`| `string` | **Yes** | Reference of the successful transaction to refund. |
| `amount` | `number` | No | Amount to refund. Omit for full 100% refund. |
| `reason` | `string` | No | Reason: `customer_request`, `duplicate_charge`, `fraudulent`, `service_cancelled`. |
| `merchant_note` | `string` | No | Internal note for reconciliation. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/refunds/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "transaction_reference": "BP-ORD-90214",
    "amount": 5000.00,
    "reason": "customer_request",
    "merchant_note": "Customer downgraded subscription tier"
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "message": "Refund processed successfully",
  "data": {
    "refund_reference": "REF-2026-88190",
    "transaction_reference": "BP-ORD-90214",
    "amount": "5000.00",
    "status": "completed",
    "currency": "NGN",
    "refunded_at": "2026-09-27T02:17:00Z"
  }
}
```

---

## 2. Disputes & Chargeback Evidence

When a cardholder files a chargeback through their issuing bank, BursaPay notifies the merchant via the `dispute.created` webhook event.

### List Active Disputes
- **Endpoint:** `GET /api/v1/disputes/`
- **Required Scope:** `payments:read`

### Submit Dispute Defense Evidence
- **Endpoint:** `POST /api/v1/disputes/<reference>/`
- **Required Scope:** `payments:write`

```json
{
  "explanation": "Customer verified delivery of physical event wristband on 2026-09-20.",
  "evidence_files": [
    "https://cdn.merchant.com/proofs/delivery_signature.pdf"
  ]
}
```
