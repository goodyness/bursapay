<?php
/**
 * BursaPay Payment Links PHP Example
 */

$secretKey = getenv('BURSAPAY_SECRET_KEY') ?: 'bp_sec_test_DEMO_KEY_HERE';
$baseUrl = getenv('BURSAPAY_BASE_URL') ?: 'https://api.bursapay.com/api/v1';

// 1. Create a dynamic payment link via cURL
$ch = curl_init("{$baseUrl}/payment-links/");
$payload = json_encode([
    'title' => 'Annual Alumni Gala 2026',
    'description' => 'Support the endowment fund and reserve your gala entry pass.',
    'amount' => 25000.00,
    'currency' => 'NGN',
    'custom_slug' => 'alumni-gala-' . time(),
    'redirect_url' => 'https://example.org/thank-you',
]);

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $payload);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    "Authorization: Bearer {$secretKey}",
    "Content-Type: application/json",
]);

$response = curl_exec($ch);
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

$data = json_decode($response, true);

header('Content-Type: application/json');
echo json_encode([
    'status_code' => $httpCode,
    'bursapay_response' => $data,
], JSON_PRETTY_PRINT);
