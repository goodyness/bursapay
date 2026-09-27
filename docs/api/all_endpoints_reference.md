# Complete Gateway API Endpoints Reference (Master Catalog)

This document provides an exhaustive, endpoint-by-endpoint reference for all routes mounted under `/api/v1/` in the BursaPay Developer Gateway.

---

## Endpoint Quick Reference Index

| Category | Method | Endpoint | Required Scope | Description |
|---|---|---|---|---|
| **Health & IP** | `GET` | `/health/` | None | Operational status check. |
| | `GET` | `/ip-check/` | None | Returns caller IP and allowlist status. |
| **Payments** | `POST` | `/payments/initialize/` | `payments:write` | Initialize a card/transfer checkout session. |
| | `GET` | `/payments/verify/` | `payments:read` | Real-time payment verification by reference. |
| | `POST` | `/payments/charge/` | `payments:write` | Direct bank/card charge initiation. |
| | `POST` | `/payments/charge-saved-card/` | `payments:write` | Tokenized 1-click recurring card charge. |
| | `POST` | `/payments/bulk/` | `payments:write` | Initialize batch of multiple payments. |
| | `GET` | `/payments/bulk/<batch_ref>/` | `payments:read` | Query status of bulk payment batch. |
| | `GET` | `/payments/callback/` | None | Paystack payment return callback handler. |
| | `GET` | `/payments/` | `payments:read` | List all payments with pagination & filters. |
| | `GET` | `/payments/<reference>/` | `payments:read` | Retrieve complete details of a single payment. |
| | `POST` | `/payments/<reference>/fulfillment/` | `payments:write` | Mark physical or digital goods fulfilled. |
| | `DELETE`| `/payments/<reference>/schedule/` | `payments:write` | Cancel a scheduled recurring payment. |
| **Payment Intents** | `POST` | `/payment-intents/` | `payments:write` | Create a 2-step auth & hold payment intent. |
| | `GET` | `/payment-intents/` | `payments:read` | List payment intents. |
| | `GET` | `/payment-intents/<intent_ref>/` | `payments:read` | Retrieve payment intent detail. |
| | `POST` | `/payment-intents/<intent_ref>/capture/` | `payments:write` | Capture held authorization funds. |
| | `POST` | `/payment-intents/<intent_ref>/cancel/` | `payments:write` | Void/cancel payment authorization. |
| **Customers** | `GET` | `/customers/` | `payments:read` | List merchant customers with pagination. |
| | `POST` | `/customers/` | `payments:write` | Create a new customer profile. |
| | `POST` | `/customers/import/` | `payments:write` | Bulk import customer records. |
| | `GET` | `/customers/<customer_ref>/` | `payments:read` | Retrieve customer profile. |
| | `PUT/PATCH`| `/customers/<customer_ref>/` | `payments:write` | Update customer metadata. |
| | `GET` | `/customers/<customer_ref>/payments/` | `payments:read` | Transaction history for specific customer. |
| | `GET` | `/customers/<customer_ref>/saved-cards/` | `payments:read` | Reusable tokenized card authorizations. |
| | `GET` | `/customers/<customer_ref>/virtual-accounts/` | `payments:read` | Virtual NUBAN accounts for customer. |
| **Disputes** | `GET` | `/disputes/` | `payments:read` | List active customer chargeback disputes. |
| | `GET` | `/disputes/<reference>/` | `payments:read` | Retrieve dispute details & evidence timeline. |
| | `POST` | `/disputes/<reference>/` | `payments:write` | Submit defense explanation and evidence files. |
| **Refunds** | `GET` | `/refunds/` | `payments:read` | List all processed refunds. |
| | `POST` | `/refunds/` | `refunds:write` | Issue full or partial payment refund. |
| | `GET` | `/refunds/<refund_ref>/` | `payments:read` | Retrieve refund settlement status. |
| **Wallet & Ledger** | `GET` | `/wallet/balance` | `transfers:read` | Available, pending, and escrow balances. |
| | `GET` | `/wallet/ledger/` | `transfers:read` | Full double-entry journal transaction trail. |
| | `GET` | `/wallet/settlements/` | `transfers:read` | List automated bank payout settlements. |
| | `GET` | `/wallet/settlements/<reference>/` | `transfers:read` | View specific settlement batch details. |
| **Webhooks & SSE** | `GET` | `/events/stream/` | `payments:read` | Server-Sent Events (SSE) live stream. |
| | `GET` | `/webhooks/events/` | None | List of all supported webhook event types. |
| | `POST` | `/webhooks/test/` | `webhooks:write` | Trigger mock test webhook dispatch. |
| | `GET` | `/webhooks/` | `webhooks:write` | List configured webhook endpoints. |
| | `POST` | `/webhooks/` | `webhooks:write` | Register new webhook endpoint. |
| | `GET` | `/webhooks/<id>/` | `webhooks:write` | Get webhook endpoint details & secret. |
| | `PUT/PATCH`| `/webhooks/<id>/` | `webhooks:write` | Update URL or subscribed event list. |
| | `DELETE`| `/webhooks/<id>/` | `webhooks:write` | Delete webhook endpoint. |
| | `GET` | `/webhooks/<id>/logs/` | `webhooks:write` | View delivery history and status codes. |
| | `GET` | `/webhooks/logs/<log_id>/` | `webhooks:write` | View request/response payload of attempt. |
| | `POST` | `/webhooks/logs/<log_id>/retry/`| `webhooks:write` | Manually retry failed webhook delivery. |
| | `GET` | `/webhooks/dead-letters/` | `webhooks:write` | List Dead Letter Queue (failed) events. |
| | `POST` | `/webhooks/dead-letters/<pk>/replay/` | `webhooks:write` | Replay dead letter event to active webhook. |
| **Transfers** | `POST` | `/transfers/` | `transfers:write` | Initiate single interbank transfer. |
| | `POST` | `/transfers/bulk/` | `transfers:write` | Initiate batch bulk bank disbursements. |
| | `GET` | `/transfers/` | `transfers:read` | List outbound transfers. |
| | `GET` | `/transfers/<reference>/` | `transfers:read` | Retrieve transfer status. |
| **Payment Links** | `GET` | `/payment-links/` | `payments:read` | List payment links. |
| | `POST` | `/payment-links/` | `payments:write` | Create hosted payment link. |
| | `POST` | `/payment-links/bulk/` | `payments:write` | Create batch of multiple payment links. |
| | `POST` | `/payment-links/ab-test/` | `payments:write` | Setup A/B conversion experiment. |
| | `GET` | `/payment-links/<code>/analytics/`| `payments:read` | Views, conversions, and revenue analytics. |
| | `GET` | `/payment-links/<code>/qr/` | `payments:read` | Generate downloadable QR image. |
| | `GET` | `/payment-links/<code>/` | `payments:read` | Retrieve payment link metadata. |
| | `PUT/PATCH`| `/payment-links/<code>/` | `payments:write` | Update title, price, or expiry. |
| | `DELETE`| `/payment-links/<code>/` | `payments:write` | Deactivate payment link. |
| **Virtual Accounts**| `POST` | `/virtual-accounts/` | `virtual_accounts:write` | Provision dedicated NUBAN bank account. |
| | `GET` | `/virtual-accounts/` | `payments:read` | List merchant virtual accounts. |
| | `GET` | `/virtual-accounts/<pk>/` | `payments:read` | Retrieve virtual account details. |
| **Subscriptions** | `POST` | `/subscriptions/plans/` | `subscriptions:write` | Create recurring billing plan. |
| | `GET` | `/subscriptions/plans/` | `subscriptions:write` | List billing plans. |
| | `POST` | `/subscriptions/` | `subscriptions:write` | Subscribe customer to plan. |
| | `GET` | `/subscriptions/` | `subscriptions:write` | List active customer subscriptions. |
| | `GET` | `/subscriptions/<pk>/` | `subscriptions:write` | Retrieve subscription detail. |
| | `POST` | `/subscriptions/<pk>/pause/` | `subscriptions:write` | Pause automated recurring charges. |
| | `POST` | `/subscriptions/<pk>/resume/` | `subscriptions:write` | Resume active billing cycle. |
| | `POST` | `/subscriptions/<pk>/cancel/` | `subscriptions:write` | Permanently terminate subscription. |
| **Invoices** | `POST` | `/invoices/` | `payments:write` | Create itemized digital invoice. |
| | `GET` | `/invoices/` | `payments:read` | List merchant invoices. |
| | `GET` | `/invoices/<reference>/` | `payments:read` | Retrieve invoice details & PDF URL. |
| **Analytics** | `GET` | `/analytics/` | `payments:read` | High-level metrics, volume & conversion rates. |
| **Reconciliation** | `GET` | `/reconciliation/` | `transfers:read` | Ledger audit & discrepancy breakdown. |
| | `GET` | `/reconciliation/export/` | `transfers:read` | Download reconciliation CSV report. |
| **Sandbox Tools** | `POST` | `/sandbox/seed/` | None (Test Mode) | Seed mock customers, cards & payments. |
| **API Schema** | `GET` | `/schema/` | None | Raw OpenAPI schema. |
| | `GET` | `/docs/` | None | Interactive Swagger UI documentation. |
| | `GET` | `/redoc/` | None | Interactive ReDoc documentation. |
