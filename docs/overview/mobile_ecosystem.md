# Mobile Ecosystem & Mobile API Architecture

The BursaPay Mobile Ecosystem delivers fast, secure, native mobile experiences for students, event organizers, and campus staff on iOS and Android.

---

## 📱 Mobile Applications

```
┌─────────────────────────────────────────────────────────────┐
│                    BURSAPAY MOBILE SUITE                    │
├──────────────────────────────┬──────────────────────────────┤
│ 🎓 STUDENT MOBILE APP        │ 🎟️ REGISTRAR & SCANNER APP   │
├──────────────────────────────┼──────────────────────────────┤
│ • Biometric FaceID/Finger    │ • Sub-second QR ticket scan  │
│ • Offline-ready PDF receipts │ • Real-time attendance count │
│ • Push fee notifications     │ • Offline validation caching │
│ • Sponsor delegation chat    │ • Multi-gate sync alerts     │
└──────────────────────────────┴──────────────────────────────┘
```

---

## 🔑 Key Mobile Features

### 1. Offline-Ready Verified Digital Passes
- Purchased event tickets and academic clearance receipts are cryptographically signed and stored in local secure mobile storage.
- Students and attendees can present their QR clearance passes at gate checkpoints even without active mobile internet connectivity.

### 2. High-Speed Camera QR Scanner
- Built for event registrars and mini-registrar gate staff.
- Scans high-density QR passes in under **300 milliseconds**.
- Real-time audio-visual feedback (Green chime for valid entry, Red alert for duplicate/reused pass).

### 3. Push Notification Architecture (Expo & Web Push)
- Powered by `expo_push_token` and Web Push VAPID standards.
- Instant alerts dispatched when:
  - Departmental fee payment is confirmed.
  - Sponsor pays a "Pay-for-Me" request.
  - Event registration is approved or 24-hour event reminder fires.
  - Vendor receives a new marketplace booking request.

---

## 📡 Mobile API Endpoints (`/api/v1/mobile/`)

- `POST /api/v1/mobile/auth/login/` — Biometric & token-based session login.
- `POST /api/v1/mobile/push/register/` — Register device push notification token.
- `GET /api/v1/mobile/tickets/active/` — Fetch attendee's offline-ready event passes.
- `POST /api/v1/mobile/scan/validate/` — Gate scanner validation endpoint.
