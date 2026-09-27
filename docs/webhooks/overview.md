# Webhooks Overview & Architecture

Webhooks notify your backend systems asynchronously in real-time whenever an event occurs on BursaPay (e.g. payment received, virtual account credited, transfer completed, dispute opened).

---

## Webhook Delivery Architecture

```
┌────────────────────────┐
│ BursaPay Event Triggers│ (e.g. Card Payment, NIP Transfer)
└───────────┬────────────┘
            │
            ▼
┌────────────────────────┐
│ Async Celery Dispatcher│ ── Signs payload with HMAC-SHA256
└───────────┬────────────┘
            │
            ├───────────────► Merchant Webhook URL (`POST https://yourdomain.com/webhooks/`)
            │                 └─ Returns `HTTP 200 OK`
            │
            ▼ (If merchant server fails or times out)
┌────────────────────────┐
│ Exponential Backoff    │ (Retries at 5m, 15m, 1h, 6h, 24h)
└───────────┬────────────┘
            │
            ▼ (After max retries)
┌────────────────────────┐
│ Dead Letter Queue (DLQ)│ ── Inspect and replay anytime via API or Portal
└────────────────────────┘
```

---

## Delivery Guarantees

1. **At-Least-Once Delivery**: BursaPay guarantees events will be delivered at least once. Your endpoint handler should be idempotent (deduplicating using `event.data.reference` or `event.id`).
2. **Timeout Window**: Webhook deliveries timeout after **10 seconds**. Ensure your handler processes heavy tasks asynchronously (e.g., in a background worker queue) and immediately returns `HTTP 200 OK`.
3. **Dead Letter Queue (DLQ)**: Failed deliveries after 5 exponential backoff retries are preserved in the DLQ (`GET /api/v1/webhooks/dead-letters/`) and can be replayed at any time.

---

## Configuring Webhooks via API

### Register a Webhook Endpoint

- **Endpoint:** `POST /api/v1/webhooks/`
- **Required Scope:** `webhooks:write`

```bash
curl -X POST https://api.bursapay.com/api/v1/webhooks/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://api.yourdomain.com/v1/bursapay-webhooks",
    "events": [
      "payment.success",
      "virtual_account.credited",
      "transfer.success",
      "dispute.created"
    ],
    "description": "Production Webhook Ingestion"
  }'
```
