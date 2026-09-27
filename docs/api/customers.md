# Customers API

The Customers API provides a unified CRM interface to create customer entities, track lifetime transaction volume, attach reusable card tokens, and store KYC identity data.

---

## 1. Create a Customer

- **Endpoint:** `POST /api/v1/customers/`
- **Required Scope:** `payments:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `email` | `string` | **Yes** | Customer email address. |
| `first_name` | `string` | **Yes** | Customer first name. |
| `last_name` | `string` | **Yes** | Customer surname. |
| `phone` | `string` | No | E.164 formatted phone number (e.g., `+2348012345678`). |
| `metadata` | `object` | No | Custom JSON attributes (e.g. matric number, student level). |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/customers/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "chidi.okonkwo@university.edu.ng",
    "first_name": "Chidi",
    "last_name": "Okonkwo",
    "phone": "+2348039876543",
    "metadata": {
      "matric_number": "ENG/2022/0491",
      "department": "Mechanical Engineering",
      "level": "400L"
    }
  }'
```

### Example Response (`201 Created`)

```json
{
  "success": true,
  "data": {
    "customer_reference": "CUST-ENG-0491",
    "email": "chidi.okonkwo@university.edu.ng",
    "first_name": "Chidi",
    "last_name": "Okonkwo",
    "phone": "+2348039876543",
    "total_spend": "0.00",
    "total_transactions": 0,
    "created_at": "2026-09-27T02:18:00Z"
  }
}
```

---

## 2. Customer Sub-Resources

| Endpoint | Method | Purpose |
|---|---|---|
| `/api/v1/customers/<ref>/payments/` | `GET` | Retrieve complete transaction history for this customer. |
| `/api/v1/customers/<ref>/saved-cards/` | `GET` | List tokenized payment cards available for 1-click charge. |
| `/api/v1/customers/<ref>/virtual-accounts/` | `GET` | List dedicated NUBAN bank accounts linked to customer. |
| `/api/v1/customers/import/` | `POST` | Bulk import hundreds of customer records simultaneously. |
