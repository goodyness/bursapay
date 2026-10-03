# BursaPay — Modern Payment Infrastructure for Africa

<div align="center">

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![API Version](https://img.shields.io/badge/API_Version-v1.0.0-green.svg)](docs/api/getting_started.md)
[![Status](https://img.shields.io/badge/Uptime-99.9%25-brightgreen.svg)](https://bursapay.com)
[![Volume Processed](https://img.shields.io/badge/Processed-₦1B%2B-blueviolet.svg)](docs/overview/ecosystem.md)
[![Institutions](https://img.shields.io/badge/Institutions-50%2B-orange.svg)](docs/overview/institutional_mode.md)

**Simplified, high-reliability payment infrastructure for institutions, public organizations, event organizers, service marketplaces, and modern developers.**

[Explore Docs](docs/overview/ecosystem.md) · [API Reference](docs/api/getting_started.md) · [Webhooks Guide](docs/webhooks/overview.md) · [Python SDK](docs/tools/python_sdk.md) · [JS SDK](docs/tools/js_sdk.md) · [Go SDK](docs/tools/go_sdk.md) · [PHP SDK](docs/tools/php_sdk.md) · [Inline Popup](docs/tools/inline_checkout.md) · [CLI](docs/tools/cli.md) · [OpenAPI Spec](specs/openapi.yaml)

</div>

---

## ⚡ What is BursaPay?

**BursaPay** is a comprehensive financial technology platform and developer gateway engineered to streamline payment collection, escrow settlements, event ticketing, and automated reconciliation across Nigeria and Africa.

Having processed over **₦1 Billion+** across **50+ tertiary institutions** and **100,000+ registered students**, BursaPay provides unified infrastructure connecting end-payers, institutions, event organizers, freelance vendors, and enterprise API developers.

```
                                  ┌───────────────────────────────┐
                                  │   BursaPay Core Platform      │
                                  └───────────────┬───────────────┘
                                                  │
        ┌───────────────────┬─────────────────────┼─────────────────────┬───────────────────┐
        ▼                   ▼                     ▼                     ▼                   ▼
┌──────────────┐    ┌──────────────┐      ┌──────────────┐      ┌──────────────┐    ┌──────────────┐
│ Institutional│    │ Public Org   │      │ Registrar &  │      │ Vendor &     │    │ Developer    │
│ Academic V2  │    │ Payment Links│      │ Event Engine │      │ Escrow Mkt   │    │ Gateway API  │
└───────┬──────┘    └───────┬──────┘      └───────┬──────┘      └───────┬──────┘    └───────┬──────┘
        │                   │                     │                     │                   │
        ▼                   ▼                     ▼                     ▼                   ▼
  - Multi-tier fees   - Public links        - Tiered tickets      - 4-Phase KYC       - REST API v1
  - Matric/Dept forms - Fundraising bars    - 5-min seat locking  - Escrow release    - Dedicated NUBAN
  - "Pay-for-Me" flow - Wall of Fame        - Mini-Registrars     - Milestone payouts - Subscriptions
  - Anti-tamper PDF   - Split sessions      - QR Code Check-in    - Dispute handling  - Real-time SSE
```

---

## 🏛️ Platform Capabilities at a Glance

| Pillar | Key Features | Target Audience | Documentation |
|---|---|---|---|
| **Institutional Mode (V2)** | Departmental dues, faculty levies, multi-category pricing (100L–500L, Staylites), dynamic custom fields, "Pay-for-Me" parent workflow, anti-tamper QR receipts (`bursapay.xxxxxxxx-xxxx`). | Universities, Student Unions, Faculties, Departments | [Read Guide](docs/overview/institutional_mode.md) |
| **Public Mode** | Public payment URLs (`bursapay.com/pay/<slug>`), fundraising campaign trackers with target bars, anonymous donor options, public Wall of Fame, installment/split checkout sessions. | NGOs, Alumni Associations, Charities, Businesses | [Read Guide](docs/overview/public_mode.md) |
| **Registrar & Events** | Custom event landing pages, multi-tiered tickets (Regular, VIP, Tables), interactive seat reservation with 5-minute locking, guest lineup cards, digital staff badges with mini-registrar permissions, QR scanning app, cryptographic audit logs. | Concerts, Conferences, Seminars, Event Organizers | [Read Guide](docs/overview/event_ticketing.md) |
| **Vendor Marketplace** | 4-phase onboarding with NIN/KYC verification, custom vanity URLs (`brand.bursapay.com`), multi-tiered escrow protection, instant partial escrow release, milestone chat, dispute arbitration engine. | Photographers, Designers, Caterers, Freelancers | [Read Guide](docs/overview/vendor_marketplace.md) |
| **Developer Gateway** | Full-featured REST API v1, Dedicated Virtual NUBAN Accounts, 2-Step Payment Intents (Authorize & Capture), Subscriptions & Invoicing, Single & Bulk Transfers, Webhook DLQ, Real-time SSE Stream. | Fintechs, SaaS apps, Mobile & Web Developers | [Read API Docs](docs/api/getting_started.md) |
| **Finance & Ledger** | Double-entry chart of accounts, immutable transaction journal entries, automated 3-way reconciliation engine (Gateway, BursaPay Ledger, Bank Statement), suspicious velocity risk scoring. | CFOs, Auditors, Financial Secretaries | [Read Guide](docs/overview/finance_and_reconciliation.md) |

---

## 🚀 Developer Gateway Quickstart

Integrate payment acceptance into your application in less than 5 minutes.

### 1. Initialize a Payment

Send a `POST` request to `/api/v1/payments/initialize/`:

```bash
curl -X POST https://api.bursapay.com/api/v1/payments/initialize/ \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE" \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: 7b31e9a2-4a5f-4a3d-a4e9-9d0a1b2c3d4e" \
  -d '{
    "amount": 25000.00,
    "email": "customer@example.com",
    "currency": "NGN",
    "reference": "ORD-2026-9812",
    "callback_url": "https://yourapp.com/checkout/callback",
    "metadata": {
      "customer_id": "CUST-4412",
      "product": "Pro Annual Plan"
    }
  }'
```

**Response (`201 Created`):**
```json
{
  "success": true,
  "message": "Payment initialized successfully",
  "data": {
    "reference": "ORD-2026-9812",
    "access_code": "0pylnd82pe",
    "authorization_url": "https://checkout.bursapay.com/pay/0pylnd82pe",
    "amount": "25000.00",
    "currency": "NGN"
  }
}
```

### 2. Verify Payment Status

Query payment status on callback or manually:

```bash
curl -X GET "https://api.bursapay.com/api/v1/payments/verify/?reference=ORD-2026-9812" \
  -H "Authorization: Bearer bp_sec_test_DEMO_KEY_HERE"
```

---

## 📦 Official Developer Tooling

| Package / Tool | Language / Runtime | Installation | Description | Documentation |
|---|---|---|---|---|
| **Python SDK** | Python 3.8+ | `pip install bursapay-sdk` | Official client with sync (`BursaPay`) & async (`AsyncBursaPay`) support. | [Python Guide](docs/tools/python_sdk.md) |
| **JavaScript / TS SDK** | Node.js 18+ / Browser | `npm install bursapay-sdk` | TypeScript-first library supporting ESM, CommonJS, and inline popup checkout. | [JS/TS Guide](docs/tools/js_sdk.md) |
| **Go SDK** | Go 1.18+ | `go get github.com/goodyness/bursapay-go-sdk` | Idiomatic Go client with `context.Context` and HMAC signature validation. | [Go Guide](docs/tools/go_sdk.md) |
| **PHP SDK** | PHP 7.4 / 8.x | `composer require bursapay/bursapay-php` | PSR-compliant client for Laravel, Symfony, WordPress, and standard PHP. | [PHP Guide](docs/tools/php_sdk.md) |
| **Inline Web Popup** | Browser JS | `<script src="https://api.bursapay.com/bursapay-embed.js">` | Drop-in modal checkout iframe overlay for web storefronts. | [Inline Guide](docs/tools/inline_checkout.md) |
| **BursaPay CLI** | Node.js (CLI) | `npm install -g bursapay-cli` | Webhook tunneling via SSE, signature verification, sandbox mocking. | [CLI Guide](docs/tools/cli.md) |

### CLI in Action:
```bash
# Listen to real-time webhooks and forward to your local development server
bursapay listen --forward-to http://localhost:8000/api/webhooks/

# Trigger mock webhook events locally
bursapay trigger payment.success --forward-to http://localhost:8000/api/webhooks/
```

---

## 📂 Repository Structure

```
bursapay-public/
├── docs/
│   ├── overview/                     # Ecosystem deep dives & platform pillar guides
│   │   ├── ecosystem.md
│   │   ├── step_by_step_guides.md    # Master step-by-step guides for Orgs, Public & Events
│   │   ├── bursa_ai_and_knowledge_base.md # Customer-facing Bursa AI & Self-Service KB
│   │   ├── whatsapp_bot_commerce.md  # WhatsApp bot: Pay dues, buy tickets & get receipts in chat
│   │   ├── fees_and_pricing.md       # Transparent fee matrix & calculation examples
│   │   ├── roles_and_permissions.md  # Deep breakdown of all 13 platform user roles
│   │   ├── institutional_mode.md     # Academic V2, multi-category dues & Pay-for-Me
│   │   ├── public_mode.md            # Public links, crowdfunding & Wall of Fame
│   │   ├── event_ticketing.md        # Events, 5-min seat lock, Mini-registrars & QR passes
│   │   ├── vendor_marketplace.md     # 4-phase KYC, escrow engine & milestone payouts
│   │   ├── mobile_ecosystem.md       # Mobile apps, offline passes & push alerts
│   │   ├── whatsapp_and_notifications.md # Automated WhatsApp receipts & reminders
│   │   ├── support_and_verification.md # Public receipt verification & dispute SLAs
│   │   ├── ambassador_program.md     # Referral rewards & campus ambassador program
│   │   └── finance_and_reconciliation.md # Double-entry ledger & 3-way reconciliation
│   ├── api/                          # REST API v1 complete specifications
│   │   ├── all_endpoints_reference.md# Master catalog of all 45+ gateway endpoints
│   │   ├── getting_started.md
│   │   ├── security_and_scopes.md
│   │   ├── sandbox_testing.md        # Test cards, simulated NUBANs & Seeding API
│   │   ├── payments.md               # Intialize, Verify, Saved Cards, Intents, Bulk, Schedule
│   │   ├── virtual_accounts.md       # Dedicated NUBAN accounts
│   │   ├── payment_links.md          # Hosted links, bulk links, QR codes & A/B testing
│   │   ├── subscriptions.md          # Recurring plans, pause, resume & cancel
│   │   ├── invoices.md               # Itemized invoicing & PDF receipts
│   │   ├── transfers_payouts.md      # Single & bulk bank payouts
│   │   ├── refunds_and_disputes.md   # Chargeback defense & automated refunds
│   │   ├── customers.md              # Customer CRM, bulk import, saved cards & history
│   │   └── wallet_and_settlements.md # Balances, ledger, analytics & 3-way reconciliation
│   ├── webhooks/                     # Webhook integration & security
│   │   ├── overview.md               # Architecture, retries, DLQ & delivery logs
│   │   ├── signatures.md             # HMAC SHA-256 verification in all languages
│   │   ├── event_catalog.md          # Exhaustive list of 25+ event types
│   │   └── sse_realtime_stream.md    # Real-time event stream via Server-Sent Events
│   └── tools/                        # SDKs, CLI and Popup documentation
│       ├── python_sdk.md             # Python sync & async SDK guide
│       ├── js_sdk.md                 # Node.js & TypeScript SDK guide
│       ├── go_sdk.md                 # Go SDK integration guide
│       ├── php_sdk.md                # PHP SDK integration guide
│       ├── inline_checkout.md        # Web modal popup checkout integration
│       └── cli.md                    # Developer CLI guide
├── specs/
│   ├── openapi.yaml                  # OpenAPI 3.0.3 YAML Schema (45+ endpoints)
│   ├── openapi.json                  # OpenAPI 3.0.3 JSON Schema
│   └── postman_collection.json       # 1-Click Postman Collection
├── examples/                         # Standalone runnable code starters
│   ├── node-express/                 # Express.js checkout + webhook handler
│   ├── python-fastapi/               # FastAPI payment intent + signature validation
│   ├── html-checkout/                # Pure client-side checkout popup demo
│   ├── vendor-payouts-go/            # Go bulk payouts & webhook listener
│   ├── payment-links-php/            # PHP dynamic payment links starter
│   ├── saas-subscriptions-python/    # Python recurring SaaS subscriptions demo
│   └── webhook-server-node/          # Node.js webhook receiver with HMAC validation
├── CONTRIBUTING.md
├── SECURITY.md
└── LICENSE
```

---

## 🔐 Security & Compliance

BursaPay is built with defense-in-depth:
- **Hashed API Credentials**: Secret keys are stored as SHA-256 hashes and displayed only once upon creation.
- **HMAC SHA-256 Webhooks**: Every webhook dispatch carries an `X-BursaPay-Signature` header signed with your account secret.
- **Granular Scopes**: Restrict API keys to specific permissions (`payments:read`, `transfers:write`, `refunds:write`, etc.).
- **IP Allowlisting**: Optional CIDR/IP restriction on server-to-server API calls.
- **Idempotency**: Prevent accidental double-charges on network retries using `Idempotency-Key`.

For security disclosures, please consult our [Security Policy](SECURITY.md).

---

## 📄 License

This public documentation and developer toolkit are licensed under the [MIT License](LICENSE).
© 2026 BursaPay Technologies Inc. All rights reserved.
