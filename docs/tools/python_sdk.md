# Python SDK (`bursapay-sdk`)

The official Python client library for the **BursaPay Developer Gateway API**, featuring both synchronous (`BursaPay`) and asynchronous (`AsyncBursaPay`) HTTP clients powered by `httpx`.

---

## Installation

```bash
pip install bursapay-sdk
```

Tested on Python 3.8, 3.9, 3.10, 3.11, and 3.12.

---

## Client Setup

```python
from bursapay import BursaPay

# Standard synchronous client
bp = BursaPay(secret_key="bp_sec_test_DEMO_KEY_HERE")

# Optional: set custom base_url or timeout
# bp = BursaPay(secret_key="bp_sec_...", base_url="http://localhost:8000/api/v1", timeout=30.0)
```

For asynchronous applications (FastAPI, Starlette, asyncio):

```python
import asyncio
from bursapay import AsyncBursaPay

async def main():
    async with AsyncBursaPay(secret_key="bp_sec_test_DEMO_KEY_HERE") as bp:
        payment = await bp.payments.initialize(
            amount=5000.00,
            email="customer@example.com"
        )
        print("Async URL:", payment["data"]["authorization_url"])

asyncio.run(main())
```

---

## Payments

### Initialize Payment Session
```python
payment = bp.payments.initialize(
    amount=15000.00,
    email="customer@example.com",
    currency="NGN",
    reference="ORD-PY-2026-001",
    callback_url="https://yourapp.com/checkout/callback",
    metadata={"cart_id": "CART-9912"},
    channels=["card", "bank_transfer", "ussd"]
)
print("Authorization URL:", payment["data"]["authorization_url"])
```

### Verify Payment
```python
result = bp.payments.verify(reference="ORD-PY-2026-001")
if result["data"]["status"] == "success":
    print(f"Verified ₦{result['data']['amount']} paid via {result['data']['channel']}")
```

### Charge Saved Card (1-Click Recurring)
```python
charge = bp.payments.charge_saved_card(
    customer_reference="CUST-99214",
    authorization_code="AUTH_8b91fa02c",
    amount=5000.00,
    reference="RECURRING-001"
)
```

### Bulk Payments
```python
bulk = bp.payments.bulk_initialize([
    {"amount": 10000.00, "email": "user1@example.com", "reference": "BATCH-001"},
    {"amount": 20000.00, "email": "user2@example.com", "reference": "BATCH-002"},
])
batch_ref = bulk["data"]["batch_reference"]

# Check bulk status
status = bp.payments.get_bulk_status(batch_ref)
```

### Cancel Scheduled Charge & Mark Fulfillment
```python
# Cancel scheduled payment
bp.payments.cancel_schedule(reference="ORD-SCHEDULED-12")

# Mark order fulfillment
bp.payments.mark_fulfillment(reference="ORD-PY-2026-001", fulfillment_data={"tracking_number": "DHL-9812"})
```

---

## Payment Intents (2-Step Authorize & Capture)

```python
# 1. Create intent (hold funds)
intent = bp.payment_intents.create(
    amount=25000.00,
    email="guest@hotel.com",
    currency="NGN"
)
intent_ref = intent["data"]["intent_reference"]

# 2. Capture held funds after checkout/service completion
captured = bp.payment_intents.capture(
    intent_reference=intent_ref,
    amount_to_capture=25000.00
)

# 3. Or cancel authorization
# bp.payment_intents.cancel(intent_reference=intent_ref)
```

---

## Dedicated Virtual Accounts

```python
# Create dedicated NUBAN for customer
va = bp.virtual_accounts.create(
    customer_email="merchant@example.com",
    bvn="12345678901",
    preferred_bank="Wema Bank"
)
print(f"Virtual Account: {va['data']['account_number']} ({va['data']['bank_name']})")

# List virtual accounts
accounts = bp.virtual_accounts.list()
```

---

## Customers

```python
# Create customer
customer = bp.customers.create(
    email="uche.ogbonna@domain.com",
    first_name="Uche",
    last_name="Ogbonna",
    phone="+2348031234567"
)

# List customer payments
payments = bp.customers.get_payments(customer_reference="CUST-99214")

# List tokenized saved cards
cards = bp.customers.get_saved_cards(customer_reference="CUST-99214")
```

---

## Subscriptions & Invoicing

```python
# Create billing plan
plan = bp.subscriptions.create_plan(
    name="Enterprise Monthly",
    amount=50000.00,
    interval="monthly"
)

# Subscribe customer
sub = bp.subscriptions.create(
    customer_email="enterprise@client.com",
    plan_code=plan["data"]["plan_code"]
)

# Pause / Resume / Cancel subscription
bp.subscriptions.pause(subscription_id=sub["data"]["id"])
bp.subscriptions.resume(subscription_id=sub["data"]["id"])
bp.subscriptions.cancel(subscription_id=sub["data"]["id"])
```

---

## Transfers & Disbursements

```python
transfer = bp.transfers.create(
    amount=50000.00,
    account_number="0123456789",
    bank_code="058",
    narration="Contractor payment",
    reference="TRF-PY-9902"
)
```

---

## Webhook Signature Verification

```python
from bursapay.webhooks import verify_signature

# Inside Flask / FastAPI / Django
is_valid = verify_signature(
    raw_body_bytes=request.body,
    signature_header=request.headers.get("X-BursaPay-Signature"),
    timestamp_header=request.headers.get("X-BursaPay-Timestamp"),
    webhook_secret="sec_wh_DEMO_SECRET_HERE"
)

if not is_valid:
    raise ValueError("Invalid signature")
```

---

## Error Handling

```python
from bursapay.exceptions import (
    BursaPayError,
    AuthenticationError,
    InvalidRequestError,
    RateLimitError
)

try:
    bp.payments.initialize(amount=-10, email="invalid")
except InvalidRequestError as e:
    print(f"Validation error: {e.message} (Param: {e.param})")
except AuthenticationError:
    print("Check your API Secret Key.")
except RateLimitError:
    print("Rate limit reached, backing off.")
except BursaPayError as e:
    print(f"BursaPay error: {e}")
```
