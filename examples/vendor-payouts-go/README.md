# Go Vendor Payouts & Transfers Example

This example demonstrates how to:
1. Accept payments via the BursaPay REST API in Go.
2. Disburse bulk vendor payouts via `/api/v1/transfers/bulk/`.
3. Handle incoming webhook notifications with HMAC SHA-256 signature verification.

---

## Setup

```bash
# Set your environment variables
export BURSAPAY_SECRET_KEY=bp_sec_test_DEMO_KEY_HERE
export BURSAPAY_WEBHOOK_SECRET=sec_wh_DEMO_SECRET_HERE

# Run the server
go run main.go

# Or run with reference to trigger automated payout
go run main.go --reference ORD-2026-9812
```
