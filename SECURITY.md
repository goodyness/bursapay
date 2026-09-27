# Security Policy

BursaPay takes platform security, data integrity, and responsible disclosure very seriously. As financial infrastructure serving universities, organizations, and developers across Nigeria, we prioritize the protection of sensitive transactional and identity data.

---

## Reporting a Vulnerability

If you discover a security vulnerability within BursaPay services, API, or client libraries, **please do not disclose it publicly**. We appreciate your help in disclosing it to us responsibly.

### How to Report

Please email our security engineering team directly:
- **Email:** `security@bursapay.com`
- **Subject Line:** `[Vulnerability Report] <Brief description>`

### What to Include
1. A detailed description of the issue.
2. Step-by-step proof-of-concept (PoC) or reproducible script/curl command.
3. Impact assessment (e.g., unauthorized data access, privilege escalation, balance manipulation).
4. Any potential mitigations or suggested fixes.

---

## Our Commitment

When you report a security vulnerability responsibly:
- We will acknowledge receipt of your report within **24 hours**.
- We will investigate and provide regular status updates until resolution.
- We will not pursue legal action against security researchers who follow responsible disclosure principles and test only against their own sandbox accounts.
- We will credit you in our public security release notes (unless you prefer to remain anonymous).

---

## Safe Harbor & Testing Rules

- **Only test against test mode credentials (`bp_sec_test_KEY_HERE` / `bp_pub_test_KEY_HERE`) or your own accounts.**
- **Never attempt to access, modify, or exfiltrate another user's or organization's data.**
- **Do not perform Denial of Service (DoS/DDoS) attacks or degrade platform performance for active users.**
