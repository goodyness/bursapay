# Frontend Inline & Popup Checkout (`bursapay-embed.js`)

BursaPay provides seamless client-side modal checkout capabilities so your customers can complete payments directly on your website without being redirected to an external page.

---

## 1. Quick Integration via `<script>` Tag

Include the BursaPay JavaScript SDK in your HTML `<head>` or before the closing `</body>` tag:

```html
<script src="https://api.bursapay.com/bursapay-embed.js"></script>
```

*(Alternatively, use the canonical inline script: `https://api.bursapay.com/v1/bursapay.js`)*

---

## 2. Triggering Popup Checkout

Attach the handler to your checkout button:

```html
<button id="pay-btn" class="btn btn-primary">Pay ₦15,000</button>

<script>
document.getElementById('pay-btn').addEventListener('click', function () {
  const handler = BursaPay.setup({
    key: 'bp_pub_test_DEMO_KEY_HERE', // Your BursaPay Public Key
    email: 'customer@example.com',
    amount: 15000.00,                 // Amount in NGN
    currency: 'NGN',
    reference: 'ORD-' + Math.floor((Math.random() * 1000000000) + 1),
    metadata: {
      custom_fields: [
        { display_name: "Cart ID", variable_name: "cart_id", value: "CART-9921" }
      ]
    },
    onSuccess: function (response) {
      // response contains { reference: "ORD-...", status: "success" }
      console.log('Payment Successful:', response);
      window.location.href = '/order-success?ref=' + response.reference;
    },
    onCancel: function () {
      console.log('Customer closed payment modal.');
    },
    onError: function (error) {
      console.error('Payment error:', error);
      alert('Payment could not be processed: ' + error.message);
    }
  });

  // Open the interactive payment modal
  handler.openIframe();
});
</script>
```

---

## 3. React / Next.js Component Example

```tsx
import React from 'react';

declare global {
  interface Window {
    BursaPay: any;
  }
}

export const CheckoutButton = ({ amount, email }: { amount: number; email: string }) => {
  const handlePayment = () => {
    if (typeof window.BursaPay === 'undefined') {
      alert('BursaPay script not loaded yet.');
      return;
    }

    const handler = window.BursaPay.setup({
      key: process.env.NEXT_PUBLIC_BURSAPAY_KEY,
      email,
      amount,
      currency: 'NGN',
      onSuccess: (res: any) => {
        alert(`Payment successful! Reference: ${res.reference}`);
      },
      onCancel: () => {
        console.log('User dismissed checkout');
      }
    });

    handler.openIframe();
  };

  return (
    <button onClick={handlePayment} className="px-6 py-3 bg-green-600 text-white font-semibold rounded-lg shadow hover:bg-green-700">
      Pay with BursaPay
    </button>
  );
};
```

---

## 4. Configuration Options

| Parameter | Type | Required | Description |
|---|---|---|---|
| `key` | `string` | **Yes** | Your BursaPay Public Key (`bp_pub_test_...` or `bp_pub_live_...`). |
| `email` | `string` | **Yes** | Payer email address for receipt delivery. |
| `amount` | `number` | **Yes** | Amount in Naira (e.g. `5000.00` for ₦5,000). |
| `currency` | `string` | No | Currency code (`NGN`, `USD`, `GHS`, `KES`). Default: `NGN`. |
| `reference` | `string` | No | Unique transaction reference. |
| `channels` | `array` | No | Allowed payment methods: `['card', 'bank_transfer', 'ussd', 'qr']`. |
| `metadata` | `object` | No | Additional custom key-value data. |
| `onSuccess` | `function` | **Yes** | Callback executed when payment succeeds. |
| `onCancel` | `function` | No | Callback executed when customer closes modal. |
| `onError` | `function` | No | Callback executed on unexpected transaction error. |
