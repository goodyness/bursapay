# JavaScript & TypeScript SDK (`bursapay-sdk`)

The official JavaScript / TypeScript client library for **Node.js 18+**, browser apps, Next.js, and serverless environments, featuring full TypeScript definitions, ESM + CommonJS builds, and zero external runtime dependencies (native `fetch`).

---

## Installation

```bash
npm install bursapay-sdk
# or
yarn add bursapay-sdk
# or
pnpm add bursapay-sdk
```

---

## Client Setup

```typescript
import { BursaPay } from 'bursapay-sdk';

// Initialize with your secret API key
const bursapay = new BursaPay({
  secretKey: process.env.BURSAPAY_SECRET_KEY!,
  // baseUrl: 'http://localhost:8000/api/v1', // Optional
});
```

---

## Payments

### 1. Initialize Hosted Checkout
```typescript
const payment = await bursapay.payments.initialize({
  amount: 25000.00,
  email: 'customer@example.com',
  currency: 'NGN',
  reference: 'ORD-JS-2026-001',
  callbackUrl: 'https://mysite.com/checkout/callback',
  metadata: {
    orderId: 'ORD-JS-2026-001',
  },
});

console.log('Authorization URL:', payment.data.authorization_url);
```

### 2. Verify Payment
```typescript
const verification = await bursapay.payments.verify('ORD-JS-2026-001');

if (verification.data.status === 'success') {
  console.log(`Payment of ₦${verification.data.amount} confirmed via ${verification.data.channel}`);
}
```

### 3. Charge Saved Card
```typescript
const charge = await bursapay.payments.chargeSavedCard({
  customerReference: 'CUST-99214',
  authorizationCode: 'AUTH_8b91fa02c',
  amount: 5000.00,
});
```

---

## Payment Intents (Authorize & Capture)

```typescript
// Create authorization intent
const intent = await bursapay.paymentIntents.create({
  amount: 50000.00,
  email: 'renter@carhire.com',
});

// Capture held funds
await bursapay.paymentIntents.capture(intent.data.intent_reference, {
  amountToCapture: 50000.00,
});

// Cancel authorization
// await bursapay.paymentIntents.cancel(intent.data.intent_reference);
```

---

## Dedicated Virtual Accounts

```typescript
const va = await bursapay.virtualAccounts.create({
  customerEmail: 'merchant@example.com',
  bvn: '12345678901',
  preferredBank: 'Wema Bank',
});

console.log(`Virtual Account: ${va.data.account_number} (${va.data.bank_name})`);
```

---

## Subscriptions & Invoicing

```typescript
// Create Plan
const plan = await bursapay.subscriptions.createPlan({
  name: 'Standard Pro Plan',
  amount: 15000.00,
  interval: 'monthly',
});

// Subscribe Customer
const sub = await bursapay.subscriptions.create({
  customerEmail: 'user@example.com',
  planCode: plan.data.plan_code,
});

// Pause / Resume / Cancel
await bursapay.subscriptions.pause(sub.data.id);
await bursapay.subscriptions.resume(sub.data.id);
await bursapay.subscriptions.cancel(sub.data.id);
```

---

## Webhook Verification in Express / Next.js

```typescript
import express from 'express';
import { verifyWebhookSignature } from 'bursapay-sdk';

const app = express();

app.post('/webhooks/bursapay', express.raw({ type: 'application/json' }), (req, res) => {
  const signature = req.headers['x-bursapay-signature'] as string;
  const timestamp = req.headers['x-bursapay-timestamp'] as string;
  const secret = process.env.BURSAPAY_WEBHOOK_SECRET!;

  const isValid = verifyWebhookSignature({
    rawBody: req.body, // Buffer
    signature,
    timestamp,
    secretKey: secret,
  });

  if (!isValid) {
    return res.status(400).json({ error: 'Invalid webhook signature' });
  }

  const event = JSON.parse(req.body.toString());

  switch (event.event) {
    case 'payment.success':
      console.log('Payment succeeded:', event.data.reference);
      break;
    case 'virtual_account.credited':
      console.log('Virtual account credited:', event.data.amount);
      break;
  }

  res.status(200).json({ received: true });
});
```

---

## Browser Inline Popup

When using in a frontend web application:

```typescript
import { BursaPay } from 'bursapay-sdk';

BursaPay.inline({
  key: 'bp_pub_test_DEMO_KEY_HERE',
  email: 'customer@example.com',
  amount: 15000.00,
  currency: 'NGN',
  onSuccess: (res) => {
    console.log('Payment completed:', res.reference);
  },
  onCancel: () => {
    console.log('Checkout dismissed');
  },
});
```
