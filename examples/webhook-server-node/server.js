const express = require('express');
const crypto = require('crypto');

const app = express();
const PORT = process.env.PORT || 3000;
const WEBHOOK_SECRET = process.env.BURSAPAY_WEBHOOK_SECRET || 'sec_wh_DEMO_SECRET_HERE';

// Capture raw body Buffer before parsing
app.use(express.raw({ type: 'application/json' }));

function verifySignature(payload, signature, timestamp, secret) {
  if (!signature || !secret) return false;

  let ts = timestamp;
  let expectedHex = signature;

  if (signature.startsWith('t=')) {
    const parts = {};
    signature.split(',').forEach(seg => {
      const idx = seg.indexOf('=');
      if (idx > 0) parts[seg.substring(0, idx)] = seg.substring(idx + 1);
    });
    ts = parts['t'];
    expectedHex = parts['v1'];
  }

  if (!ts || !expectedHex) return false;

  const parsed = JSON.parse(payload.toString('utf8'));
  // Sorted keys serialization
  const compact = JSON.stringify(parsed, Object.keys(parsed).sort());
  const message = `${ts}.${compact}`;

  const hmac = crypto.createHmac('sha256', secret);
  hmac.update(message);
  const computedHex = hmac.digest('hex');

  return crypto.timingSafeEqual(Buffer.from(computedHex, 'utf8'), Buffer.from(expectedHex, 'utf8'));
}

app.post('/webhooks/bursapay', (req, res) => {
  const sig = req.headers['x-bursapay-signature'];
  const ts = req.headers['x-bursapay-timestamp'];

  if (!verifySignature(req.body, sig, ts, WEBHOOK_SECRET)) {
    console.warn('Webhook signature check failed');
    return res.status(401).json({ error: 'Unauthorized signature' });
  }

  const event = JSON.parse(req.body.toString('utf8'));
  console.log(`[Webhook Event Received] ${event.event}:`, event.data);

  switch (event.event) {
    case 'payment.success':
      console.log(`Payment confirmed: ${event.data.reference} for ₦${event.data.amount}`);
      break;
    case 'virtual_account.credited':
      console.log(`Virtual Account credited: ₦${event.data.amount}`);
      break;
    case 'transfer.success':
      console.log(`Transfer completed: ${event.data.transfer_reference}`);
      break;
  }

  return res.status(200).json({ received: true });
});

app.listen(PORT, () => {
  console.log(`BursaPay Webhook Server listening on port ${PORT}`);
});
