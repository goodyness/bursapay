# Registrar & Event Ticketing Engine Guide

The **Registrar System** is BursaPay's end-to-end event management, ticket sales, interactive seating reservation, and access control platform designed for music concerts, tech conferences, academic summits, dinners, and gala nights.

---

## 🎪 PART 1: Registrar Onboarding & Account Setup

1. **Register as a Registrar:** Navigate to `https://bursapay.com/registrar/register/`.
2. **Account Verification:** New registrar accounts start with `is_verified=False` for trust safety. Platform admins verify identity and event track record.
3. **Configure Payout Bank Details:** Add your settlement bank account at `/events/registrar/account-details/`.
4. **Access Dashboard:** Log into `/events/registrar/dashboard/` to view live revenue, ticket velocity, and active events.

---

## 🎟️ PART 2: How to Create an Event (Step-by-Step)

```
┌─────────────────────────────────────────────────────────────┐
│ 🎪 CREATE NEW EVENT                                         │
│                                                             │
│ Event Title: Lagos Tech Summit & Music Festival 2026        │
│ Category: [ Tech / Conference / Concert      ▼ ]            │
│ Public Location: Landmark Event Centre, Victoria Island     │
│ Hidden Location: Hall B (Revealed only to VIP ticket holders)│
│ Start Date & Time: 2026-11-20 at 18:00                      │
│ Enable Seating & Tables: [✔] Yes                            │
│ Escrow Protection: [✔] Yes (Holds ₦500,000 artist deposit)  │
│ Require Registration Approval: [ ] No                       │
│ Send 24-Hour Email Reminders: [✔] Yes                       │
└─────────────────────────────────────────────────────────────┘
```

### Step 1: Base Event Setup
1. Go to **Events > Create Event** (`GET/POST /events/create/`).
2. Input the **Title**, **Description**, **Category**, and **Date/Time**.
3. **Public vs Hidden Location:**
   - **Public Location:** Displayed on event cards and social links (e.g., *Landmark Beach, VI, Lagos*).
   - **Hidden Location:** Private suite/room number displayed only on the verified ticket pass after purchase (ideal for VIP dinners or exclusive executive roundtables).
4. Click **Save Event**. BursaPay auto-generates a unique event custom ID (e.g. `BURSAL-XK9P2Z`) and downloadable event promotional QR code.

---

### Step 2: Configure Multi-Tier Ticket Pricing
Go to **Configure Tickets** (`/events/configure-tickets/<event_id>/`):

| Ticket Tier | Price | Quantity Cap | Perks & Inclusions |
|---|---|---|---|
| **Early Bird** | ₦5,000 | 200 passes | General Admission access (Limited promo) |
| **Regular** | ₦10,000 | 1,000 passes | Standard Entry + Event Swag Bag |
| **VIP Pass** | ₦35,000 | 150 passes | VIP Lounge, Fast-Track Entry & Complimentary Drinks |
| **Table for 10** | ₦500,000 | 20 tables | Front Stage Table, Premium Bottle Service & Buffet |

---

### Step 3: Interactive Seating Engine & 5-Minute Real-Time Locking

For seated banquets, award nights, and theater conferences:

```
┌─────────────────────────────────────────────────────────────┐
│ 💺 INTERACTIVE TABLE & SEAT GRID                           │
│                                                             │
│ [Table 1: Gold Tier]        [Table 2: Platinum Tier]        │
│ (1) (2) (3) (4) (5)         (1) (2) (3) (4) (5)             │
│ (6) (7) [8] (9) (10)        (6) (7) (8) (9) (10)            │
│                                                             │
│ 🟢 Seat Available   🔒 [8] Locked (5m hold)   🔴 Sold       │
└─────────────────────────────────────────────────────────────┘
```

1. **Enable Seating:** Check `enable_seat_numbering=True` and select `table_based` or `general`.
2. **Table Configuration:** Set total number of tables (e.g., 25), seats per table (e.g., 10), and table naming schemes (e.g., *Gold Table 1*, *Diamond Table A*).
3. **The 5-Minute Anti-Double Booking Lock:**
   - When an attendee clicks a specific seat, BursaPay places an atomic lock on `Seat.locked_at`.
   - The seat is reserved for **5 minutes** while the attendee completes checkout.
   - If the payment succeeds, the seat status converts to `is_sold=True`.
   - If the payment is abandoned or times out, the lock automatically releases back to the public pool.

---

## 🌟 PART 3: Amazing Event Tools Built into BursaPay

### 1. Performer & Guest Lineup Cards
Add stage talent and keynote speakers to your public event page:
- Route: `/events/<event_id>/guest-lineup/`
- Supported Roles: **Artiste, Keynote Speaker, DJ, Comedian, MC, Host, Panelist, Influencer**.
- Features: High-resolution profile photo, stage name, bio, social media handles, and auto-generated shareable lineup cards for Instagram/Twitter.

### 2. Custom Registration Form Builder
Collect custom information before ticket issuance:
- Route: `/events/<event_id>/form-questions/`
- Field Types: **Short text, Multi-line bio, Dropdown select, Yes/No, Document/Photo file upload** (e.g. proof of student ID or pitch deck).

### 3. Mini-Registrars (Staff Delegation with Digital ID Badges)
Delegate on-the-ground responsibilities to team members without sharing your master login credentials or wallet access:
- Assign roles: `check_in_staff`, `ticket_manager`, `accountant`, or `custom`.
- **Verifiable Digital Staff Passes:** Each staff member gets a verifiable digital ID badge with an embedded QR code for entry into production areas.
- **Audit Logging:** Every scan, edit, or override is tracked in the `MiniRegistrarAuditLog`.

### 4. Attendee Networking & "Meet & Connect"
For conferences and summits, attendees can opt into the event networking hub:
- Create professional profiles (Job Title, Company, Bio, Interests).
- Send connection requests to fellow attendees.
- In-platform 1-on-1 networking chat.

### 5. High-Speed Gate Entry QR Scanner
- Attendants open `/events/registrar/scan/<event_id>/` on any smartphone camera or tablet.
- Scans high-density ticket QR codes in under **0.3 seconds**.
- Real-time deduplication alerts prevent counterfeit or reused tickets.

### 6. OTP Self-Service Ticket Recovery
Attendees who lost their ticket email can visit `/events/resend-ticket/email-lookup/`, enter their email, verify with a 6-digit OTP, and immediately re-download their passes.

### 7. Cryptographic Event Audit Reports
After the event, generate a tamper-proof financial and attendance audit sheet:
- Total gross sales, platform fee deductions, net payout, and verified gate check-in counts.
- **HMAC-SHA256 Cryptographic Signature:** Guarantees report integrity for sponsors and institutional auditors.

---

## 🛒 PART 4: How Attendees Purchase Tickets

```
1. Attendee visits Event Page ──► 2. Selects Ticket Tier & Seat Map ──► 3. Fills Custom Questions
                                                                                  │
                                                                                  ▼
6. Emailed QR Ticket Pass ◄── 5. Instant Payment Webhook ◄── 4. Card / USSD / Bank Transfer
```

1. **Visit Event Page:** Attendee opens `https://bursapay.com/events/<custom_id>/`.
2. **Select Tier & Seats:** Selects ticket tier (e.g. *VIP Pass*) and clicks desired seats on the interactive map (initiating the 5-minute lock).
3. **Fill Form:** Enters attendee name, email, phone number, and answers required organizer questions.
4. **Checkout:** Pays via Debit Card, Bank Transfer, or USSD.
5. **Receive Pass:** Instant ticket pass generated with a 40-character unique ID (`custom_ticket_id`) and high-density QR code delivered to their email and mobile screen.
