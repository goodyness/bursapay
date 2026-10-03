# Webhooks Overview & Architecture

BursaPay uses webhooks to notify your backend asynchronously whenever critical events occur on your account — such as successful payments, dedicated virtual account credits, dispute alerts, subscription billing renewals, and outbound transfer status updates.

---

## Webhook Delivery Characteristics

- **Transport:** HTTP `POST` requests with a `Content-Type: application/json` header.
- **Security:** Every payload is signed with an HMAC SHA-256 signature in the `X-BursaPay-Signature` and timestamped in `X-BursaPay-Timestamp`.
- **Automatic Retry Policy:** Exponential backoff with jitter across **5 retry attempts** (`1m`, `5m`, `15m`, `1h`, `6h`).
- **Dead Letter Queue (DLQ):** Webhooks that exhaust all 5 retries are captured in the merchant's Dead Letter Queue for inspection and manual 1-click replay.
- **Response Expectation:** Your endpoint must return an HTTP status code in the `2xx` range (such as `200 OK` or `202 Accepted`) within **5 seconds**.

---

## 1. Webhook Management Endpoints

| Method | Endpoint | Required Scope | Description |
|---|---|---|---|
| `GET` | `/api/v1/webhooks/` | `webhooks:write` | List configured webhook endpoints. |
| `POST` | `/api/v1/webhooks/` | `webhooks:write` | Register a new webhook endpoint URL. |
| `GET` | `/api/v1/webhooks/<id>/` | `webhooks:write` | Retrieve webhook endpoint details & secret. |
| `PUT` / `PATCH` | `/api/v1/webhooks/<id>/` | `webhooks:write` | Update target URL or subscribed events. |
| `DELETE` | `/api/v1/webhooks/<id>/` | `webhooks:write` | Remove webhook endpoint. |
| `POST` | `/api/v1/webhooks/test/` | `webhooks:write` | Dispatch mock test event to endpoint. |

---

## 2. Webhook Delivery Logs & Manual Retry

Inspect historical delivery attempts and HTTP response status codes:

- **List Logs for Webhook:** `GET /api/v1/webhooks/<id>/logs/`
- **View Log Payload & Response:** `GET /api/v1/webhooks/logs/<log_id>/`
- **Manually Retry Specific Delivery:** `POST /api/v1/webhooks/logs/<log_id>/retry/`

---

## 3. Dead Letter Queue (DLQ)

When your server is temporarily down or returning 500 errors across all automatic retries, failed events are preserved in the DLQ:

- **List Dead Letters:** `GET /api/v1/webhooks/dead-letters/`
- **Replay Dead Letter:** `POST /api/v1/webhooks/dead-letters/<pk>/replay/`

---

## 4. Real-Time Server-Sent Events (SSE) Stream

For low-latency local development or reactive browser UIs without public webhook URLs, BursaPay provides an authenticated SSE stream:

- **Endpoint:** `GET /api/v1/events/stream/`
- **Required Scope:** `payments:read`
- **Header:** `Authorization: Bearer bp_sec_...`
- **Guide:** [Read SSE Stream Documentation](sse_realtime_stream.md)
