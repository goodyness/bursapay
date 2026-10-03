# Payment Links & A/B Testing API

The Payment Links API enables merchants to programmatically generate hosted checkout pages, bulk links, QR codes, and run split-test experiments without building custom payment UI.

---

## 1. Create a Payment Link

- **Endpoint:** `POST /api/v1/payment-links/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `title` | `string` | **Yes** | Public title displayed on the checkout page. |
| `description` | `string` | No | Markdown-formatted product or service details. |
| `amount` | `number` | No | Fixed payment amount. Leave blank for flexible donations. |
| `currency` | `string` | No | `NGN`, `USD`, `GHS`. Default: `NGN`. |
| `custom_slug` | `string` | No | Custom vanity path (e.g. `tech-summit-pass`). Auto-generated if omitted. |
| `redirect_url`| `string` | No | Redirection destination after successful payment. |
| `is_reusable` | `boolean`| No | Whether multiple payers can use the link (`true` by default). |
| `expires_at`  | `string` | No | ISO 8601 expiry timestamp. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/payment-links/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Design System Workshop 2026",
    "description": "2-day hands-on workshop covering advanced Figma & React design systems.",
    "amount": 35000.00,
    "currency": "NGN",
    "custom_slug": "design-workshop-2026",
    "redirect_url": "https://designclub.org/workshop-confirmed"
  }'
```

---

## 2. Bulk Payment Link Creation

Create multiple payment links in a single request.

- **Endpoint:** `POST /api/v1/payment-links/bulk/`
- **Required Scope:** `payments:write`

### Request Body
```json
{
  "links": [
    {
      "title": "Early Bird Ticket",
      "amount": 10000.00,
      "custom_slug": "early-bird-pass"
    },
    {
      "title": "VIP All-Access Ticket",
      "amount": 35000.00,
      "custom_slug": "vip-all-access"
    }
  ]
}
```

---

## 3. Generate Payment Link QR Code

Generate a downloadable high-resolution QR image pointing directly to the payment page.

- **Endpoint:** `GET /api/v1/payment-links/<link_code>/qr/`
- **Required Scope:** `payments:read`
- **Response:** PNG QR Code image stream or data URI.

---

## 4. Link Analytics & Conversion Tracking

Inspect views, unique visits, conversion rate, and total volume collected per payment link.

- **Endpoint:** `GET /api/v1/payment-links/<link_code>/analytics/`
- **Required Scope:** `payments:read`

### Response Example:
```json
{
  "success": true,
  "data": {
    "link_code": "design-workshop-2026",
    "total_views": 1420,
    "unique_visitors": 1180,
    "successful_payments": 142,
    "conversion_rate": "12.03%",
    "total_volume_collected": "4970000.00",
    "currency": "NGN"
  }
}
```

---

## 5. Payment Link A/B Split Testing

Create an automated traffic split between two different checkout link variants to optimize conversion rate.

- **Endpoint:** `POST /api/v1/payment-links/ab-test/`
- **Required Scope:** `payments:write`

### Request Body
```json
{
  "name": "Checkout Copy Optimization Experiment",
  "link_a_id": 102,
  "link_b_id": 105,
  "traffic_split": 50
}
```

---

## 6. List, Retrieve, Update & Deactivate

| Action | Method | Endpoint | Required Scope |
|---|---|---|---|
| **List Links** | `GET` | `/api/v1/payment-links/` | `payments:read` |
| **Retrieve Link** | `GET` | `/api/v1/payment-links/<link_code>/` | `payments:read` |
| **Update Link** | `PUT` / `PATCH` | `/api/v1/payment-links/<link_code>/` | `payments:write` |
| **Deactivate Link** | `DELETE` | `/api/v1/payment-links/<link_code>/` | `payments:write` |
