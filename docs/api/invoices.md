# Invoicing API

The Invoicing API enables B2B businesses, agencies, and institutions to issue professional, itemized digital invoices with automated email delivery, tax calculations, and instant pay links.

---

## 1. Create an Invoice

- **Endpoint:** `POST /api/v1/invoices/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `customer_reference` | `string` | **Yes** | BursaPay customer reference. |
| `due_date` | `string` | **Yes** | ISO-8601 due date (`YYYY-MM-DD`). |
| `currency` | `string` | No | Currency code (`NGN`, `USD`). Default: `NGN`. |
| `items` | `array` | **Yes** | List of line item objects. |
| `tax_percent` | `number` | No | Percentage tax (e.g. `7.5` for VAT). |
| `discount_amount` | `number` | No | Flat discount deducted from total. |
| `send_notification` | `boolean`| No | Automatically email invoice link to customer (Default: `true`). |

### Line Item Object Structure:
```json
{
  "description": "Enterprise API Integration & Setup",
  "quantity": 1,
  "unit_price": 250000.00
}
```

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/invoices/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_reference": "CUST-99214",
    "due_date": "2026-10-15",
    "tax_percent": 7.5,
    "items": [
      {
        "description": "Cloud Architecture Review",
        "quantity": 2,
        "unit_price": 75000.00
      },
      {
        "description": "Security Penetration Audit",
        "quantity": 1,
        "unit_price": 120000.00
      }
    ]
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "data": {
    "reference": "INV-2026-0814",
    "status": "pending",
    "subtotal": "270000.00",
    "tax_amount": "20250.00",
    "total_amount": "290250.00",
    "due_date": "2026-10-15",
    "payment_url": "https://checkout.bursapay.com/invoice/INV-2026-0814",
    "pdf_url": "https://api.bursapay.com/api/v1/invoices/INV-2026-0814/pdf/"
  }
}
```

---

## 2. Retrieve Invoice Details

- **Endpoint:** `GET /api/v1/invoices/<reference>/`
- **Required Scope:** `payments:read`
