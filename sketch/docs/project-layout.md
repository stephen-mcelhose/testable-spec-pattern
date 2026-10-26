# Project Layout

```
.
├── .agents/
│   └── skills/
│       ├── requirements-gathering/
│       │   ├── domain-use-case-interview/     ← exists (needs move)
│       │   └── domain-use-cases/              ← exists (needs move)
│       │
│       ├── contracts-and-domain-modeling/
│       │   ├── model-as-code/                 ← exists (needs move)
│       │   └── resource-oriented-design/      ← exists (needs move)
│       │
│       ├── specifications/
│       │   └── rfc-2119-spec-writer/          ← exists (needs move)
│       │
│       ├── planning/
│       │   └── (none yet)
│       │
│       ├── implementation/
│       │   └── go-project-scaffold/           ← exists (needs move)
│       │
│       └── testing/
│           └── go-behavioral-spec/            ← exists (needs move)
│
└── sketch/
    └── docs/
        ├── index.md                           ← Table of Contents
        ├── project-layout.md                  ← this file
        │
        ├── idea/
        │   └── index.md
        │
        ├── requirements-gathering/
        │   ├── index.md
        │   ├── use-cases.md
        │   ├── domain-types.md
        │   └── gherkin-scenarios.md
        │
        ├── contracts-and-domain-modeling/
        │   ├── index.md
        │   ├── contracts.md
        │   └── domain-model.md
        │
        ├── specifications/
        │   ├── index.md
        │   └── specification.md
        │
        ├── planning/
        │   ├── index.md
        │   └── plan.md
        │
        ├── implementation/
        │   └── index.md
        │
        └── testing/
            ├── index.md
            └── runnable-specification.md
```

## Skills not mapped to a lifecycle stage

These exist but sit outside the `<lifecycle_stage>/` convention:

| Skill               | Current location              | Note                              |
| ------------------- | ----------------------------- | --------------------------------- |
| `docs-wiki`         | `.agents/skills/docs-wiki/`   | Meta — doc management, not a lifecycle stage |
| `write-issue`       | `.agents/skills/write-issue/` | Meta — process tooling            |
