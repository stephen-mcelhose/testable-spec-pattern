# Ideation: Domain Names as the Bridge

**Status:** hypothesis — not yet validated as discipline  
**Related:** [contract-first.md](./contract-first.md)

---

## The observation

When code uses domain names — not technical names — the spec test becomes a domain statement. The same word appears in the requirement, the type, and the test name. That coincidence is not cosmetic: it is a traceability chain.

```
Requirement:  "A Booking MUST be rejected if the CancellationPolicy deadline has passed."
Type:          type Booking struct { ... }  /  type CancellationPolicy struct { ... }
Test name:    "MUST reject Booking when CancellationPolicy deadline has passed"
```

Compare to a technically-named equivalent:

```
Requirement:  (implied, never written down)
Type:          type Record struct { ... }  /  cutoffOffset int
Test name:    "MUST return error when offset exceeds cutoff"
```

The second version passes tests but the chain is broken. A domain expert cannot read the test. An engineer cannot trace the test back to a requirement. When the requirement changes, no one knows which test to update.

---

## Why this matters for testable specs

The testable spec pattern tries to make tests readable as contracts. Domain names are what make that possible in practice. A MUST statement in RFC 2119 style is only as clear as the nouns in it. If those nouns are technical (`record`, `offset`, `cutoff`), the test is a contract but not a *domain* contract — it contracts the implementation, not the requirement.

Domain names force the engineer to model the domain before writing the test. That modeling is the actual design work. The test is just its artifact.

---

## The hypothesis

> Using domain names in types and functions is a necessary precondition for testable specs to bridge implementation and requirements. Without domain names, testable specs degrade into implementation contracts that survive refactors but do not survive requirement changes.

---

## Open questions

- Does this require a full domain model up front, or can names be introduced incrementally as the spec is written?
- How does this interact with generic infrastructure code (repositories, clients, caches) that is inherently non-domain?
- Is a GLOSSARY.md (à la Matt Pocock's pattern) the right way to enforce this, or does it happen naturally when you write the spec first?

---

## What this is NOT yet

This is not a rule in the discipline. It is an observation. The discipline consequence — if any — might be something like: "Before writing spec tests, establish domain names for the core entities." That instruction has not been validated against enough examples to be codified.
