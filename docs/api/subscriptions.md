# Subscriptions & Recurring Billing API

The Subscriptions API automates recurring card charges for SaaS platforms, gym memberships, co-working spaces, and recurring membership dues.

---

## 1. Create a Subscription Plan

Define a recurring billing interval and pricing tier.

- **Endpoint:** `POST /api/v1/subscriptions/plans/`
- **Required Scope:** `subscriptions:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `name` | `string` | **Yes** | Plan name (e.g., "Pro Developer Monthly"). |
| `amount` | `number` | **Yes** | Amount billed per cycle. |
| `interval` | `string` | **Yes** | Billing frequency: `weekly`, `monthly`, `quarterly`, `annually`. |
| `currency` | `string` | No | Currency code (`NGN`, `USD`). Default: `NGN`. |
| `description` | `string` | No | Plan description. |

### Example Request

```bash
curl -X POST https://api.bursapay.com/api/v1/subscriptions/plans/ \
  -H "Authorization: Bearer bp_sec_live_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cloud Workspace Pro",
    "amount": 18000.00,
    "interval": "monthly",
    "currency": "NGN"
  }'
```

---

## 2. Subscribe a Customer

Attach an active customer's card authorization token to a subscription plan.

- **Endpoint:** `POST /api/v1/subscriptions/`
- **Required Scope:** `subscriptions:write`

### Request Body

| Parameter | Type | Required | Description |
|---|---|---|---|
| `customer_reference`| `string` | **Yes** | Customer ID. |
| `plan_id` | `integer`| **Yes** | Target subscription plan ID. |
| `start_date` | `string` | No | ISO-8601 start date (default: immediate). |
| `authorization_code`| `string` | No | Specific tokenized card authorization code. |

---

## 3. Subscription Lifecycle Management

| Operation | Endpoint | Description |
|---|---|---|
| **Get Details** | `GET /api/v1/subscriptions/<id>/` | View current status, billing dates, and invoices. |
| **Pause** | `POST /api/v1/subscriptions/<id>/pause/` | Temporarily halt automated recurring charges. |
| **Resume** | `POST /api/v1/subscriptions/<id>/resume/` | Reactivate paused recurring billing schedule. |
| **Cancel** | `POST /api/v1/subscriptions/<id>/cancel/` | Permanently terminate subscription. |

### Lifecycle State Machine:
```
[active] ──(pause)──► [paused] ──(resume)──► [active]
   │
 (cancel)
   ▼
[cancelled]
```
