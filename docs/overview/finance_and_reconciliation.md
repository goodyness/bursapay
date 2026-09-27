# Double-Entry Finance & Automated Reconciliation

BursaPay incorporates enterprise-grade accounting rigor, utilizing double-entry bookkeeping, atomic journal transactions, and automated 3-way reconciliation to guarantee that every kobo is balanced.

---

## Double-Entry Ledger Architecture

Every monetary transaction in BursaPay produces balanced debit and credit entries across strict account categories:

| Account Type | Normal Balance | Purpose |
|---|---|---|
| **Asset (Bank / Float)** | Debit | Physical funds held with banking and gateway partners. |
| **Liability (User Wallets / Escrow)** | Credit | Platform obligations owed to organizations, vendors, and registrars. |
| **Revenue (Platform Fees)** | Credit | BursaPay processing fees, commission revenue, and service charges. |
| **Expense (Gateway Charges)** | Debit | Underlying processing costs paid to payment rails and NIP switches. |

### Atomic Balance Guarantee
All wallet operations (deposits, transfers, escrow locks, withdrawals) are executed inside database transaction blocks with row-level locks (`select_for_update()`), completely preventing race conditions or double-spending.

---

## Automated 3-Way Reconciliation Engine

BursaPay continuously executes automated reconciliation audits between three independent data sources:

```
                  ┌──────────────────────────────┐
                  │   1. Payment Gateway Logs    │
                  │  (Paystack / NIP Switch API) │
                  └──────────────┬───────────────┘
                                 │
                                 ▼
┌─────────────────────────┐             ┌─────────────────────────┐
│  2. BursaPay Internal   │◄───────────►│   3. Bank Settlement    │
│     Database Records    │   Matched   │    Statement Feeds      │
└─────────────────────────┘             └─────────────────────────┘
```

### Discrepancy Classification:
- **Matched (`STATUS: RECONCILED`)**: Amounts, references, and fees match across all three sources.
- **Timing Difference (`STATUS: PENDING_SETTLEMENT`)**: Gateway confirmed, awaiting interbank clearing window.
- **Amount Mismatch (`STATUS: DISCREPANCY`)**: Flags discrepancy to platform risk team for manual review.
- **Unclaimed Credit (`STATUS: UNMATCHED`)**: Bank received funds without matching invoice reference.

---

## Platform Safety & Risk Engine

To safeguard merchant funds and maintain compliance with financial regulations:

1. **Velocity Limits**: Automatically rate-limits rapid transaction bursts from unverified IP addresses.
2. **Withdrawal Fraud Scoring**: Payout requests are evaluated against user history, past dispute rates, and KYC tier before approval.
3. **Automated Ledger Freezing**: Any mathematical discrepancy across account balances triggers an immediate circuit breaker to prevent fund leakage.
