# Institutional Payment Architecture (V2)

The Institutional Mode is BursaPay's flagship academic payment solution, engineered to solve the complex fee collection and accounting challenges faced by African universities, faculties, departments, and Student Union Governments (SUG).

---

## Key Problems Solved

1. **Elimination of Fake Bank Tellers**: Manual bank tellers and receipts are prone to forgery and duplication. BursaPay enforces cryptographic receipt verification.
2. **Category-Specific Dues**: Fresh students (100L) often pay different faculty registration fees than returning students (200L–500L) or Direct Entry admits.
3. **Third-Party Payer Support ("Pay-for-Me")**: Most African undergraduate fees are paid by parents, guardians, or sponsors who are not present on campus.
4. **Instant Financial Clearance**: Departmental officers can instantly verify payment status using student matric numbers or receipt QR codes.

---

## System Workflow & Lifecycle

```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│ Organization    │       │ Student Form    │       │ Checkout & Gate │
│ Configures V2   ├──────►│ Entry & Category├──────►│ Paystack / Bank │
│ Payment & Slugs │       │ Selection       │       │ Transfer Rail   │
└─────────────────┘       └────────┬────────┘       └────────┬────────┘
                                   │                         │
                       "Pay for Me"│                         │ Successful
                       Delegation  ▼                         ▼ Webhook
                          ┌─────────────────┐       ┌─────────────────┐
                          │ Sponsor Email   │       │ Anti-Tamper PDF │
                          │ UUID Link       │       │ Receipt & QRs   │
                          │ Payer Checkout  │       │ Generated       │
                          └─────────────────┘       └─────────────────┘
```

---

## Core Capabilities

### 1. Multi-Tier Student Categories & Pricing
Organizations can define tiered pricing structures per payment item:
- **Freshers (100L / ND1 / HND1)**: ₦12,000 (includes departmental kit & handbook)
- **Returning Students (200L–500L)**: ₦7,500
- **Direct Entry (DE)**: ₦10,000
- **Postgraduate / Alumni**: Custom fee structures

### 2. Custom Dynamic Form Fields
Organizations can collect custom institutional metadata during checkout without writing code:
- Matriculation / JAMB Registration Number
- Level (100L, 200L, 300L, 400L, 500L)
- Faculty & Department selection dropdowns
- Residential Hall or Hostel identification
- Custom text, numeric, and dropdown fields

### 3. "Pay For Me" Sponsor Delegation Flow
When a student does not have sufficient mobile wallet or card balance:
1. The student completes the checkout form and checks **"Pay for Me"**.
2. They input their sponsor's (parent/guardian) full name and email address.
3. BursaPay generates a secure, cryptographically signed UUID token and dispatches an email to the sponsor with the checkout link: `https://bursapay.com/pay/for-me/<token>/`.
4. The sponsor views the exact breakdown of the student's departmental dues and settles payment.
5. Upon successful settlement, the **student** receives the verified receipt, and the organization's roster is immediately updated.

### 4. Anti-Tamper Digital PDF Receipts
Every completed transaction automatically generates a PDF receipt stamped with:
- **Invoice Reference Number**: Standardized format `bursapay.xxxxxxxx-xxxx`.
- **Institution & Department Logo / Stamp**: Uploaded by the organization administrator.
- **Embedded Security QR Code**: Directs campus officials to `https://bursapay.com/dashboard/verify-receipt/` to validate authenticity against the database.
- **Payer Details**: Full student name, matric number, level, and timestamp.

---

## Verification & Audit Portal

Campus administrators, lecturers, and departmental executives can verify student payment status through:
- **Live Search**: Query by Matric Number or Email in the Organization Dashboard (`/dashboard/org/v2/payers/`).
- **QR Scanning**: Scan the physical or digital receipt via mobile camera at examination hall entrances.
- **CSV & Excel Export**: Generate comprehensive class payment lists filtered by academic session, semester, level, or payment status.
