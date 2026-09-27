# Institutional Payment Architecture & User Guide (V2)

The **Institutional Mode (V2)** is BursaPay's specialized academic payment system engineered for tertiary institutions (Universities, Polytechnics, Colleges of Education), Student Union Governments (SUG), Faculties, and Academic Departments.

---

## 🏛️ PART 1: Organization Onboarding & Setup

### Step 1: Registration & Admin Approval
1. Navigate to `https://bursapay.com/register/` and select **Organization**.
2. New accounts start with `is_approved=False` to prevent fake university bodies.
3. Once approved by BursaPay Super Admin, log in to your dashboard.

### Step 2: Select Mode & Link Institution
1. On first login, select **"Institutional Mode"**.
2. Select your tertiary institution from the accredited university list (`/dashboard/org/select-institution/`).
3. Complete institutional profile (`/dashboard/org/v2/onboarding/`) with:
   - **Full Organization Name** (e.g., *National Association of Computer Science Students - UNILAG Chapter*).
   - **Organization Acronym** (e.g., *NACOSS*).
   - **Official Contact Number**.
4. Configure settlement bank account under `/dashboard/bank-setup/` (NUBAN account verified instantly via Paystack).

### Step 3: Configure Branding (Logo & Digital Stamp)
1. Go to **Dashboard > Branding** (`/dashboard/org/v2/branding/`).
2. Upload your **Faculty/Department Logo** and **Official Executive Stamp**.
3. These assets are automatically stamped on all student PDF receipts with high-density anti-tamper security.

---

## 🛠️ PART 2: How to Create an Academic Payment (Step-by-Step)

Organizations can create fee collections for departmental dues, faculty handbooks, lab coats, excursion fees, and gala tickets.

```
┌─────────────────────────────────────────────────────────────┐
│ 📝 CREATE INSTITUTIONAL PAYMENT (V2)                        │
│                                                             │
│ Payment Title: 2026/2027 NACOSS Annual Departmental Dues    │
│ Payment Category: [ Departmental Dues          ▼ ]          │
│ University: [ University of Lagos             ▼ ]          │
│ Deadline / Expiry: [ 2026-12-15               📅 ]          │
│ Status: [ Active (Live)                       ▼ ]          │
└─────────────────────────────────────────────────────────────┘
```

### Step 1: Create Base Payment Record
1. Go to **Dashboard > Create Payment** (`POST /dashboard/org/v2/payments/create/`).
2. Enter the **Title**, select the **Payment Category**, and set an optional **Deadline**.
3. Click **Save & Continue**. BursaPay generates a unique public checkout slug (e.g., `bursapay.com/pay/nacoss-unilag-dues-2026`).

### Step 2: Configure Multi-Tier Student Categories & Pricing
Academic dues differ by student level and entry mode. BursaPay lets you set custom pricing per tier:

| Student Category | Sample Price | Description / Perks Included |
|---|---|---|
| **100 Level (Freshers)** | ₦12,500 | Includes Departmental Handbook, T-shirt & Lab Access |
| **Direct Entry (200L DE)** | ₦10,000 | Includes Departmental Handbook & Kit |
| **Returning (200L – 400L)**| ₦6,500 | Standard Annual Session Dues |
| **Final Year (500L)** | ₦8,500 | Includes Final Year Project Levy |

1. Go to **Manage Categories & Pricing** (`/dashboard/org/v2/categories/student/`).
2. Select the student category and input the exact amount for each tier.

### Step 3: Add Dynamic Custom Form Fields
Collect necessary departmental data during checkout:
- **Matriculation / JAMB Number** (Required text field)
- **Level / Academic Year** (Dropdown: 100L, 200L, 300L, 400L, 500L)
- **Residential Hall / Hostel** (Optional text field)
- **Specialization / Option** (e.g., Software Engineering vs Cyber Security)

### Step 4: Publish & Share Checkout URL
Your payment link is now live at `https://bursapay.com/pay/<slug>/`. Share it on departmental WhatsApp groups, Telegram channels, and official notice boards.

---

## 💳 PART 3: How Students / Payers Complete Payment

### Option A: Standard Direct Student Payment Flow

```
1. Student opens Link ──► 2. Selects Category (e.g. 100L) ──► 3. Fills Matric/Dept Info
                                                                        │
                                                                        ▼
6. Verified PDF Receipt ◄── 5. Instant Payment Confirmation ◄── 4. Card / Transfer / USSD
```

1. **Open Payment Link**: Student visits `bursapay.com/pay/<slug>/`.
2. **Select Category**: Selects their student level (e.g., "100 Level Freshers"). The system automatically updates the exact fee payable (₦12,500).
3. **Fill Information**: Inputs Full Name, Email, Matric Number, and Department.
4. **Choose Payment Method**: Card, Bank Transfer, USSD, or Apple Pay.
5. **Instant Confirmation**: Transaction is confirmed via Paystack webhook; status converts immediately to `is_paid=True`.
6. **Download Anti-Tamper Receipt**: A PDF receipt stamped with the department logo, executive signature, and security QR code is sent to the student's email and downloadable on-screen.

---

### Option B: "Pay for Me" (Parent / Sponsor Delegation Flow)

When a student relies on a parent, guardian, or external sponsor to pay:

```
┌─────────────────────────────────────────────────────────────┐
│ 👨‍👦 "PAY FOR ME" SPONSOR WORKFLOW                          │
│                                                             │
│ 1. Student checks [✔] "Pay for Me" on checkout form         │
│ 2. Enters Sponsor Name: Chief Bamidele Ade                  │
│ 3. Enters Sponsor Email: sponsor.ade@gmail.com              │
│ 4. BursaPay sends UUID checkout link to Sponsor             │
│ 5. Sponsor pays remotely via Bank Transfer or Card          │
│ 6. Verified receipt is sent directly to the STUDENT's email │
└─────────────────────────────────────────────────────────────┘
```

1. The student completes their details on the checkout form, checks the **"Pay for Me"** checkbox, and enters their sponsor's name and email.
2. BursaPay creates a pending `GuestPaymentRecord` and issues a cryptographically secure UUID link: `https://bursapay.com/pay/for-me/<token>/`.
3. The sponsor receives an email breakdown with the student's name, matric number, and departmental fee items.
4. The sponsor clicks the link and pays.
5. Upon successful settlement, the **student** receives the verified clearance receipt, and the organization's roster updates in real time.

---

## 🔍 PART 4: Verification & Departmental Screening

Departmental executives and examination officers can verify fee clearance using 3 methods:

1. **Instant QR Scanner**: Scan the QR code on the student's phone or printed receipt using the camera scanner at `/dashboard/verify-receipt/`.
2. **Matric Number Lookup**: Search the real-time payer registry at `/dashboard/org/v2/payers/` to see complete transaction timestamps and student details.
3. **CSV & Excel Export**: Download the official class roster (`/dashboard/org/v2/payments/export/`) filtered by level or session.
