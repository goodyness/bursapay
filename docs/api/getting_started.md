# Getting Started with the Developer Gateway API

The BursaPay Developer Gateway REST API allows software teams, startups, and institutions to embed African payment collection, dedicated virtual bank accounts, split transfers, automated subscriptions, and digital invoicing into their software stack.

---

## Base URLs

| Environment | Base URL |
|---|---|
| **Production / Live** | `https://api.bursapay.com/api/v1` |
| **Sandbox / Test** | `https://api.bursapay.com/api/v1` |

*Note: Environments are determined entirely by whether you authenticate with a Test Key (`bp_sec_test_KEY_HERE`) or a Live Key (`bp_sec_live_KEY_HERE`).*

---

## Authentication

Every API call to BursaPay is authenticated over HTTPS using standard Bearer token authentication in the `Authorization` HTTP header.

```http
Authorization: Bearer bp_sec_test_DEMO_KEY_HERE
```

### API Key Types

| Key Name | Prefix | Visibility | Permitted Context |
|---|---|---|---|
| **Secret Test Key** | `sk_test_` | Private | Backend servers / Sandbox environment |
| **Publishable Test Key** | `pk_test_` | Public | Mobile Apps / Frontend JS / Sandbox checkout |
| **Secret Live Key** | `sk_live_` | Private | Backend servers / Production environment |
| **Publishable Live Key** | `pk_live_` | Public | Mobile Apps / Frontend JS / Production checkout |

> ⚠️ **CRITICAL SECURITY RULE:** Never expose secret keys (`bp_sec_live_KEY_HERE` or `bp_sec_test_KEY_HERE`) in frontend code, client-side mobile bundles, or public git repositories.

---

## Standard Request Headers

```http
Authorization: Bearer bp_sec_test_DEMO_KEY_HERE
Content-Type: application/json
Accept: application/json
Idempotency-Key: <unique-uuid-v4>
```

---

## Standard Response Structure

All successful API responses return an HTTP `2xx` status code wrapped in a consistent JSON envelope:

```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": {
    "reference": "BP-PAY-9812456",
    "amount": "15000.00",
    "status": "success"
  }
}
```

### Standard Error Structure

Error responses return appropriate HTTP `4xx` or `5xx` status codes with structured diagnostic data:

```json
{
  "success": false,
  "error_code": "RESOURCE_NOT_FOUND",
  "message": "The requested payment reference does not exist",
  "details": {
    "reference": "Invalid reference 'BP-NON-EXISTENT'"
  }
}
```

---

## HTTP Status Codes Reference

| Status Code | Meaning | Description |
|---|---|---|
| `200 OK` | Success | Request succeeded and returned data. |
| `201 Created` | Created | Resource successfully created. |
| `400 Bad Request` | Validation Error | Request body failed schema validation. |
| `401 Unauthorized` | Invalid Key | Missing or invalid API key. |
| `403 Forbidden` | Forbidden | Key lacks the required permission scope. |
| `404 Not Found` | Not Found | Target resource was not found. |
| `409 Conflict` | Conflict / Duplicate | Resource already exists or duplicate idempotency key with different payload. |
| `429 Too Many Requests`| Rate Limited | Request quota exceeded for this key/IP. |
| `500 Server Error` | Internal Error | An unexpected server error occurred. |
