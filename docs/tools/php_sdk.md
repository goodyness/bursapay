# PHP SDK (`bursapay-php`)

The official PHP SDK for integrating with the **BursaPay Developer Gateway API**, featuring PSR-7/18 compliant HTTP communication, typed parameter validation, and robust exception handling for Laravel, Symfony, WordPress, and standard PHP applications.

---

## Installation

Install via Composer:

```bash
composer require bursapay/bursapay-php
```

Requires PHP 7.4 or 8.x with `cURL` and `JSON` extensions.

---

## Quickstart

```php
<?php
require_once __DIR__ . '/vendor/autoload.php';

use BursaPay\BursaPay;

// Initialize client with secret key
$bursapay = new BursaPay(getenv('BURSAPAY_SECRET_KEY'));

// Optional: configure custom timeout or base URL
// $bursapay = new BursaPay('bp_sec_test_...', [
//     'base_url' => 'http://localhost:8000/api/v1',
//     'timeout' => 30
// ]);
```

---

## Payments

### 1. Initialize Checkout Session
```php
$response = $bursapay->payments->initialize([
    'amount' => 15000.00,
    'email' => 'customer@example.com',
    'currency' => 'NGN',
    'reference' => 'ORDER-PHP-2026-001',
    'callback_url' => 'https://myshop.com/checkout/callback',
    'metadata' => [
        'order_id' => 'ORDER-PHP-2026-001',
        'customer_id' => 'CUST-8812',
    ],
]);

// Redirect customer to authorization URL
$checkoutUrl = $response['data']['authorization_url'];
$reference = $response['data']['reference'];
header("Location: " . $checkoutUrl);
exit;
```

### 2. Verify Payment
```php
$reference = 'ORDER-PHP-2026-001';
$verification = $bursapay->payments->verify($reference);

if ($verification['data']['status'] === 'success') {
    $amountPaid = $verification['data']['amount'];
    $channel = $verification['data']['channel'];
    // Fulfill customer order...
    echo "Payment of ₦{$amountPaid} verified via {$channel}!";
}
```

### 3. Charge Saved Card
```php
$charge = $bursapay->payments->chargeSavedCard([
    'customer_reference' => 'CUST-99214',
    'authorization_code' => 'AUTH_8b91fa02c',
    'amount' => 5000.00,
]);
```

---

## Dedicated Virtual Accounts

```php
$virtualAccount = $bursapay->virtualAccounts->create([
    'customer_email' => 'merchant@example.com',
    'bvn' => '12345678901',
    'preferred_bank' => 'Wema Bank',
]);

$nuban = $virtualAccount['data']['account_number'];
$bank = $virtualAccount['data']['bank_name'];
echo "Pay via bank transfer to {$nuban} ({$bank})";
```

---

## Transfers & Disbursements

```php
$transfer = $bursapay->transfers->create([
    'amount' => 25000.00,
    'account_number' => '0123456789',
    'bank_code' => '058', // GTBank
    'narration' => 'Vendor payment',
    'reference' => 'TRF-PHP-9901',
]);
```

---

## Subscriptions & Invoices

```php
// Create Subscription Plan
$plan = $bursapay->subscriptions->createPlan([
    'name' => 'Gold Membership',
    'amount' => 12000.00,
    'interval' => 'monthly',
    'currency' => 'NGN',
]);

// Create Itemized Invoice
$invoice = $bursapay->invoices->create([
    'customer_email' => 'billing@client.com',
    'due_date' => '2026-11-15',
    'items' => [
        ['description' => 'Web Development', 'quantity' => 1, 'unit_price' => 150000.00],
    ],
]);
```

---

## Webhook Signature Verification

```php
<?php
use BursaPay\Resources\Webhooks;
use BursaPay\Exceptions\SignatureVerificationException;

$payload = file_get_contents('php://input');
$signature = $_SERVER['HTTP_X_BURSAPAY_SIGNATURE'] ?? '';
$timestamp = $_SERVER['HTTP_X_BURSAPAY_TIMESTAMP'] ?? '';
$secret = getenv('BURSAPAY_WEBHOOK_SECRET');

try {
    $event = Webhooks::constructEvent($payload, $signature, $timestamp, $secret);

    switch ($event['event']) {
        case 'payment.success':
            $payment = $event['data'];
            // Handle successful payment
            break;
        case 'transfer.success':
            $transfer = $event['data'];
            // Handle transfer completion
            break;
    }

    http_response_code(200);
    echo json_encode(['status' => 'success']);
} catch (SignatureVerificationException $e) {
    http_response_code(400);
    echo json_encode(['error' => 'Invalid signature: ' . $e->getMessage()]);
}
```

---

## Exception Hierarchy

All exceptions in the SDK extend `BursaPay\Exceptions\BursaPayException`:

```php
use BursaPay\Exceptions\AuthenticationException;
use BursaPay\Exceptions\InvalidRequestException;
use BursaPay\Exceptions\RateLimitException;
use BursaPay\Exceptions\BursaPayException;

try {
    $response = $bursapay->payments->initialize([...]);
} catch (AuthenticationException $e) {
    // Invalid secret key or insufficient scope (HTTP 401/403)
} catch (InvalidRequestException $e) {
    // Missing or malformed parameters (HTTP 400/422)
} catch (RateLimitException $e) {
    // Rate limit hit (HTTP 429)
} catch (BursaPayException $e) {
    // Generic API or network failure
}
```
