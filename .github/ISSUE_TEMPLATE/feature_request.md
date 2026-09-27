name: Feature Request
description: Propose a new feature or improvement for the BursaPay API, SDKs, or docs
title: "[FEATURE] "
labels: ["enhancement"]
body:
  - type: markdown
    attributes:
      value: Share your ideas for improving the BursaPay developer platform!
  - type: textarea
    id: problem
    attributes:
      label: Problem / Motivation
      placeholder: Is your feature request related to a specific problem or workflow?
    validations:
      required: true
  - type: textarea
    id: solution
    attributes:
      label: Proposed Solution / Feature
      placeholder: Describe the solution you'd like to see, including potential endpoint designs or SDK methods.
    validations:
      required: true
  - type: textarea
    id: alternatives
    attributes:
      label: Alternatives Considered
      placeholder: Any alternative solutions or workarounds you've explored.
