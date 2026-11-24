# Contains Duplicate

> NeetCode 150 · Arrays & Hashing · Blind 75 — a seen-set with early exit: the first repeat ends the scan. Go, stdlib only, tested.

## Problem

Does any value occur at least twice in the array? The curtain-raiser of the NeetCode list — one notch simpler than Find-the-Duplicate (dsa-010), which demanded O(1) space and no mutation.

## Approach

Insert each value into a set, answering true the moment an insert finds its value already present. The check-before-insert ordering gives the early exit; sorting first (O(n log n), then adjacent compare) trades speed for zero allocation, the alternative worth naming. `map[int]struct{}` rather than `map[int]bool` is the idiomatic Go set — the empty struct occupies zero bytes, stating "membership only, no payload" in the type.

**Complexity:** O(n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
