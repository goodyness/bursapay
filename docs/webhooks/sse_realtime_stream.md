# Real-Time Event Stream (Server-Sent Events)

For browser dashboards, command-line interfaces (such as `bursapay-cli`), and internal monitoring workers that cannot expose a public webhook endpoint, BursaPay provides a real-time event stream using standard **Server-Sent Events (SSE)**.

---

## Stream Endpoint

- **Endpoint:** `GET /api/v1/events/stream/`
- **Required Scope:** `payments:read`
- **Transport:** HTTP/1.1 or HTTP/2 SSE (`text/event-stream`)

### Stream Request Example

```bash
curl -N -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
     -H "Accept: text/event-stream" \
     https://api.bursapay.com/api/v1/events/stream/
```

---

## Stream Output Format

The stream pushes newline-delimited event chunks conforming to the SSE specification:

```http
event: payment.success
data: {"reference": "BP-ORD-90214", "amount": "15000.00", "currency": "NGN", "status": "success"}

event: virtual_account.credited
data: {"account_number": "0123456789", "amount": "50000.00", "currency": "NGN"}
```

---

## Consuming SSE in JavaScript

```javascript
const EventSource = require('eventsource'); // or native EventSource in browser

const eventSource = new EventSource('https://api.bursapay.com/api/v1/events/stream/', {
  headers: {
    'Authorization': 'Bearer bp_sec_test_DEMO_KEY_HERE'
  }
});

eventSource.addEventListener('payment.success', (e) => {
  const payment = JSON.parse(e.data);
  console.log('⚡ Real-time Payment Settled:', payment.reference, payment.amount);
});

eventSource.onerror = (err) => {
  console.error('SSE Connection error:', err);
};
```
