# Public Mode & Community Fundraising User Guide

**Public Mode** is designed for non-academic organizations—such as Non-Governmental Organizations (NGOs), Alumni Associations, Charities, Religious Bodies, Sports Clubs, and Event Drives—that need to collect payments or run fundraising campaigns from the general public without student login barriers.

---

## 🌐 PART 1: Public Organization Setup

### Step 1: Switch to Public Mode
1. Register at `https://bursapay.com/register/` or switch your existing account at `/dashboard/org/switch-mode/` to **"Public Mode"**.
2. Complete your Public Profile (`/dashboard/org/profile/`) with:
   - **Organization Name** (e.g., *Green Earth Nigeria Foundation*).
   - **CAC / Government Registration Number** (if registered).
   - **Contact Information & Official Website**.
   - **Paystack Subaccount Code** (Optional: for instant split payouts).

### Step 2: Configure Bank Details
Link your commercial bank account under `/dashboard/bank-setup/` for automated weekly or on-demand balance withdrawals.

---

## 🛠️ PART 2: How to Create a Public Payment Link

Public organizations can create fixed-price payment links (e.g. membership dues, conference registration) or open crowdfunding drives.

```
┌─────────────────────────────────────────────────────────────┐
│ 🚀 CREATE PUBLIC PAYMENT LINK                               │
│                                                             │
│ Title: 2026 Community Clean Water Crowdfund                 │
│ Description: Help us drill 5 solar-powered boreholes in Oyo │
│ Amount: [ ₦0.00 (Flexible / Open Donation)   ]              │
│ Enable Fundraising Mode: [✔] Yes                            │
│ Target Goal: [ ₦5,000,000.00                 ]              │
│ Show Wall of Fame: [✔] Yes                                  │
│ Live Contribution Toasts: [✔] Yes                           │
│ Organizer Absorbs Fees: [✔] Yes                             │
│ Custom Slug: bursapay.com/pay/clean-water-oyo-2026          │
└─────────────────────────────────────────────────────────────┘
```

### Step 1: Base Configuration
1. Go to **Dashboard > Public Payments > Create** (`/dashboard/org/pay/create/`).
2. Input the **Title**, **Description**, and **Custom URL Slug**.
3. Choose Pricing Model:
   - **Fixed Price:** Enter an exact amount (e.g., `₦10,000.00`).
   - **Open / Crowdfunding Amount:** Check `is_fundraising=True` and set a minimum contribution (e.g., `₦500.00`).

### Step 2: Enable Social Proof & Crowdfunding Tools
- **Target Goal & Progress Bar (`target_amount`):** Automatically computes and displays the real-time completion percentage.
- **Wall of Fame (`show_wall_of_fame`):** Displays recent donor names on the page.
- **Live Donation Toasts (`show_contribution_toasts`):** Shows animated real-time popups when new contributions land.
- **Fee Absorption (`organizer_bears_fees`):** When enabled, donors pay the exact intended sum, and fees are deducted from the organization wallet balance.

### Step 3: Add Custom Data Fields
Attach custom form questions to your public payment link:
- **T-shirt Size** (Dropdown: S, M, L, XL, XXL)
- **Home City / Chapter** (Text)
- **Special Message / Dedication** (Textarea)

---

## 💳 PART 3: How Payers Complete a Public Payment

```
1. Payer lands on URL ──► 2. Enters Donation Amount ──► 3. Toggles Anonymous / Wall of Fame
                                                                      │
                                                                      ▼
6. Emailed PDF Receipt ◄── 5. Instant Payment Webhook ◄── 4. Card / USSD / Bank Transfer
```

### Step-by-Step Payer Flow:
1. **Visit Link:** The payer visits `bursapay.com/pay/clean-water-oyo-2026`.
2. **View Campaign:** Sees the mission description, progress bar (*₦3,250,000 of ₦5,000,000 raised*), and recent contributors.
3. **Choose Amount:** Selects a suggested tier (₦5,000, ₦10,000, ₦50,000) or types a custom sum.
4. **Privacy Options:** Can check **"Donate Anonymously"** (masks name on the public Wall of Fame as `Anonymous`).
5. **Settle Payment:** Pays seamlessly via Debit Card, Pay with Transfer, USSD, or Apple Pay.
6. **Confirmation & Receipt:** Receives instant on-screen success confirmation and an official PDF donation tax receipt via email.

---

## 🔄 PART 4: Split Payment Sessions (Installment Payments)

For high-ticket community dues or building fund pledges:
- Payers can activate **Installment Checkout**.
- BursaPay issues a resumable session link: `https://bursapay.com/p/resume/<session_id>/`.
- The session tracks payments progressively until the total target is satisfied.
