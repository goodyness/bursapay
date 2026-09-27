name: Bug Report
description: Create a report to help us fix errors in documentation, SDKs, or examples
title: "[BUG] "
labels: ["bug"]
body:
  - type: markdown
    attributes:
      value: Thanks for taking the time to report a bug!
  - type: input
    id: endpoint_or_tool
    attributes:
      label: Affected Endpoint / Tool / SDK
      placeholder: e.g. /api/v1/payments/initialize or python-sdk
    validations:
      required: true
  - type: textarea
    id: description
    attributes:
      label: Describe the Bug
      placeholder: A clear description of what the bug is.
    validations:
      required: true
  - type: textarea
    id: reproduction
    attributes:
      label: Steps to Reproduce
      placeholder: |
        1. Call endpoint with payload...
        2. Inspect response...
        3. See unexpected behavior...
    validations:
      required: true
  - type: textarea
    id: expected
    attributes:
      label: Expected Behavior
      placeholder: What you expected to happen.
    validations:
      required: true
