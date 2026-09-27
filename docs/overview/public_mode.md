# Public Mode & Community Fundraising

BursaPay's Public Mode empowers non-academic entities—such as non-governmental organizations (NGOs), alumni associations, faith communities, sports clubs, and business groups—to collect payments and run crowdfunding campaigns with zero technical setup.

---

## Capabilities Overview

### 1. Hosted Public Payment Links
Organizations can create custom hosted payment pages accessible at:
```
https://bursapay.com/pay/<custom-slug>/
```
- **Custom Vanity Slugs**: Descriptive, branded URLs (e.g., `/pay/alumni-annual-reunion-2026`).
- **Flexible Amounts**: Choose between fixed pricing (e.g., ₦10,000 association membership) or custom open amounts (e.g., donations).
- **Custom Branding**: Upload organization banners, logos, and mission statements.

---

## Crowdfunding & Social Proof Features

When `is_fundraising=True`, the public payment link transforms into an interactive crowdfunding portal:

```
┌─────────────────────────────────────────────────────────────┐
│ 🎯 ALUMNI BUILDING FUNDRAISER 2026                          │
│                                                             │
│ ₦4,250,000 raised of ₦5,000,000 goal                       │
│ [██████████████████████████████████░░░░] 85%                │
│                                                             │
│ 👥 Wall of Fame:                                            │
│ - Ade*** S. contributed ₦100,000 (2 mins ago)               │
│ - Anonymous contributed ₦50,000 (15 mins ago)               │
│ - Chi*** O. contributed ₦25,000 (1 hour ago)                │
└─────────────────────────────────────────────────────────────┘
```

### Key Crowdfunding Features:
1. **Target Amount & Live Progress Bar**: Real-time percentage progress calculated from verified contributions.
2. **Wall of Fame**: Displays recent contributors with privacy-preserving masking (e.g., `Ade*** S.`) to encourage social proof.
3. **Anonymous Donation Mode**: Donors can opt out of public display with a single toggle.
4. **Live Contribution Toasts**: Real-time on-page popup alerts celebrating new donations.

---

## Split Payment Sessions (Installment Support)

For high-ticket contributions or event packages, BursaPay supports **Split Payment Sessions**:
- Payers can break down large payments into manageable installments.
- BursaPay issues a resumable session URL: `https://bursapay.com/p/resume/<session_id>/`.
- The session tracks `total_amount`, `amount_paid`, and state transitions:
  `pending` ➔ `partially_paid` ➔ `completed`.

---

## Automated Split Payouts & Subaccounts

Public organizations with their own Paystack subaccounts can configure **Instant Split Settlements**:
- Route a percentage of incoming revenue directly to partner or department accounts.
- Retain platform accounting records while automating co-organizer payouts.
- Immutable audit logs captured under `WithdrawalRequest` with `instant_payout` status.
