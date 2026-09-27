import os
import hmac
import hashlib
import json
import httpx
from dotenv import load_dotenv
from fastapi import FastAPI, Header, HTTPException, Request, Response, status
from pydantic import BaseModel

load_dotenv()

app = FastAPI(title="BursaPay FastAPI Integration Starter")

BURSAPAY_SECRET_KEY = os.getenv("BURSAPAY_SECRET_KEY", "bp_sec_test_DEMO_KEY_HERE")
BURSAPAY_BASE_URL = "https://api.bursapay.com/api/v1"

class CheckoutRequest(BaseModel):
    amount: float
    email: str
    reference: str = None
    callback_url: str = "http://localhost:8000/checkout/success"

@app.post("/api/checkout")
async def create_checkout(payload: CheckoutRequest):
    """
    Initializes a BursaPay checkout session.
    """
    async with httpx.AsyncClient() as client:
        res = await client.post(
            f"{BURSAPAY_BASE_URL}/payments/initialize/",
            headers={
                "Authorization": f"Bearer {BURSAPAY_SECRET_KEY}",
                "Content-Type": "application/json"
            },
            json={
                "amount": payload.amount,
                "email": payload.email,
                "reference": payload.reference,
                "callback_url": payload.callback_url
            }
        )
        if res.status_code != 201:
            raise HTTPException(status_code=res.status_code, detail=res.json())
        return res.json()

@app.post("/api/webhooks")
async def handle_webhook(
    request: Request,
    x_bursapay_signature: str = Header(None)
):
    """
    Webhook handler validating HMAC-SHA256 signature against raw bytes.
    """
    if not x_bursapay_signature:
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Missing signature header")

    body_bytes = await request.body()
    
    expected_signature = hmac.new(
        key=BURSAPAY_SECRET_KEY.encode("utf-8"),
        msg=body_bytes,
        digestmod=hashlib.sha256
    ).hexdigest()

    if not hmac.compare_digest(expected_signature, x_bursapay_signature):
        raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Invalid webhook signature")

    event = json.loads(body_bytes.decode("utf-8"))
    event_type = event.get("event")
    event_data = event.get("data", {})

    print(f"✅ Verified Webhook Received: {event_type} -> Reference: {event_data.get('reference')}")

    # Process business logic here...

    return {"received": True}
