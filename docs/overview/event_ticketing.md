# Registrar & Event Ticketing Engine

The Registrar system is BursaPay's end-to-end event ticketing, seating reservation, and access control platform designed for conferences, campus festivals, concerts, banquets, and summits.

---

## Event Lifecycle Architecture

```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│ Registrar       │       │ Attendee Buys   │       │ On-Day Check-in │
│ Creates Event & ├──────►│ Ticket & Locks  ├──────►│ QR Code Scan &  │
│ Configures Seats│       │ 5-min Seat      │       │ Mini-Registrars │
└─────────────────┘       └─────────────────┘       └─────────────────┘
```

---

## Key Features & Capabilities

### 1. Multi-Tier Ticket Pricing
Registrars can configure granular ticket tiers with distinct perks and quantity caps:
- **Early Bird**: Limited time/quantity discounted passes.
- **Regular**: General admission passes.
- **VIP & VVIP**: Premium seating with backstage access.
- **Table Packages**: Group bookings (e.g., Table for 10) with custom table naming.

### 2. Interactive Seating Engine & 5-Minute Locking
To prevent double-booking during peak ticket drops:
- Table-based and general seat maps.
- When an attendee clicks a seat, the system places a **5-minute lock** on the seat record.
- If the checkout completes within 5 minutes, the seat status permanently converts to `is_sold=True`.
- If the session expires or is abandoned, the lock automatically releases back into the pool.

### 3. Mini-Registrar Delegation (Staff Badges)
Event organizers can delegate operational roles to venue staff without exposing financial balances or admin access.

| Staff Role | Key Permissions |
|---|---|
| `check_in_staff` | QR Code scanner access only; validate ticket authenticity at gates. |
| `ticket_manager` | Add, edit, or adjust ticket inventory; cannot view financial wallets. |
| `accountant` | View real-time revenue and ticket sales analytics. |
| `custom` | Granular permission selection across 11 distinct operational privileges. |

- **Digital Staff ID Cards**: Every assigned mini-registrar receives a verifiable digital badge with an embedded QR code for entry into organizer areas.
- **Audit Logs**: Every mini-registrar scan, edit, or check-in is recorded in the `MiniRegistrarAuditLog`.

---

## Access Control & Check-in

### 1. Cryptographic QR Check-in
- Each purchased ticket generates a unique 40-character alphanumeric `custom_ticket_id` encoded into a high-density QR code.
- Gate attendants scan the pass using the BursaPay web scanner or mobile app.
- Real-time deduplication prevents counterfeit duplicate ticket entries.

### 2. OTP Ticket Recovery
Attendees who misplace their confirmation emails can retrieve their passes using self-service OTP verification at `/events/resend-ticket/email-lookup/`.

### 3. Cryptographic Event Audit Reports
After the event concludes, BursaPay compiles a comprehensive `EventAuditReport`:
- Total gross revenue, platform fees, net settlement, and verified attendance count.
- Anomaly risk factor calculation and fraud scoring.
- **HMAC-SHA256 Digital Signature**: Guarantees report cannot be tampered with after generation.
