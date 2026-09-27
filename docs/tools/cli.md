# BursaPay Developer CLI (`bursapay-cli`)

The official developer CLI for testing, tunneling webhooks, validating cryptographic signatures, and seeding sandbox test data.

---

## Installation

```bash
# Direct execution without permanent installation
npx bursapay-cli listen --forward-to http://localhost:8000/webhooks/

# Or install globally
npm install -g bursapay-cli
```

---

## Core Commands

### 1. `bursapay login` / `logout`
Save your developer API key securely in `~/.bursapay/config.json`.
```bash
bursapay login --api-key bp_sec_test_DEMO_KEY_HERE
```

---

### 2. `bursapay listen` (Webhook Tunneling)
Tunnels real-time webhook events directly from your BursaPay sandbox account to your local development server without exposing port 80/443 or configuring ngrok.

```bash
bursapay listen --forward-to http://localhost:5000/webhooks/bursapay

# Filter specific event types
bursapay listen --forward-to http://localhost:5000/webhooks/ --filter payment.success,virtual_account.credited
```

---

### 3. `bursapay trigger` (Mock Webhook Dispatcher)
Dispatch simulated webhook payloads to test your local handlers.

```bash
# Trigger payment success event
bursapay trigger payment.success --forward-to http://localhost:5000/webhooks/

# Trigger virtual account credit event
bursapay trigger virtual_account.credited --forward-to http://localhost:5000/webhooks/
```

---

### 4. `bursapay verify` (Signature Inspector)
Validate HMAC-SHA256 signatures offline against local payload files:

```bash
bursapay verify --payload ./payload.json --signature 7b31e9... --secret bp_sec_test_KEY_HERE
```
