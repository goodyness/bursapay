# Transparent Fee Structure & Pricing Model

BursaPay operates with absolute transparency. Our pricing structure is engineered to be simple, predictable, and fair for students, institutions, public organizations, event organizers, vendors, and developers.

---

## 📊 Comprehensive Platform Fee Matrix

| Platform Pillar | Product / Service | BursaPay Fee | Underlying Gateway Fee | Who Bears the Fee? |
|---|---|---|---|---|
| **Institutional Mode (V2)** | Academic Session & Departmental Dues | ₦150 – ₦300 flat per transaction | Standard Paystack / NIP rail | Configurable (Student by default) |
| **Public Mode** | Public Payment Links & Association Dues | 1.5% (Capped at ₦2,000) | Included | Configurable (Organizer absorbs or Payer bears) |
| **Public Mode** | Crowdfunding & Donations | 1.5% | Included | Configurable (Organizer absorbs or Donor bears) |
| **Event Ticketing** | Paid Event Tickets | 2.5% + ₦100 per ticket | Included | Configurable (Registrar absorbs or Attendee bears) |
| **Event Ticketing** | Free Tickets & RSVPs | **₦0.00 (100% Free)** | **₦0.00** | Free for everyone |
| **Vendor Marketplace** | Escrow Service (≤ ₦20,000) | **5.0%** | 1.5% + ₦100 (capped at ₦2,000) | Vendor or Client (configurable) |
| **Vendor Marketplace** | Escrow Service (₦20,001 – ₦50,000) | **4.0%** | 1.5% + ₦100 | Vendor or Client |
| **Vendor Marketplace** | Escrow Service (₦50,001 – ₦100,000)| **3.5%** | 1.5% + ₦100 | Vendor or Client |
| **Vendor Marketplace** | Escrow Service (> ₦100,000) | **3.0%** | 1.5% + ₦100 | Vendor or Client |
| **Developer Gateway** | Online Card / USSD / Transfer Charge | 1.5% + ₦100 (Capped at ₦2,000) | Included | Merchant / Developer |
| **Developer Gateway** | Dedicated Virtual Account (NUBAN) | ₦50 flat per inbound deposit | Included | Merchant / Developer |
| **Transfers & Payouts**| Outbound Bank Payout (≤ ₦5,000) | **₦10.00** | Included | Account balance |
| **Transfers & Payouts**| Outbound Bank Payout (₦5,001 – ₦50,000)| **₦25.00** | Included | Account balance |
| **Transfers & Payouts**| Outbound Bank Payout (> ₦50,000) | **₦50.00** | Included | Account balance |

---

## 🧮 How Fees are Calculated (Step-by-Step Examples)

### Example 1: Institutional Departmental Dues
- **Base Fee Set by Department:** ₦10,000.00
- **Platform Processing Fee:** ₦200.00
- **Gateway Charge:** ₦150.00
- **Total Paid by Student:** **₦10,350.00**
- **Net Credited to Department Wallet:** **₦10,000.00** (100% of base amount)

---

### Example 2: Event Ticket (Registrar Bears Fees vs Attendee Bears Fees)

#### Case A: Attendee Bears Fees (Default)
- **Ticket Price Set by Registrar:** ₦20,000.00
- **Ticket Fee (2.5% + ₦100):** ₦600.00
- **Total Paid by Attendee:** **₦20,600.00**
- **Net Credited to Registrar Wallet:** **₦20,000.00**

#### Case B: Registrar Absorbs Fees (`organizer_bears_fees=True`)
- **Ticket Price Set by Registrar:** ₦20,000.00
- **Total Paid by Attendee:** **₦20,000.00** (Exact intended price)
- **Fee Deducted:** ₦600.00
- **Net Credited to Registrar Wallet:** **₦19,400.00**

---

### Example 3: Vendor Marketplace Escrow Booking
- **Client Books Photography Package:** ₦80,000.00
- **Tier Platform Fee (3.5% on ₦50k-₦100k tier):** ₦2,800.00
- **Paystack Processing Fee (1.5% + ₦100):** ₦1,300.00
- **Upfront Escrow Deposit Paid by Client:** ₦80,000.00
- **Final Net Payout Released to Vendor:** **₦75,900.00** (upon project completion and client sign-off)

---

## 🔒 No Hidden Charges Guarantee

- **Zero Maintenance Fees:** No monthly or annual account keeping fees.
- **Zero Inactive Account Penalties:** Keep your account open without minimum balance requirements.
- **Free Sandbox:** 100% free sandbox environment for developers with unlimited synthetic test seeding.
