# Bursa AI & Self-Service Knowledge Base

BursaPay incorporates **Bursa AI**—an intelligent, customer-facing conversational assistant—and an enterprise **Knowledge Base (KB)** to provide 24/7 instant guidance for students, payers, event attendees, vendors, and developers.

---

## 🤖 Bursa AI: Customer-Facing Conversational Assistant

**Bursa AI** is embedded directly into BursaPay web dashboards, checkout pages, and support touchpoints. Powered by modern large language models (**Google Gemini**), Redis asynchronous state caching, and domain-grounded knowledge pipelines, Bursa AI answers questions and guides users in real time.

```
┌─────────────────────────────────────────────────────────────┐
│ 💬 USER INQUIRY                                             │
│ "How do I pay my 200L faculty dues with my father's card?"  │
└──────────────┬──────────────────────────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────────────────────────┐
│ 🛡️ DOMAIN GUARD & INTENT ROUTER                             │
│ • Validates message is BursaPay-related                     │
│ • Routes intent: 'payment_pay_for_me' & 'institutional_v2'  │
└──────────────┬──────────────────────────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────────────────────────┐
│ 📚 KNOWLEDGE BASE GROUNDING ENGINE                          │
│ • Retrieves authoritative KB articles & verified policies   │
│ • Prevents AI hallucinations with strict domain boundaries  │
└──────────────┬──────────────────────────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────────────────────────┐
│ ⚡ BURSA AI STREAMED RESPONSE                               │
│ "You can use BursaPay's 'Pay for Me' feature:               │
│  1. Open your department's payment link.                    │
│  2. Select '200 Level Returning'.                           │
│  3. Check 'Pay for Me' and enter your father's email.       │
│  4. He will receive a secure checkout link to pay, and the │
│     official receipt will be emailed directly to you!"      │
└─────────────────────────────────────────────────────────────┘
```

---

## ⚡ What Bursa AI Does for Users

| User Persona | Common Inquiries & Tasks Handled by Bursa AI |
|---|---|
| **🎓 Students / Payers** | • Guides students through selecting student tiers (100L vs DE vs Staylite).<br>• Explains the "Pay for Me" sponsor checkout process.<br>• Helps locate lost payment references and receipts.<br>• Provides departmental clearance instructions. |
| **🎟️ Event Attendees** | • Explains how the 5-minute seating reservation lock works.<br>• Assists with lost ticket recovery and QR pass resends.<br>• Answers questions regarding venue rules, refund policies, and dress codes. |
| **🎨 Marketplace Clients & Vendors**| • Explains milestone escrow protection and how funds are held safely.<br>• Guides vendors through the 4-phase NIN/CAC KYC onboarding process.<br>• Explains dispute escalation and delivery approval timelines. |
| **💻 Developers & Merchants** | • Answers API authentication, rate limits, and webhook signature questions.<br>• Recommends appropriate SDKs (Python, JS) and CLI commands.<br>• Explains idempotency keys and testing with sandbox cards. |

---

## 🛡️ Architecture & Safety Controls

1. **Strict Domain Guard:** Bursa AI filters out unrelated queries, ensuring conversations remain focused entirely on payments, ticketing, escrow, and platform guidance.
2. **Intent-Specific Fast Paths:** Recognizes common intents (`event_create`, `payment_failed`, `ticket_retrieve`, `withdrawal_status`, `vendor_trust`) to provide immediate, structured answers.
3. **Session Memory with Redis:** Maintains contextual multi-turn conversation history across up to 10 dialogue turns.
4. **Automated KB Grounding Metrics (`KBGroundingEvent`):** Platform tools continuously measure response accuracy against approved Knowledge Base articles to prevent misinformation.

---

## 📚 Self-Service Knowledge Base (KB)

The BursaPay Knowledge Base is a comprehensive, searchable library of help documentation, interactive tutorials, and FAQs organized by category:

```
Knowledge Base Taxonomy:
├── 🏛️ Institutional Payments & Dues (V2 Guides, Receipts, Pay-for-Me)
├── 🎟️ Events & Ticketing (Seating, QR Passes, Gate Check-in)
├── 🎨 Vendor Marketplace & Escrow (KYC, Milestone Deliveries, Disputes)
├── 💻 Developer Gateway API (Keys, Webhooks, Idempotency, SDKs)
├── 💳 Wallet & Bank Settlements (Withdrawal schedules, Fee breakdowns)
└── 🔐 Security & Account Verification (2FA, NIN verification, Scams)
```

### Knowledge Base Features:
- **Article Feedback System (`KBArticleFeedback`):** Users vote on whether an article was helpful or not, driving automated content quality improvements.
- **Search Analytics (`KBSearchEvent`):** Tracks popular search queries to identify missing help topics and optimize user search discovery.
- **Automated AI Health Scoring (`kb_ai_health`):** Evaluates article readability, completeness, and keyword coverage.
