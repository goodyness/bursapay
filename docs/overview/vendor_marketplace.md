# Vendor Marketplace & Escrow Engine

The BursaPay Vendor Marketplace connects verified service professionals (photographers, videographers, event decorators, web developers, caterers, MCs) with clients under an automated, milestone-backed escrow protection system.

---

## 4-Phase KYC Onboarding

To maintain platform trust and prevent fraudulent listings, vendors undergo a strict 4-phase onboarding verification:

```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│ Phase 1:        │       │ Phase 2:        │       │ Phase 3:        │       │ Phase 4:        │
│ Business Setup  ├──────►│ Social Proof    ├──────►│ KYC Identity    ├──────►│ Bank Payout     │
│ Name & Category │       │ Portfolio Links │       │ NIN & Slip / RC │       │ Account Setup   │
└─────────────────┘       └─────────────────┘       └─────────────────┘       └─────────────────┘
```

1. **Business Setup**: Business name, taxonomy classification, operational location, and bio.
2. **Social Proof**: Verified social handles (Instagram, X, Facebook, or Website) and portfolio media.
3. **Identity / KYC**: 11-digit National Identification Number (NIN), NIN slip document upload, and CAC business registration documents (for incorporated entities).
4. **Payout Configuration**: Automated NUBAN bank account resolution and verification.

---

## Custom Vanity Domains

Every verified vendor receives a dedicated public storefront:
- **Default URL**: `https://bursapay.com/v/<vendor-slug>/services/`
- **Custom Subdomain**: `https://<brand_name>.bursapay.com` (routed via dynamic front-door dispatch middleware).

---

## The Milestone Escrow Engine

BursaPay's escrow engine protects both parties from default:

```
┌──────────────────┐
│  Client Books    │
│  Service Package │
└────────┬─────────┘
         │ Upfront Deposit (e.g. 50%)
         ▼
┌──────────────────┐       Instant Release %      ┌──────────────────┐
│  BursaPay        ├─────────────────────────────►│  Vendor Wallet   │
│  Escrow Vault    │       (Configurable)         │  (Initial Float) │
└────────┬─────────┘                              └──────────────────┘
         │
         │ Vendor Submits Deliverables & Proof
         ▼
┌──────────────────┐
│  Client Review & │ ─── Satisfied ───► Full Escrow Balance Released to Vendor
│  Approval Window │
└────────┬─────────┘
         │
         └─── Dispute Raised ───► Admin Arbitration & Ledger Freeze
```

---

## Tiered Platform Fee Structure

BursaPay applies transparent, volume-discounted fee rates to vendor service bookings:

| Booking Amount Tier | Platform Fee Rate |
|---|---|
| **≤ ₦20,000** | 5.0% |
| **₦20,001 – ₦50,000** | 4.0% |
| **₦50,001 – ₦100,000** | 3.5% |
| **> ₦100,000** | 3.0% |

*Note: Gateway payment processing fees (1.5% + ₦100, capped at ₦2,000) can be configured to be absorbed by the vendor or passed to the client.*

---

## Dispute Arbitration & Reviews

- **Correction Requests**: Clients can request adjustments without freezing escrow.
- **Formal Disputes**: If work is not delivered, a `MarketplaceDispute` locks escrow funds until platform administrators review chat logs and submitted deliverables.
- **Verified Reviews**: Only clients with completed, paid bookings can publish 1-5 star ratings and reviews on vendor profiles.
