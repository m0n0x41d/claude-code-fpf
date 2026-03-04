---
name: fpf
description: "Query the First Principles Framework (FPF) for reasoning guidance — principles, methods, templates, and terminology."
argument-hint: "[concept or question]"
---

# FPF — First Principles Framework Reference

You have access to the full FPF specification via a search tool. Use it when you need precise definitions, procedures, or principles from the framework.

## When to search

- You encounter FPF terminology you're unsure about
- You need a specific template, procedure, or principle
- The user references FPF concepts in their request
- You need to apply principled reasoning (problem framing, variant generation, evidence, etc.)

## How to search

Run the binary at the path shown below. It returns relevant FPF sections as markdown.

```bash
# Snippet search (default)
~/.claude/skills/fpf/references/fpf-rag search "<query>"

# Full sections (more context)
~/.claude/skills/fpf/references/fpf-rag search "<query>" --full

# Limit results
~/.claude/skills/fpf/references/fpf-rag search "<query>" --limit 5

# Get a specific section by heading
~/.claude/skills/fpf/references/fpf-rag section "<heading>"

# Version info
~/.claude/skills/fpf/references/fpf-rag info
```

## Inline index — key concepts and search terms

Use this index to decide WHEN to search. If a concept below is relevant, run a search for deeper context.

### Core reasoning cycle
- **ADI cycle** — Abduction → Deduction → Induction. Frame hypotheses, derive predictions, test against evidence.
- **Lifecycle stages** — Explore → Shape → Evidence → Operate. Always state current stage.

### Principles
- **WLNK (Weak Link)** — System reliability = min(component reliabilities). The weakest link bounds the whole.
- **MONO (Monotonicity)** — Adding components adds weak links. Benefit must justify new vulnerabilities.
- **BLP (Better Learning Policy)** — At comparable budgets, prefer methods with better scaling slopes.
- **NQD (Name-Quality-Duration)** — Never collapse to single score. Q references multi-dimensional indicators.
- **F-G-R** — Formality (ordinal, min), ClaimScope (set-valued, NOT ordinal), Reliability ([0,1], min).
- **CL 0-3** — Commensurability levels: 0=Opposed, 1=Comparable, 2=Translatable, 3=Near-identity.
- **Parity** — Fair comparison requires equal conditions. Use Parity Plan before comparing variants.
- **Anti-Goodhart** — Distinguish observation indicators, acceptance criteria, and optimization targets (1-3 max).
- **E/E policy** — Explore for unfamiliar; exploit for known. Default: explore. Preserve stepping stones.

### Problem design
- **Problem card (PROB-*)** — ≥3 hypotheses, trade-off axes, goldilocks assessment, acceptance spec
- **Anomaly record (ANOM-*)** — Observations, hypotheses, test plan
- **Characterization (CHR-*)** — Characteristic space, indicators, comparison rules
- **Problem portfolio (PPORT-*)** — Multiple problems with selection rules and diversification

### Solution design
- **SoTA survey (SOTA-*)** — ≥2 traditions, bridge matrix with CL 0-3, gaps
- **Strategy card (STRAT-*)** — Method family bet, invalidation conditions, variant generation axes
- **Solution portfolio (SPORT-*)** — ≥3 genuinely distinct variants, NQD per indicators, stepping stones
- **Selection (SEL-*)** — Pareto analysis, policy stated before applying, stepping-stone bets
- **Parity plan (PAR-*)** — Equal conditions for fair comparison

### Evidence and decisions
- **Evidence record (EVID-*)** — Predictions before testing, commands + outputs, F-G-R assessment
- **Decision record (DRR-*)** — Options, rationale, risks, rollback plan for irreversible decisions

### Key distinctions
- Plan ≠ reality
- Object ≠ description ≠ carrier
- "Process" → resolve to: Role | Capability | Method | Work | WorkPlan
- Design-time (Plan & Model) vs Run-time (Actions & Observations)

### Methodology
- **Coupled double-loop factories** — Problem factory (creative: problematization) + Solution factory (creative: strategizing + variants) + Factory of factories (meta)
- **Task tiers** — T1 (trivial) → T2 (localized) → T3 (substantive) → T4 (architectural)
- **Constraints** — Creative (C1-C4), Assurance (C5-C8), Session (C9-C10)
