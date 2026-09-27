# Python SDK (`bursapay-sdk`)

The official Python client library for the BursaPay Developer Gateway API, featuring synchronous and asynchronous HTTP transports powered by `httpx`.

---

## Installation

```bash
pip install bursapay-sdk
```

---

## Quickstart (Synchronous)

```python
from bursapay import BursaPay

client = BursaPay(secret_key="bp_sec_test_DEMO_KEY_HERE")

# 1. Initialize a checkout session
payment = client.payments.initialize(
    amount=15000.00,
    email="customer@example.com",
    reference="ORDER-10023",
    callback_url="https://yourapp.com/callback"
)
print("Authorization URL:", payment.data.authorization_url)

# 2. Verify payment status
verified = client.payments.verify(reference="ORDER-10023")
print("Status:", verified.data.status)
```

---

## Asynchronous Usage (FastAPI / asyncio)

```python
import asyncio
from bursapay import AsyncBursaPay

async def main():
    async with AsyncBursaPay(secret_key="bp_sec_test_KEY_HERE") as client:
        # Create dedicated virtual account
        account = await client.virtual_accounts.create(
            customer_reference="CUST-4491",
            preferred_bank="wema-bank"
        )
        print("Generated NUBAN:", account.data.account_number)

asyncio.run(main())
```

---

## Handling Webhooks in Python

```python
from bursapay.webhooks import verify_signature

# Inside your Flask / FastAPI / Django route
is_valid = verify_signature(
    raw_body_bytes=request_body,
    signature_header=request.headers.get("X-BursaPay-Signature"),
    secret_key="bp_sec_live_KEY_HERE"
)
```
