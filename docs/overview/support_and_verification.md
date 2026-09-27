# Public Verification Portals, Support & Dispute SLAs

BursaPay provides public verification tools and clear support service level agreements (SLAs) to guarantee trust, accountability, and institutional compliance.

---

## 🔍 Public Verification Tools

### 1. Public Receipt Verification Portal
Anyone (university bursars, departmental officers, parents) can independently verify the authenticity of a BursaPay receipt without logging in:
- **URL:** `https://bursapay.com/dashboard/verify-receipt/`
- **Lookup Method:** Enter the unique invoice reference (e.g. `bursapay.a1b2c3d4-e5f6`) or scan the physical QR code with a camera.
- **Verification Response:** Returns verified payer name, matriculation number, institution, payment date, fee breakdown, and digital verification seal.

---

### 2. Cryptographic Event Audit Verification Portal
Institutions, sponsors, and auditors can verify official event revenue and attendance reports:
- **URL:** `https://bursapay.com/dashboard/audit/verify/<report_id>/`
- **Integrity Check:** Compares the stored HMAC-SHA256 digital signature against database records to prove that financial figures and attendance counts have not been altered.

---

## ⏱️ Dispute Resolution Service Level Agreements (SLAs)

| Dispute Stage | BursaPay Commitment | Actions Taken |
|---|---|---|
| **Dispute Acknowledgment** | **< 24 Hours** | Immediate temporary fund lock in escrow and notification dispatched to both parties. |
| **Evidence Submission Window** | **72 Hours** | Merchant / Vendor provides proof of fulfillment, delivery logs, or attendee sign-off. |
| **Arbitration & Resolution** | **< 5 Business Days**| Super Admin arbitrates based on submitted chat transcripts and proof; funds released or refunded. |

---

## 🎫 Customer Support Channels

- **Developer Support:** `developers@bursapay.com`
- **General Support & Inquiries:** `support@bursapay.com`
- **Security & Responsible Disclosure:** `security@bursapay.com`
- **In-App Support Ticket Center:** Accessible directly from any logged-in dashboard (`/support/`).
