# Webhook Event Catalog

Below is the complete catalog of event types dispatched by the BursaPay platform with their corresponding payload schemas.

---

## 1. Payments & Checkout Events

### `payment.success`
Triggered when a charge (card, USSD, transfer, QR) is successfully settled.

```json
{
  "event": "payment.success",
  "id": "evt_01J8X9Q2",
  "timestamp": "2026-09-27T02:15:00Z",
  "data": {
    "reference": "BP-ORD-90214",
    "amount": "15000.00",
    "currency": "NGN",
    "status": "success",
    "channel": "card",
    "gateway_fee": "225.00",
    "net_amount": "14775.00",
    "customer": {
      "email": "sarah.ade@example.com",
      "customer_code": "CUST-99214"
    },
    "metadata": {
      "cart_id": "CART-8812"
    }
  }
}
```

### `payment.failed`
Triggered when a charge attempt fails due to insufficient funds, card rejection, or user abandonment.

---

## 2. Virtual Account Events

### `virtual_account.credited`
Triggered when a customer completes an inbound interbank transfer to their assigned dedicated NUBAN.

```json
{
  "event": "virtual_account.credited",
  "id": "evt_01J8X9Q3",
  "timestamp": "2026-09-27T02:15:22Z",
  "data": {
    "account_number": "0123456789",
    "bank_name": "Wema Bank",
    "amount": "50000.00",
    "currency": "NGN",
    "fee": "150.00",
    "net_amount": "49850.00",
    "payer_name": "JOHN DOE",
    "customer_reference": "CUST-99214",
    "transaction_reference": "VA-DEP-883104"
  }
}
```

---

## 3. Transfers & Payout Events

### `transfer.success`
Triggered when an outbound bank payout to a recipient account is confirmed by the interbank switch.

```json
{
  "event": "transfer.success",
  "id": "evt_01J8X9Q4",
  "timestamp": "2026-09-27T02:16:30Z",
  "data": {
    "reference": "TRF-2026-90214",
    "amount": "75000.00",
    "currency": "NGN",
    "bank_name": "Guaranty Trust Bank",
    "account_number": "0123456789",
    "recipient_name": "TUNDE BAKARE",
    "status": "success"
  }
}
```

### `transfer.failed` / `transfer.reversed`
Triggered if a transfer is rejected by the destination financial institution.

---

## 4. Subscriptions & Invoicing Events

| Event Code | Trigger Condition |
|---|---|
| `subscription.created` | New recurring subscription registered. |
| `subscription.charged` | Automated renewal billing succeeded. |
| `subscription.payment_failed` | Recurring card charge attempt declined. |
| `subscription.cancelled` | Subscription deactivated by customer or merchant. |
| `invoice.created` | New digital invoice generated. |
| `invoice.paid` | Customer settled invoice in full. |

---

## 5. Refunds & Disputes Events

| Event Code | Trigger Condition |
|---|---|
| `refund.completed` | Refund credited back to customer's account/card. |
| `dispute.created` | Cardholder filed an official chargeback dispute. |
| `dispute.resolved` | Bank resolved dispute in merchant or customer favor. |
