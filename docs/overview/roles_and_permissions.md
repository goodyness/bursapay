# Comprehensive Guide to User Roles & Permissions in BursaPay

BursaPay is built around a comprehensive Role-Based Access Control (RBAC) model. Every user on the platform has a distinct `role` that dictates their dashboard access, financial permissions, operational capabilities, and security boundaries.

---

## Master Role Matrix

| # | Role Identifier | Display Name | Unique ID Prefix | Target User / Persona | Core Capabilities |
|---|---|---|---|---|---|
| **1** | `super_admin` | Super Admin | `BURSA-ADM-` | BursaPay Platform Owners & Core Ops | Full root system access, KYC approvals, dispute arbitration, financial ledger oversight, risk overrides, platform fee adjustments. |
| **2** | `organization` | Organization | `BURSA-ORG-` | Universities, SUG, Faculties, NGOs, Non-Profits | Mode switching (Institutional / Public / Gateway), multi-tier student dues, custom payment forms, "Pay-for-Me" workflows, branded receipts. |
| **3** | `student` | Student / Payer | `BURSA-STU-` | Undergraduate & Postgraduate Students | Profile tied to Matric No, Faculty, Dept, Level; instant payment history, digital fee receipts, examination clearance verification. |
| **4** | `registrar` | Event Registrar | `BURSA-REG-` | Event Organizers & Promoters | Event creation, multi-tier ticketing (VIP, Regular, Tables), 5-min interactive seat locking, guest lineups, revenue payouts. |
| **5** | `mini_registrar`| Mini-Registrar | `BURSA-MNR-` | Venue Staff & Gate Attendants | Delegated event roles (Check-in scanner, Ticket Manager, Accountant) with digital staff ID cards and time-bounded permissions. |
| **6** | `mini_org` | Mini-Org Staff | `BURSA-MNO-` | Departmental Staff & Faculty Clerks | Delegated departmental clearance verification, payment roster lookup, and receipt validation under an Organization. |
| **7** | `vendor` | Service Vendor | `BURSA-VND-` | Creative Freelancers & Agencies | 4-phase KYC onboarding (NIN/RC), service packages, vanity domain (`brand.bursapay.com`), milestone escrow payouts, portfolio. |
| **8** | `client` | Service Client | `BURSA-CLI-` | Event Hosts & Individuals Hiring Vendors | Service booking, upfront escrow deposits, in-platform project messaging, milestone deliverables review, dispute escalation. |
| **9** | `developer` | API Developer | `BURSA-DEV-` | Software Engineers & Fintechs | REST API v1 access, Live & Test API keys (`sk_`, `pk_`), dedicated virtual NUBANs, real-time SSE stream, webhook DLQ. |
| **10**| `ambassador` | Brand Ambassador | `BURSA-AMB-` | Campus Student Influencers | Referral link generation, student signup tracking, real-time transaction commission calculations, automated payouts. |
| **11**| `founder` | Platform Founder | `BURSA-FOU-` | Company Shareholders & Founders | Platform gross margin visibility, dividend distribution ledgers, atomic cap table accounting, executive financial reports. |
| **12**| `intern` | Platform Intern | `BURSA-INT-` | Junior Support & Operations Team | Level-1 student support ticketing, manual receipt logging assistance, merchant verification preprocessing under strict supervision. |
| **13**| `financial_secretary`| Financial Secretary | `BURSA-FSC-` | University & Faculty Auditors | Read-only ledger audit access, session dues balance tracking, reconciliation reports, and class clearance exports. |

---

## Deep Breakdown of Each Role

### 1. Super Admin (`super_admin`)
The platform owner role. Has unrestricted access across all database tables, dashboards, and financial operations.
- **KYC & Merchant Approvals**: Reviews submitted National Identification Numbers (NIN), CAC incorporation certificates, and bank account proofs for new organizations and vendors (`is_approved=True`).
- **Dispute Resolution & Arbitration**: Acts as the neutral arbiter for marketplace disputes between clients and vendors, with the power to refund clients or release escrow to vendors.
- **Risk & Fraud Controls**: Views platform-wide velocity anomalies, suspends suspicious accounts, and inspects dead-letter webhook queues.
- **Fee Configuration**: Manages base platform fees, institutional flat charges, and vendor volume tiers.

---

### 2. Organization (`organization`)
Organizations collect funds from groups of payers. An organization account can switch between three modes:
- **Institutional Mode**:
  - Configures departmental and faculty registration dues (V2 System).
  - Configures multi-tier student pricing (e.g. 100L vs 400L vs Direct Entry).
  - Adds custom form fields (Matric Number, Level, Residential Hall).
  - Uploads institutional logos and digital stamps for anti-tamper PDF receipts.
  - Monitors "Pay-for-Me" requests sent to parents and sponsors.
- **Public Mode**:
  - Creates open payment links (`bursapay.com/pay/<slug>/`).
  - Launches crowdfunding campaigns with target amounts, live progress bars, and Wall of Fame donor walls.
  - Supports installment split payment sessions.
- **Gateway Mode**:
  - Unlocks programmatic developer access to the REST API v1.

---

### 3. Student (`student`)
Academic end-payers enrolled in recognized Nigerian universities, polytechnics, and colleges.
- **Profile Attributes**: `matric_number`, `institution`, `faculty`, `department`, `level` (100–600L), `student_type` (Fresher, Staylite, Direct Entry).
- **Payment History**: Accesses all past session dues and departmental payments.
- **Examination Clearance**: Generates cryptographically verifiable receipts with standard invoice references (`bursapay.xxxxxxxx-xxxx`).
- **Sponsor Delegation**: Initiates "Pay-for-Me" links allowing sponsors to pay remotely via card or bank transfer.

---

### 4. Event Registrar (`registrar`)
Event creators and promoters managing ticket sales, venue access, and stage lineups.
- **Event Creation**: Sets event details, public vs hidden venue locations, categories (Concerts, Conferences, Banquets).
- **Multi-Tier Ticketing**: Defines Unlimited or Capped tiers (Early Bird, Regular, VIP, VVIP, Table for 10).
- **Interactive Seating Engine**: Configures seat grids with 5-minute atomic locking during checkout to prevent duplicate sales.
- **Guest Lineup Cards**: Adds artists, speakers, DJs, MCs, and hosts with social links and share cards.
- **Access Control & Check-in**: Scans high-density ticket QR codes using the built-in camera scanner.
- **Audit Reports**: Generates HMAC-signed financial and attendance audit reports.

---

### 5. Mini-Registrar (`mini_registrar`)
Delegated event staff appointed by a Registrar to assist during live events.
- **Staff Badges**: Receives a verifiable digital staff pass with a QR code and strict time-bounded access.
- **Granular Permissions**:
  - `check_in_staff`: Fast QR scanning and attendee validation at gates.
  - `ticket_manager`: Adjusts ticket inventory and pricing tiers.
  - `accountant`: Monitors ticket volume and gross sales in real time.
- **Immutable Audit Trail**: All actions logged in `MiniRegistrarAuditLog`.

---

### 6. Mini Organization Staff (`mini_org`)
Departmental clerks and faculty executives assisting Organization administrators.
- Verifies student fee clearance during physical faculty screening.
- Searches payer rosters by matriculation number or email.
- Generates manual receipt logs for physical cash payments when permitted.

---

### 7. Marketplace Vendor (`vendor`)
Verified creative professionals and freelancers offering specialized services (Photography, Video, Decor, Tech, Catering).
- **4-Phase Onboarding**: Business setup, Social proof, Identity KYC (NIN/CAC), and Bank payout setup.
- **Storefront & Vanity Domain**: Public storefront at `brand.bursapay.com` or `bursapay.com/v/<slug>/`.
- **Milestone Escrow Protection**: Receives upfront deposits safely held in `BookingEscrow`; funds released upon client satisfaction.
- **Instant Release Option**: Configures an instant float percentage (e.g., 20%) released immediately on booking to cover materials.
- **Calendar & Availability**: Manages weekly schedules, blackout dates, and instant booking rules.

---

### 8. Marketplace Client (`client`)
Individuals, corporations, and event planners booking services through the marketplace.
- Browses verified vendor portfolios and package pricing.
- Submits custom booking requests with date and location requirements.
- Pays upfront milestone deposits into BursaPay escrow.
- Communicates with vendors via in-app booking chat.
- Reviews submitted deliverables, requests revisions, or opens dispute arbitrations.
- Leaves verified 1-5 star reviews.

---

### 9. API Developer (`developer`)
Software engineers and enterprise technical teams integrating BursaPay capabilities into external apps.
- Full access to the Developer Gateway REST API v1 (`/api/v1/`).
- API Key management with Test (`sk_test_`, `pk_test_`) and Live (`sk_live_`, `pk_live_`) environments.
- Provisions dedicated virtual accounts (NUBAN) with instant transfer notifications.
- Manages webhook endpoints with automatic retries, signature verification, and Dead Letter Queue (DLQ).
- Subscribes to real-time Server-Sent Events (SSE) at `/api/v1/events/stream/`.

---

### 10. Brand Ambassador (`ambassador`)
Student leaders and campus advocates promoting BursaPay adoption.
- Unique referral tracking code and shareable URLs.
- Real-time conversion tracking across student onboarding and departmental signups.
- Automated commission calculations and transparent wallet payout requests.

---

### 11. Founder (`founder`)
Company founders and platform equity partners.
- High-level financial reporting dashboard showing total volume, gateway processing costs, and gross margins.
- Dividend and profit distribution ledgers.
- Cap table distribution and automated quarterly payout calculations.

---

### 12. Intern (`intern`)
Operational team interns supporting platform operations under restricted supervision.
- Handles tier-1 user support inquiries and ticket routing.
- Assists with manual cash receipt indexing and verification logs.
- Pre-screens merchant KYC documentation for Super Admin review.

---

### 13. Financial Secretary (`financial_secretary`)
University, faculty, or departmental finance executives with independent oversight.
- Read-only financial ledger access to verify all collections against institutional bank statements.
- Real-time clearance roster verification for exam eligibility.
- Automated generation of CSV and PDF financial reports for university bursary audits.
