# Webhook Signature Verification

Every webhook request sent by BursaPay includes an `X-BursaPay-Signature` HTTP header. This header contains the HMAC-SHA256 signature of the raw request payload, computed using your account's secret key or endpoint webhook secret.

> 🔒 **Always verify webhook signatures** to ensure requests originated from BursaPay and have not been intercepted or forged.

---

## Verification Logic

1. Read the **raw, unparsed** HTTP request body bytes.
2. Compute the HMAC-SHA256 hash using your secret key (`bp_sec_live_KEY_HERE` or endpoint secret).
3. Compare the computed hex digest against the value in `X-BursaPay-Signature` using a constant-time comparison function.

---

## Code Examples

### 1. Python (FastAPI / Flask / Django)

```python
import hmac
import hashlib

def verify_bursapay_webhook(raw_payload_bytes: bytes, signature_header: str, secret_key: str) -> bool:
    """
    Verifies that a webhook payload was genuinely signed by BursaPay.
    """
    expected_signature = hmac.new(
        key=secret_key.encode('utf-8'),
        msg=raw_payload_bytes,
        digestmod=hashlib.sha256
    ).hexdigest()
    
    return hmac.compare_digest(expected_signature, signature_header)
```

### 2. JavaScript / Node.js (Express)

```javascript
const crypto = require('crypto');

function verifyBursaPayWebhook(rawBodyBuffer, signatureHeader, secretKey) {
  const hash = crypto
    .createHmac('sha256', secretKey)
    .update(rawBodyBuffer)
    .digest('hex');
    
  return crypto.timingSafeEqual(Buffer.from(hash), Buffer.from(signatureHeader));
}

// In Express: use express.raw({ type: 'application/json' }) to retain raw body
app.post('/webhooks/bursapay', express.raw({ type: 'application/json' }), (req, res) => {
  const signature = req.headers['x-bursapay-signature'];
  if (!verifyBursaPayWebhook(req.body, signature, process.env.BURSAPAY_SECRET_KEY)) {
    return res.status(401).send('Invalid signature');
  }
  
  const event = JSON.parse(req.body.toString('utf8'));
  console.log('Received event:', event.event);
  res.status(200).json({ received: true });
});
```

### 3. PHP

```php
<?php
function verifyBursaPayWebhook($rawPayload, $signatureHeader, $secretKey) {
    $expectedSignature = hash_hmac('sha256', $rawPayload, $secretKey);
    return hash_equals($expectedSignature, $signatureHeader);
}

$rawPayload = file_get_contents('php://input');
$signatureHeader = $_SERVER['HTTP_X_BURSAPAY_SIGNATURE'] ?? '';

if (!verifyBursaPayWebhook($rawPayload, $signatureHeader, getenv('BURSAPAY_SECRET_KEY'))) {
    http_response_code(401);
    exit('Invalid signature');
}

$event = json_decode($rawPayload, true);
http_response_code(200);
echo json_encode(['received' => true]);
```

### 4. Go (Golang)

```go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func VerifyBursaPayWebhook(rawPayload []byte, signatureHeader string, secretKey string) bool {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(rawPayload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedMAC), []byte(signatureHeader))
}
```
