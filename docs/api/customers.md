# Customers API

The Customers API provides a unified CRM interface to create customer entities, track lifetime transaction volume, attach reusable card tokens, manage dedicated virtual NUBAN accounts, and bulk import records.

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

## 2. List & Retrieve Customers

### List Customers
- **Endpoint:** `GET /api/v1/customers/`
- **Required Scope:** `payments:read`
- **Query Params:** `search`, `page`, `page_size`

### Retrieve Customer Detail
- **Endpoint:** `GET /api/v1/customers/<customer_reference>/`
- **Required Scope:** `payments:read`

### Update Customer Profile
- **Endpoint:** `PUT` / `PATCH /api/v1/customers/<customer_reference>/`
- **Required Scope:** `payments:write`

---

## 3. Bulk Customer Import

Import dozens or hundreds of customer records simultaneously.

- **Endpoint:** `POST /api/v1/customers/import/`
- **Required Scope:** `payments:write`

### Request Body
```json
{
  "customers": [
    {
      "email": "student1@funaab.edu.ng",
      "first_name": "Bolu",
      "last_name": "Ade",
      "phone": "+2348011111111",
      "metadata": {"matric": "20201001"}
    },
    {
      "email": "student2@funaab.edu.ng",
      "first_name": "Kehinde",
      "last_name": "Ojo",
      "phone": "+2348022222222",
      "metadata": {"matric": "20201002"}
    }
  ]
}
```

---

## 4. Customer Sub-Resources

### 4.1 List Customer Transaction History
- **Endpoint:** `GET /api/v1/customers/<customer_reference>/payments/`
- **Required Scope:** `payments:read`

### 4.2 List Saved Cards (Tokenized Authorizations)
- **Endpoint:** `GET /api/v1/customers/<customer_reference>/saved-cards/`
- **Required Scope:** `payments:read`
- **Response:** Array of reusable authorization codes (`AUTH_xxx`), card brand, last 4 digits, expiry date, issuing bank.

### 4.3 List Customer Virtual Accounts
- **Endpoint:** `GET /api/v1/customers/<customer_reference>/virtual-accounts/`
- **Required Scope:** `payments:read`
- **Response:** Dedicated virtual NUBAN accounts associated with this customer.
