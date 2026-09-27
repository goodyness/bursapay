# BursaPay Node.js / Express Integration Starter

A ready-to-run Express.js backend demonstrating payment initialization and cryptographic HMAC-SHA256 webhook signature verification.

---

## Setup & Run

1. **Install Dependencies:**
   ```bash
   npm install
   ```

2. **Configure Environment:**
   Create a `.env` file in this directory:
   ```ini
   BURSAPAY_SECRET_KEY=bp_sec_test_DEMO_KEY_HERE
   PORT=5000
   ```

3. **Start the Server:**
   ```bash
   npm start
   ```

4. **Test Webhooks with BursaPay CLI:**
   ```bash
   bursapay listen --forward-to http://localhost:5000/api/webhooks
   ```
