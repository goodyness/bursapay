"""
BursaPay SaaS Subscriptions & Recurring Billing Example (Flask)
"""
import os
from flask import Flask, request, jsonify
from bursapay import BursaPay
from bursapay.webhooks import verify_signature

app = Flask(__name__)

bp = BursaPay(secret_key=os.getenv("BURSAPAY_SECRET_KEY", "bp_sec_test_DEMO_KEY_HERE"))
WEBHOOK_SECRET = os.getenv("BURSAPAY_WEBHOOK_SECRET", "sec_wh_DEMO_SECRET_HERE")

@app.route("/api/create-plan", methods=["POST"])
def create_plan():
    data = request.json or {}
    plan = bp.subscriptions.create_plan(
        name=data.get("name", "Pro Plan"),
        amount=float(data.get("amount", 10000.00)),
        interval=data.get("interval", "monthly"),
        currency="NGN"
    )
    return jsonify(plan), 201

@app.route("/api/subscribe", methods=["POST"])
def subscribe():
    data = request.json or {}
    subscription = bp.subscriptions.create(
        customer_email=data.get("email"),
        plan_code=data.get("plan_code"),
    )
    return jsonify(subscription), 201

@app.route("/webhooks/bursapay", methods=["POST"])
def webhook():
    sig = request.headers.get("X-BursaPay-Signature")
    ts = request.headers.get("X-BursaPay-Timestamp")
    
    if not verify_signature(request.data, sig, ts, WEBHOOK_SECRET):
        return jsonify({"error": "Invalid signature"}), 400
        
    event = request.json
    event_type = event.get("event")
    
    if event_type == "subscription.charged":
        print(f"Recurring charge successful: {event['data']['subscription_code']}")
    elif event_type == "subscription.canceled":
        print(f"Subscription canceled: {event['data']['subscription_code']}")
        
    return jsonify({"received": True}), 200

if __name__ == "__main__":
    app.run(port=5000, debug=True)
