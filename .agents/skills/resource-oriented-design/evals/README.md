# Evals for `resource-oriented-design`

Two independent eval tracks. Run them separately — they measure different things.

---

## Track 1 — Quality evals (`evals.json`)

**What it measures:** Does the skill produce an AIP-correct proto when loaded?

Each eval in `evals.json` has a `prompt` (a design task) and `expectations` (specific proto assertions). The grader reads a run transcript and checks each expectation pass/fail.

### Running manually (no automation yet)

For each eval:

1. **Execute** — Spawn a subagent with the skill loaded, give it the eval's `prompt`. Ask it to produce a complete `.proto` snippet. Save the response.

2. **Grade** — Spawn the skill-creator grader (`grader.md`) with:
   - `expectations`: the eval's `expectations` list
   - `transcript_path`: the executor's saved output
   - `outputs_dir`: directory with the produced proto

3. **Aggregate** — Collect grading.json files into a `benchmark.json` (schema in skill-creator `references/schemas.md`).

### Running with skill-creator (automated)

```bash
# From the skill-creator workspace, point at this skill:
python scripts/run_loop.py \
  --skill-path path/to/resource-oriented-design \
  --eval-set evals/evals.json \
  --mode test
```

The loop spawns executor and grader subagents for each eval automatically.

### Interpreting results

- A **high pass rate** with the skill loaded vs without it is the signal. Run evals in both configurations to get the delta.
- Individual expectation failures identify specific AIP rules the skill isn't teaching clearly enough.
- The grader's `eval_feedback` section surfaces suggestions for improving the evals themselves.

### Adding new evals

1. Add an entry to `evals.json` with a sequential integer `id`, a narrow `prompt` (one AIP topic), an `expected_output` description, and `expectations` as verifiable strings.
2. Tag each expectation with the AIP rule it maps to and its RFC 2119 normative level (MUST / SHOULD / MUST NOT).
3. Only include expectations that are observable in proto output — no runtime behavior.

---

## Track 2 — Trigger evals (`trigger-eval.json`)

**What it measures:** Does the skill's *description* cause the agent to read the skill when it should?

Each entry is `{ "query": "...", "should_trigger": true|false }`. True = the skill should be invoked; false = it should not.

### Running

```bash
# From the skill-creator workspace:
python scripts/run_eval.py \
  --skill-path path/to/resource-oriented-design \
  --eval-set evals/trigger-eval.json \
  --runs-per-query 3 \
  --verbose
```

### Optimizing the description

If trigger evals fail (skill not firing on should-trigger queries, or firing on should-not queries), run the optimization loop:

```bash
python scripts/run_loop.py \
  --skill-path path/to/resource-oriented-design \
  --eval-set evals/trigger-eval.json \
  --mode optimize-description
```

The loop proposes description variants and measures trigger rates.

### Adding trigger evals

- **should_trigger: true** — Real design tasks: "design an X API", "is this AIP-compliant", "CRUD methods for resource Y"
- **should_trigger: false** — Adjacent but distinct tasks: "write Go tests", "set up gRPC server", "parse JSON", "handle auth"
- Keep false cases realistic — they should be things a user might actually ask in the same session.

---

## Skill iteration workflow

```
1. Run quality evals → identify failing expectations
2. Read failing expectations → find which SKILL.md section is unclear or missing
3. Edit SKILL.md
4. Re-run quality evals → confirm pass rate improves
5. Run trigger evals → confirm description still fires correctly
6. Commit with scorecard note (v1/SCORECARD-iterN.md)
```

The quality evals drive content. The trigger evals guard the description. Don't conflate them.
