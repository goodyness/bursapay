# Security, Scopes & Idempotency

BursaPay provides advanced security controls to ensure high-value financial operations remain safe, tamper-proof, and resilient against network failures.

---

## 1. Idempotency (`Idempotency-Key`)

To prevent accidental double-charging during network timeouts, client retries, or server restarts, pass a unique `Idempotency-Key` header with mutation requests (`POST`, `PUT`, `PATCH`).

```http
POST /api/v1/payments/initialize/ HTTP/1.1
Host: api.bursapay.com
Authorization: Bearer bp_sec_live_KEY_HERE
Idempotency-Key: e82b36a1-944a-4d1e-8e40-fcbb12345678
Content-Type: application/json

{
  "amount": 50000.00,
  "email": "customer@example.com"
}
```

### Idempotency Behavior
1. **First Request**: BursaPay executes the operation, caches the HTTP response payload and status code, and associates it with the key for **24 hours**.
2. **Subsequent Retries with Same Key and Same Payload**: BursaPay returns the exact cached response without re-executing the payment charge or database mutation.
3. **Mismatched Payload Conflict**: If a subsequent request reuses a key with a altered payload, BursaPay rejects the call with an HTTP `409 Conflict`.

---

## 2. Granular API Key Scopes

You can restrict API keys to least-privilege permissions in the Developer Portal:

| Scope Identifier | Permitted Operations |
|---|---|
| `payments:read` | Query transactions, list payments, verify payment status. |
| `payments:write` | Initialize payments, charge cards, create payment intents. |
| `transfers:read` | View transfer history, check wallet balance. |
| `transfers:write` | Initiate single and bulk bank transfers / payouts. |
| `refunds:write` | Issue full and partial refunds to customers. |
| `virtual_accounts:write` | Provision dedicated and dynamic NUBAN bank accounts. |
| `subscriptions:write` | Create and manage recurring billing plans and subscriptions. |
| `webhooks:write` | Create, update, rotate secrets, and delete webhook endpoints. |

If an API key attempts an operation outside its assigned scope, BursaPay responds with:
```json
{
  "success": false,
  "error_code": "INSUFFICIENT_SCOPE",
  "message": "Your API key lacks the required scope: 'transfers:write'"
}
```

---

## 3. IP Allowlisting

For high-security server environments (such as dedicated payout services), you can restrict your secret keys to specific static IP addresses or CIDR blocks:
- Configurable under **Developer Portal > Settings > IP Allowlist**.
- Requests originating from non-whitelisted IPs receive an immediate `403 Forbidden` response.

---

## 4. Rate Limiting

To maintain system stability, BursaPay enforces the following default rate limits per API key:

| Endpoint Category | Sandbox Limit | Production Limit |
|---|---|---|
| **Standard Reads (`GET`)** | 120 req/min | 600 req/min |
| **Payment Creation (`POST`)** | 60 req/min | 300 req/min |
| **Transfers & Payouts (`POST`)** | 30 req/min | 100 req/min |
| **Real-time SSE Stream (`GET`)** | 5 active streams | 25 active streams |

When a rate limit is exceeded, BursaPay responds with `HTTP 429 Too Many Requests` containing standard `Retry-After` headers.
