---
name: resource-oriented-design
description: Guide API design using Resource Oriented Design (ROD). Use when designing a new API, reviewing a proto file, naming resources or methods, or deciding whether to use a standard vs custom method. Trigger on "design this API", "how should I name this resource", "ROD review", "AIP compliant", "resource-oriented design".
license: MIT
allowed-tools: Read
---

# Resource Oriented Design

Design APIs as a hierarchy of named **resources** acted on by a small, stable set of standard methods over a stateless protocol. Follow [AIP-121](https://google.aip.dev/121) as the canonical reference.

## Design order

Work through these in sequence before writing any method signatures:

1. What **resources** (nouns) does the API expose?
2. What **relationships and hierarchy** exist between them?
3. What is the **schema** of each resource?
4. What **methods** does each resource need — and can standard methods cover them?

## Standard methods

Prefer standard methods. Custom methods are for operations that don't cleanly map to any standard verb — not a general escape hatch.

| Method | HTTP   | Request body | Response body |
| ------ | ------ | ------------ | ------------- |
| List   | GET    | None         | Resource list |
| Get    | GET    | None         | Resource      |
| Create | POST   | Resource     | Resource      |
| Update | PATCH  | Resource     | Resource      |
| Delete | DELETE | None         | None          |

## Key AIPs

| AIP                                   | What it covers                              |
| ------------------------------------- | ------------------------------------------- |
| [AIP-121](https://google.aip.dev/121) | Resource-oriented design — foundational philosophy |
| [AIP-122](https://google.aip.dev/122) | Resource names — `{collection}/{id}` patterns |
| [AIP-123](https://google.aip.dev/123) | Resource types — `{Service}/{Type}` annotation |
| [AIP-131](https://google.aip.dev/131) | Standard Get                                |
| [AIP-132](https://google.aip.dev/132) | Standard List                               |
| [AIP-133](https://google.aip.dev/133) | Standard Create                             |
| [AIP-134](https://google.aip.dev/134) | Standard Update                             |
| [AIP-135](https://google.aip.dev/135) | Standard Delete                             |
| [AIP-154](https://google.aip.dev/154) | Resource freshness validation (ETags)       |
| [AIP-158](https://google.aip.dev/158) | Pagination                                  |

For AIPs not listed above, consult the full index at **https://google.aip.dev/**.
