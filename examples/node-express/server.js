require('dotenv').config();
const express = require('express');
const crypto = require('crypto');
const axios = require('axios');

const app = express();
const PORT = process.env.PORT || 5000;
const BURSAPAY_SECRET_KEY = process.env.BURSAPAY_SECRET_KEY || 'bp_sec_test_DEMO_KEY_HERE';

// 1. Initialize a checkout session
app.post('/api/create-checkout', express.json(), async (req, res) => {
  try {
    const { amount, email, reference } = req.body;
    
    const response = await axios.post(
      'https://api.bursapay.com/api/v1/payments/initialize/',
      {
        amount: amount || 10000.00,
        email: email || 'customer@example.com',
        reference: reference || `ORD-${Date.now()}`,
        callback_url: 'http://localhost:5000/callback'
      },
      {
        headers: {
          'Authorization': `Bearer ${BURSAPAY_SECRET_KEY}`,
          'Content-Type': 'application/json'
        }
      }
    );

    res.json(response.data);
  } catch (error) {
    res.status(error.response?.status || 500).json(error.response?.data || { error: error.message });
  }
});

// 2. Webhook receiver with cryptographic HMAC-SHA256 signature verification
app.post('/api/webhooks', express.raw({ type: 'application/json' }), (req, res) => {
  const signature = req.headers['x-bursapay-signature'];
  
  if (!signature) {
    return res.status(401).send('Missing X-BursaPay-Signature header');
  }

  const expectedHash = crypto
    .createHmac('sha256', BURSAPAY_SECRET_KEY)
    .update(req.body)
    .digest('hex');

  const isValid = crypto.timingSafeEqual(Buffer.from(expectedHash), Buffer.from(signature));

  if (!isValid) {
    console.error('❌ Invalid Webhook Signature');
    return res.status(401).send('Invalid signature');
  }

  const event = JSON.parse(req.body.toString('utf8'));
  console.log(`✅ Verified Webhook Event: ${event.event}`, event.data);

  // Handle specific event types
  switch (event.event) {
    case 'payment.success':
      console.log(`Payment confirmed for reference: ${event.data.reference}`);
      break;
    case 'virtual_account.credited':
      console.log(`Deposit detected on account: ${event.data.account_number}`);
      break;
    default:
      console.log(`Unhandled event type: ${event.event}`);
  }

  // Always return 200 OK promptly
  res.status(200).json({ received: true });
});

app.listen(PORT, () => {
  console.log(`🚀 BursaPay Node.js starter running at http://localhost:${PORT}`);
});
