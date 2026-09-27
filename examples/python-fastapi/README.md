# BursaPay Python / FastAPI Integration Starter

A high-performance async Python backend demonstrating payment checkout session creation and raw byte HMAC-SHA256 webhook signature validation.

---

## Setup & Run

1. **Create Virtual Environment & Install Dependencies:**
   ```bash
   python -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
   pip install -r requirements.txt
   ```

2. **Configure Environment:**
   Create a `.env` file:
   ```ini
   BURSAPAY_SECRET_KEY=bp_sec_test_DEMO_KEY_HERE
   ```

3. **Start the FastAPI Server:**
   ```bash
   uvicorn main:app --reload --port 8000
   ```

4. **Test Webhooks with BursaPay CLI:**
   ```bash
   bursapay listen --forward-to http://localhost:8000/api/webhooks
   ```
