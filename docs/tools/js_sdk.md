# JavaScript & TypeScript SDK (`bursapay-sdk`)

The official JavaScript / TypeScript client library for Node.js, Next.js, and modern browser environments, built on native `fetch`.

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

## Quickstart (Node.js / TypeScript)

```typescript
import { BursaPay } from 'bursapay-sdk';

const bursapay = new BursaPay({
  secretKey: process.env.BURSAPAY_SECRET_KEY!,
});

async function run() {
  // 1. Initialize a checkout
  const payment = await bursapay.payments.initialize({
    amount: 25000,
    email: 'payer@example.com',
    reference: 'REF-NODE-001',
    callbackUrl: 'https://mysite.com/payment/complete'
  });

  console.log('Checkout URL:', payment.data.authorization_url);

  // 2. Query wallet balance
  const wallet = await bursapay.wallets.getBalance();
  console.log('Available Balance:', wallet.data.available_balance);
}

run();
```

---

## Webhook Verification in TypeScript

```typescript
import { verifyWebhookSignature } from 'bursapay-sdk';

const isValid = verifyWebhookSignature({
  rawBody: req.body, // Raw Buffer
  signature: req.headers['x-bursapay-signature'] as string,
  secretKey: process.env.BURSAPAY_SECRET_KEY!,
});
```
