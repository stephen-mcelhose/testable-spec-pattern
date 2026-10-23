---
type: proposal
title: "Contract-First as the Foundation"
description: Hypothesis that defining a contract first (OpenAPI, protobuf, ROD) naturally produces the spec that testable spec tests verify, resolving the hardest question — what is a MUST — before implementation begins.
tags: [contract-first, openapi, rod, testable-spec, ideation]
timestamp: 2026-10-23T22:45:43Z
---

# Ideation: Contract-First as the Foundation

**Status:** hypothesis — not yet validated as discipline  
**Related:** [domain-names-as-bridge.md](./domain-names-as-bridge.md)

---

## The observation

A contract-first approach — defining the interface before writing implementation — produces a natural spec. The contract (OpenAPI, protobuf, ROD resource definition) *is* the requirement, expressed in a machine-readable form. Everything that follows is an attempt to satisfy it.

This creates a different starting point than the "write code, then test it" loop:

```
contract-first flow:
  Contract (OpenAPI / proto / ROD)
    → domain types derived from the contract
    → testable spec written against the contract's guarantees
    → implementation satisfies the spec
    → integration test verifies the contract end-to-end

code-first flow:
  Implementation
    → tests that describe what the code does
    → (contract inferred after the fact, if at all)
```

The contract-first flow produces a testable spec for free — the contract *is* the spec. The code-first flow requires working backwards to discover what the contract should have been.

---

## Why this matters for testable specs

In the testable spec pattern, we write tests that read like RFC 2119 requirements. The hardest part is knowing *what* to write as a MUST — what is actually a contract vs. what is an implementation detail.

A contract (OpenAPI spec, proto definition, ROD resource) answers that question before you start. The fields, operations, error codes, and invariants are explicit. You are not discovering the contract from the code; you are verifying that the code satisfies a contract that existed first.

This also connects directly to the [domain names](./domain-names-as-bridge.md) idea: contract formats like OpenAPI and ROD are inherently domain-named (resources, not records; operations, not functions). The domain model arrives with the contract.

---

## Formats in scope

| Format     | When to reach for it                                                  |
| ---------- | --------------------------------------------------------------------- |
| **OpenAPI**  | HTTP APIs — request/response shape, status codes, error bodies        |
| **Protobuf** | gRPC or multi-language services — strong typing, streaming            |
| **ROD**      | Resource-Oriented Design (AIP pattern) — resource lifecycle, standard methods (`Get`, `List`, `Create`, `Update`, `Delete`), and long-running operations |

ROD is worth calling out specifically because it is not just a wire format — it is a design discipline. Resources have names, revisions, and lifecycle states. Defining a resource in ROD terms forces you to name things in domain language before writing a line of implementation, which directly feeds the [domain names](./domain-names-as-bridge.md) pattern.

---

## The hypothesis

> Contract-first development (OpenAPI, proto, ROD) provides the spec that testable spec tests verify. The contract resolves the hardest question — "what is a MUST?" — before any implementation is written. Domain names follow naturally from the contract's resource definitions.

---

## Open questions

- Can this pattern work retroactively — writing a contract from existing code, then deriving tests from the contract? Or does it only work cleanly when the contract is written first?
- How do you handle contracts that evolve? If the OpenAPI spec changes, the spec tests must change too. Is that friction or is it exactly the right friction?
- For purely internal functions (helpers, transformations) with no wire contract, what plays the role of the contract? Is `discipline.md`-style naming sufficient, or does something else need to anchor the MUST?

---

## What this is NOT yet

This is not a rule in the discipline. It is a hypothesis about the upstream condition that makes testable specs easier to write well. The practical consequence — if validated — might be: "Start every new service with an OpenAPI or ROD definition. Write your first spec tests against that definition before touching implementation."
