# Master Step-by-Step Guides & Sample Workflows

This document serves as an end-to-end visual walkthrough and operational cheatsheet for **Institutional Organizations**, **Public Organizations**, and **Event Registrars**.

---

## 📋 Table of Contents

1. [Institutional Organization Workflow: Creating Dues & Making Payments](#1-institutional-organization-workflow)
2. [Public Organization Workflow: Creating Crowdfunds & Making Donations](#2-public-organization-workflow)
3. [Registrar Workflow: Creating Events, Configuring Seating & Ticket Sales](#3-registrar-workflow)
4. [Event Check-in & Security Staff Workflow](#4-event-check-in--security-staff-workflow)

---

## 1. Institutional Organization Workflow

### A. Creating an Academic Payment (V2)

```
[ Step 1: Login ] ──► [ Step 2: Create Payment ] ──► [ Step 3: Set Tier Pricing ] ──► [ Step 4: Add Custom Fields ]
  • Role: ORG           • Title: "2026 NACOSS Dues"     • 100L: ₦12,500                • Matric No (Required)
  • Mode: Institutional • Slug: /pay/nacoss-2026        • 200L-400L: ₦6,500            • Department Dropdown
                                                        • 500L: ₦8,500                 • Level (100L-500L)
```

1. **Navigate to Payment Creator:** `/dashboard/org/v2/payments/create/`.
2. **Input Base Details:**
   - **Title:** `2026/2027 NACOSS Annual Departmental Dues`
   - **Category:** `Departmental Dues`
   - **University:** `University of Lagos`
3. **Configure Category Pricing:**
   - Go to `/dashboard/org/v2/categories/student/`
   - Add Tier: `100 Level (Freshers)` ➔ `₦12,500.00`
   - Add Tier: `Returning (200L–400L)` ➔ `₦6,500.00`
   - Add Tier: `Final Year (500L)` ➔ `₦8,500.00`
4. **Attach Branding:**
   - Upload official department crest and executive signature stamp at `/dashboard/org/v2/branding/`.

---

### B. Student Payer Checkout Flow

```
1. Student lands on: bursapay.com/pay/nacoss-2026
2. Selects Student Tier: [ 100 Level (Freshers) - ₦12,500 ]
3. Fills Personal Info:
   - Full Name: Sarah Ade
   - Email: sarah.ade@unilag.edu.ng
   - Matric Number: 190408012
   - Level: 100L
4. Settle Payment via Paystack / Bank Transfer.
5. Receives stamped anti-tamper PDF receipt with QR code (bursapay.xxxxxxxx-xxxx).
```

---

### C. "Pay-for-Me" (Parent / Sponsor Delegation)

```
1. Student checks: [✔] "Pay for Me"
2. Enters Sponsor: Chief Ade (sponsor.ade@gmail.com)
3. System sends secure UUID link to sponsor: bursapay.com/pay/for-me/<token>/
4. Sponsor reviews breakdown & pays remotely.
5. Student immediately receives cleared official clearance receipt.
```

---

## 2. Public Organization Workflow

### A. Creating a Public Crowdfunding Drive

```
[ Step 1: Switch Mode ] ──► [ Step 2: Create Public Link ] ──► [ Step 3: Enable Crowdfund ] ──► [ Step 4: Share URL ]
  • Mode: Public              • Title: "Oyo Water Fund"          • Target: ₦5,000,000             • bursapay.com/pay/
  • Set Bank Details          • Fixed / Open Amount              • Wall of Fame: [✔]                clean-water-2026
                                                                 • Live Toasts: [✔]
```

1. **Navigate to Public Link Creator:** `/dashboard/org/pay/create/`.
2. **Input Campaign Details:**
   - **Title:** `2026 Clean Water for Rural Oyo Community`
   - **Description:** `Fundraising for 5 solar-powered boreholes.`
   - **Custom Slug:** `clean-water-2026`
3. **Configure Crowdfunding Options:**
   - Check `is_fundraising = True`.
   - Set **Target Amount:** `₦5,000,000.00`.
   - Toggle `show_wall_of_fame = True` and `show_contribution_toasts = True`.

---

### B. Public Donor Flow

```
1. Donor opens: bursapay.com/pay/clean-water-2026
2. Views live goal tracker: [████████████░░░░] 75% (₦3,750,000 raised)
3. Selects contribution amount: ₦25,000
4. Privacy toggle: [✔] "Donate Anonymously"
5. Pays via Card, Apple Pay, or Bank Transfer.
6. Receives official donation tax receipt via email.
```

---

## 3. Registrar Workflow (Event Management)

### A. Creating an Event with Multi-Tier Tickets & Seating

```
[ Step 1: Create Event ] ──► [ Step 2: Configure Tickets ] ──► [ Step 3: Table & Seat Map ] ──► [ Step 4: Lineup Cards ]
  • Title: Lagos Gala 2026     • Early Bird: ₦5,000 (200 qty)    • 20 Tables (10 seats/table)     • Artists, DJs, MCs
  • Venue: Landmark Centre     • VIP: ₦35,000 (100 qty)          • 5-min Anti-Double Lock         • Social share cards
  • Hidden Location: Hall B    • Table for 10: ₦500,000          • Table Naming: Gold/Platinum
```

1. **Create Base Event:** `/events/create/`
   - **Title:** `Lagos Gala & Awards Night 2026`
   - **Public Location:** `Landmark Centre, Victoria Island, Lagos`
   - **Hidden Location:** `Banquet Hall 3 (Revealed only on VIP passes)`
2. **Configure Ticket Pricing:** `/events/configure-tickets/<event_id>/`
   - Add **Regular Pass:** `₦10,000` (1,000 capacity)
   - Add **VIP Lounge Pass:** `₦35,000` (150 capacity)
   - Add **Executive Table (10 Seats):** `₦500,000` (20 capacity)
3. **Configure Interactive Seating Grid:**
   - Enable `enable_seat_numbering = True`.
   - Set tables and seat numbers with automatic **5-minute checkout locks**.
4. **Publish Lineup:** Add guest artists, keynote speakers, and performers at `/events/<event_id>/guest-lineup/`.

---

### B. Attendee Ticket Purchase & Seat Selection

```
1. Attendee visits: bursapay.com/events/BURSAL-XK9P2Z
2. Selects VIP Lounge Tier.
3. Clicks Table 1 -> Seat 4 (System locks Seat 4 for 5 minutes).
4. Fills attendee name & phone number.
5. Pays ₦35,000 via Debit Card.
6. Receives digital ticket pass with high-density QR code (40-char custom ID).
```

---

## 4. Event Check-in & Security Staff Workflow

### A. Delegating Mini-Registrars (Staff Badges)
1. Event owner goes to **Mini-Registrar Management** (`/events/registrar/delegation/`).
2. Inputs staff email and assigns role: `check_in_staff`.
3. Staff member receives a **Digital Staff ID Card** with a security QR code for venue access.

### B. Live Gate Entry QR Scanning
1. Gate attendants open `/events/registrar/scan/<event_id>/` on any mobile browser.
2. Directs phone camera at attendee's ticket QR code.
3. System responds in **< 0.3 seconds**:
   - `✅ VALID: Sarah Ade - VIP Lounge (Table 1, Seat 4)`
   - `❌ ALREADY SCANNED: Scanned 12 mins ago at Gate 2`
   - `❌ INVALID: Counterfeit / Unrecognized ticket reference`

---

### C. Post-Event Cryptographic Audit Report
1. Event concludes; registrar navigates to `/events/registrar/audit/<event_id>/`.
2. BursaPay compiles total gross sales, platform fees, net settlement, and verified attendance count.
3. Generates an **HMAC-SHA256 digitally signed audit certificate** for sponsors and institutional auditors.
