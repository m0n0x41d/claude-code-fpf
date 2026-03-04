---
name: fpf
description: "Apply First Principles Framework (FPF) reasoning — principled problem design, variant generation, and evidence-based decisions."
argument-hint: "[task, concept, or question]"
---

# FPF — First Principles Framework

FPF is a systems thinking methodology by Anatoly Levenchuk. This skill gives you its operational core — apply it to reason about problems, solutions, and decisions. Use the RAG search for deeper context when needed.

---

## Core thinking algorithm

When asked to apply FPF or when facing a non-trivial decision, follow this loop:

### 1. Frame the problem BEFORE solving it

The bottleneck is **problem quality**, not solution speed. Before generating any solution:

- **State what's anomalous** — what observation doesn't fit the current model?
- **Generate ≥3 hypotheses** — genuinely distinct explanations, not variations of one
- **Identify trade-off axes** — what dimensions are in tension? (speed vs safety, generality vs performance, etc.)
- **Define acceptance criteria** — how will you know the problem is solved? Separate:
  - **Optimization targets** (1-3 max) — what you're actively improving
  - **Acceptance criteria** — hard constraints that must hold
  - **Observation indicators** — things you monitor but don't optimize (Anti-Goodhart: monitoring what you don't optimize prevents reward hacking)

### 2. Characterize before comparing

Before evaluating anything, define the **characteristic space** — what dimensions matter and how they're measured. Without this, comparisons are arbitrary.

### 3. Generate genuinely distinct variants

- Produce **≥3 variants** that differ in **kind**, not degree
- For each variant, assess quality as **multi-dimensional** (NQD) — never collapse to a single score
- Identify each variant's **weakest link** (WLNK) — the component that bounds overall quality
- Preserve **1-2 stepping stones** — variants that open future possibilities even if not optimal now

### 4. Select from the Pareto front

- **State the selection policy BEFORE applying it** — what matters most and why
- **Ensure parity** — fair comparison requires equal conditions (same inputs, same constraints)
- Hold the **Pareto front** — don't discard non-dominated options prematurely
- When a variant adds complexity over a simpler one, the added components must justify the new weak links they introduce (**MONO**)
- At comparable budgets, prefer methods with **better scaling slopes** over hand-tuned solutions (**BLP**)

### 5. Test against reality

- **Predict before testing** — state what you expect to observe if the hypothesis is correct AND if it's wrong
- **Record evidence** — commands run, outputs observed, interpretation
- **Assess confidence** — using F-G-R:
  - **F** (Formality): how rigorous is the method? (ordinal, min across chain)
  - **G** (ClaimScope): what exactly does the claim cover? (set-valued, NOT ordinal)
  - **R** (Reliability): how likely is the claim true? ([0,1], min across chain)
- **Close the loop** — evidence either corroborates or refutes. If refuted, update the problem framing and iterate.

---

## Reasoning cycle (ADI)

All thinking follows: **Abduction → Deduction → Induction**

| Phase | What happens | Output |
|-------|-------------|--------|
| **Abduction** | Generate hypotheses, frame problems, propose explanations | Problem cards, anomaly records, candidate hypotheses |
| **Deduction** | Derive predictions, define what MUST follow if hypothesis is true | Falsifiable predictions, acceptance specs, logical consequences |
| **Induction** | Test predictions against evidence, update confidence | Evidence records, corroboration/refutation, confidence update |

**Anti-patterns to avoid:**
- Jumping to solutions without framing the problem (skipping abduction)
- Testing without predictions (skipping deduction — "data dredging")
- Claiming "verified" without recorded evidence (skipping induction)

---

## Lifecycle stages

Every artifact progresses: **Explore → Shape → Evidence → Operate**

| Stage | Activity | ADI phase |
|-------|----------|-----------|
| **Explore** | Generate possibilities, brainstorm, question assumptions | Abduction |
| **Shape** | Select direction, define architecture, ensure internal consistency | Deduction |
| **Evidence** | Test against reality, validate claims, measure performance | Induction |
| **Operate** | Deploy, monitor, maintain | Continuous induction |

Always state which stage you're in. Don't skip stages.

---

## Core invariants (the rules that always hold)

- **WLNK** — System quality = min(component qualities). The weakest link bounds the whole. Always identify it.
- **MONO** — Improving a part cannot worsen the whole. Adding a part adds a new potential weak link — justify the cost.
- **IDEM** — Evaluating a single element in isolation must return that element unchanged (no accidental upgrade/downgrade).
- **COMM/LOC** — For independent components, evaluation order and location don't matter. When dependencies exist, order matters and must be controlled.

---

## Key distinctions (always maintain these)

- **Plan ≠ Reality** — a model is not the thing it models
- **Object ≠ Description ≠ Carrier** — the system, its spec, and its implementation are three different things
- **Resolve "process"** — always disambiguate into: Role | Capability | Method | Work | WorkPlan
- **Design-time vs Run-time** — planning and modeling vs acting and observing
- **Commensurability (CL 0-3)** — before comparing two things, assess how comparable they are:
  - 0 = Opposed (contradictory frames)
  - 1 = Comparable (same domain, different frameworks)
  - 2 = Translatable (systematic mapping exists)
  - 3 = Near-identity (same framework, minor differences)

---

## When to search the RAG

The above is enough for applying FPF reasoning. Search `fpf-rag` when you need:

- **Specific templates** — exact format for problem cards, evidence records, decision records, etc.
- **Deep definitions** — formal specification of a concept beyond the summary above
- **Conformance checklists** — detailed rules for a specific FPF pattern
- **Aggregation rules** — how to compose assessments across components (Γ flavours)
- **Specific patterns** — A.* (concept patterns) or B.* (process patterns) by number

```bash
# Quick search
~/.claude/skills/fpf/references/fpf-rag search "<query>"

# Full section content
~/.claude/skills/fpf/references/fpf-rag search "<query>" --full

# Specific section by heading
~/.claude/skills/fpf/references/fpf-rag section "<heading>"
```

---

## Concept index (search terms)

**Problem design:** problem card, PROB, anomaly, ANOM, characterization, CHR, problem portfolio, PPORT, goldilocks, trade-off axes, acceptance spec

**Solution design:** SoTA survey, SOTA, strategy card, STRAT, method family, invalidation conditions, solution portfolio, SPORT, variant generation, NQD, stepping stones

**Selection:** Pareto front, selection policy, SEL, parity plan, PAR, fair comparison, Pareto analysis

**Evidence:** evidence record, EVID, predictions, corroboration, refutation, F-G-R, assurance level, L0, L1, L2

**Decisions:** decision record, DRR, irreversible, rollback plan, options, rationale

**Aggregation:** Gamma, fold, Quintet, IDEM, COMM, LOC, WLNK, MONO, weakest link, cutset

**Reasoning:** ADI cycle, abduction, deduction, induction, explore, shape, evidence, operate, lifecycle

**Comparison:** commensurability, CL 0-3, bridge matrix, translation, near-identity, opposed
