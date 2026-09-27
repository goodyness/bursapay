# WhatsApp Conversational Commerce Bot

BursaPay features a native, end-to-end **WhatsApp Conversational Commerce Engine** (`whatsapp/`) that brings the entire payment, ticketing, receipt generation, and support ecosystem directly into WhatsApp.

---

## 📱 Conversational Commerce Capabilities

```
┌─────────────────────────────────────────────────────────────┐
│ 💬 BURSAPAY OFFICIAL WHATSAPP BOT                           │
│                                                             │
│ "Welcome to BursaPay. Your payments & events hub.           │
│  Tap below to see our services:"                            │
│                                                             │
│ ┌─────────────────────────────────────────────────────────┐ │
│ │ 🎟 Buy Ticket       — Purchase event passes in chat      │ │
│ │ 💳 Make Payment     — Pay school bills & public dues    │ │
│ │ 📄 Get Receipt      — Download official PDF receipts    │ │
│ │ 📲 Get QR Code      — Resend ticket QR entry passes     │ │
│ │ 🛠 Report Issue     — Resolve payment discrepancies     │ │
│ │ 📞 Contact Support  — Connect with human support agents │ │
│ └─────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

---

## 🚀 Key Interactive WhatsApp User Flows

### 1. 💳 Pay School Bills & Public Dues in WhatsApp (`flow_make_payment`)
Students and payers can settle fees without leaving WhatsApp:
1. Payer taps **"Make Payment 💳"**.
2. Bot displays an interactive list of accredited universities and organizations.
3. Payer selects their institution, department, and student category (e.g. *100L Freshers*).
4. Bot returns a secure checkout session link or dynamic virtual account NUBAN.
5. Once settled, the bot immediately uploads the **official anti-tamper PDF receipt** directly into the WhatsApp thread.

---

### 2. 🎟 Buy Event Tickets Directly in WhatsApp (`flow_buy_event`)
Attendees can discover and purchase tickets directly within chat:
1. Payer taps **"Buy Ticket 🎟"**.
2. Bot presents upcoming featured concerts, summits, and festivals.
3. Attendee selects their preferred ticket tier (Regular, VIP, Table) and quantity.
4. Completes checkout via card or instant bank transfer.
5. Bot sends the **high-density Ticket QR entry pass** and ticket ID (`BURSAL-XXXXXX`) directly into WhatsApp for quick presentation at the gate.

---

### 3. 📄 Retrieve Lost PDF Receipts in WhatsApp (`flow_get_receipt`)
1. User taps **"Get Receipt 📄"**.
2. Enters their transaction reference or registered matriculation number/email.
3. Bot validates the payment in real time and attaches the official PDF receipt.

---

### 4. 📲 Instant Ticket QR Retrieval (`flow_get_qr`)
1. Attendee taps **"Get QR Code 📲"**.
2. Enters their email or phone number.
3. Bot delivers a fresh QR code image directly to their screen ready for gate scanning.

---

### 5. 🛠 Report Payment Issues & Open Support Tickets (`flow_report_issue`)
1. User taps **"Report Issue 🛠"**.
2. Submits bank transaction references and details.
3. An official `SupportTicket` is created in BursaPay's support queue, and the user receives a unique tracking ID (`TKT-XXXXXX`).

---

## 👑 Admin & Founder WhatsApp Operations

Authorized platform administrators and founders can manage operations securely via WhatsApp:

- **Pending Approvals (`adm_pending_orgs`, `adm_pending_regs`):** Review and approve new institutional organizations and event registrars directly from chat.
- **Withdrawal Approvals (`adm_org_withdrawals`, `adm_reg_withdrawals`):** Review merchant payout requests and approve disbursements.
- **Founder Financial Insights (`fnd_balance`, `fnd_transactions`):** Check real-time gross volume, platform fee earnings, and recent profit settlements.
- **Real-time Platform Health (`adm_platform_stats`):** Query live system uptime, active transaction counts, and gateway success rates.
