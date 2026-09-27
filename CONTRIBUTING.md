# Contributing to BursaPay Open Platform & Docs

We love contributions! Whether you're fixing a typo in the documentation, suggesting a new API feature, improving SDK wrappers, or submitting an integration example, your help makes BursaPay better for the entire developer community.

---

## Code of Conduct

We are committed to providing a welcoming, diverse, and harassment-free environment for everyone. Please be respectful, constructive, and collaborative in all discussions and code reviews.

---

## How to Contribute

### 1. Reporting Bugs
- Check the [open issues](https://github.com/bursapay/bursapay-public/issues) to avoid duplicates.
- Use our [Bug Report Template](.github/ISSUE_TEMPLATE/bug_report.md) with steps to reproduce, expected vs actual behavior, and relevant environment details.

### 2. Suggesting Enhancements
- Open a feature request using our [Feature Request Template](.github/ISSUE_TEMPLATE/feature_request.md).
- Describe the business use case and proposed developer interface.

### 3. Submitting Pull Requests
1. Fork this repository and create a new feature branch (`git checkout -b feature/awesome-feature`).
2. Make your improvements. Ensure markdown documents are well formatted and code examples are tested.
3. Commit your changes with a clear, concise commit message following [Conventional Commits](https://www.conventionalcommits.org/):
   - `docs: update webhook signature verification guide`
   - `feat(sdk): add idempotency header support in python example`
   - `fix(specs): correct refund response schema in openapi.yaml`
4. Push to your branch and submit a Pull Request against `main`.

---

## Documentation Guidelines

- **Accuracy**: Code examples and endpoint schemas must exactly match the BursaPay Developer Gateway API v1.
- **Security First**: Never commit real secret keys, private credentials, or live API credentials in examples or docs. Use mock keys like `bp_sec_test_KEY_HERE` or `bp_pub_test_KEY_HERE`.
- **Clarity**: Use fenced code blocks with language identifiers, clear callouts, and clean formatting.

Thank you for building with BursaPay! 🚀
