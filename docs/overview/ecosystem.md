# BursaPay Ecosystem Overview

BursaPay is Nigeria's specialized payment infrastructure platform, connecting higher education institutions, public organizations, event organizers, service marketplaces, and fintech developers into a unified payment settlement network.

---

## Key Metrics & Scale

- **₦1 Billion+** in transaction volume processed.
- **50+ Tertiary Institutions** integrated across Nigeria.
- **100,000+ Students and End-Payers** onboarded.
- **99.9% Uptime** across high-traffic academic registration deadlines.
- **Multi-Gateway Redundancy** with automated failover routing between primary and secondary payment rails.

---

## User Roles & Hierarchy

Every account in BursaPay operates under a role-based access control (RBAC) model:

| Role Code | Role Name | Identifier Prefix | Core Responsibilities |
|---|---|---|---|
| `organization` | Organization | `ORG-` | Educational faculties, student unions, and public non-profits collecting levies or dues. |
| `student` | Student / Payer | `STU-` | Academic payers tied to a university matriculation number, faculty, department, and level. |
| `registrar` | Event Registrar | `REG-` | Event organizers creating ticketing campaigns, managing seating, and scanning entry passes. |
| `mini_registrar` | Mini-Registrar | `MNR-` | Delegated event staff (ticket managers, QR scanners, financial oversight) with granular permissions. |
| `vendor` | Marketplace Vendor | `VND-` | Creative and professional freelancers listing services under milestone-based escrow. |
| `client` | Marketplace Client | `CLI-` | Individuals or businesses hiring vendors through the BursaPay service marketplace. |
| `ambassador` | Brand Ambassador | `AMB-` | Campus and community advocates driving adoption and earning referral commissions. |
| `financial_secretary` | Financial Secretary | `FSC-` | University or departmental officers auditing balances, ledger entries, and payment records. |
| `super_admin` | Platform Admin | `ADM-` | System administrators managing KYC approvals, risk alerts, dispute arbitrations, and settlements. |

---

## System Architecture Blueprint

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                             CLIENT TOUCHPOINTS                              │
├───────────────────┬───────────────────┬───────────────────┬─────────────────┤
│ Web Dashboard     │ Guest Checkout    │ Mobile App        │ Developer SDKs  │
│ (Org/Registrar)   │ (Slug URLs)       │ (iOS / Android)   │ (Python / JS)   │
└─────────┬─────────┴─────────┬─────────┴─────────┬─────────┴────────┬────────┘
          │                   │                   │                  │
          ▼                   ▼                   ▼                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                             APPLICATION LAYER                               │
├───────────────────┬───────────────────┬───────────────────┬─────────────────┤
│ Django REST Core  │ Gateway API (v1)  │ Webhook Dispatch  │ SSE Live Stream │
│ Authentication    │ Serializers & ORM │ Dead Letter Queue │ Real-time Feed  │
└─────────┬─────────┴─────────┬─────────┴─────────┬─────────┴────────┬────────┘
          │                   │                   │                  │
          ▼                   ▼                   ▼                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                            FINANCIAL & RISK CORE                            │
├───────────────────┬───────────────────┬───────────────────┬─────────────────┤
│ Double-Entry      │ Escrow Engine     │ 3-Way Bank        │ Risk & Fraud    │
│ Ledger Engine     │ (Vendor/Milestone)│ Reconciliation    │ Anomaly Scorer  │
└─────────┬─────────┴─────────┬─────────┴─────────┬─────────┴────────┬────────┘
          │                   │                   │                  │
          ▼                   ▼                   ▼                  ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                            INFRASTRUCTURE & RAILS                           │
├───────────────────┬───────────────────┬───────────────────┬─────────────────┤
│ Paystack Rail     │ Dedicated NUBANs  │ Supabase / R2     │ Redis / Celery  │
│ Primary Gateway   │ Virtual Accounts  │ File & Doc Storage│ Asynchronous Qs │
└───────────────────┴───────────────────┴───────────────────┴─────────────────┘
```

---

## Platform Operating Modes

Organizations on BursaPay can switch their operating persona based on their business model:

### 1. Institutional Mode
Optimized for universities, polytechnics, colleges of education, student union governments (SUG), faculty associations, and departmental bodies.
- Pre-populated student rosters and matric number validation.
- Academic session and semester dues breakdown.
- Multi-tier student pricing (Freshers, Returning, Direct Entry, Final Year).

### 2. Public Mode
Engineered for non-academic organizations, charities, non-profits, faith-based organizations, and community drives.
- Public hosted links with zero authentication barriers for payers.
- Dynamic custom forms (text, dropdown, file upload).
- Crowdfunding trackers with live progress bars and Wall of Fame recognition.

### 3. Gateway Mode (Developer API)
Programmatic access for software engineers and enterprises wanting to embed payment processing, dedicated virtual accounts, and automated payouts directly into their proprietary applications.
