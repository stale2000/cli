# Session Context

## User Prompts

### Prompt 1

let's add in tests to benchmark against main to show differences. Let's add it in our mise.toml file

### Prompt 2

fix these issues:

Similar to the SeedShadowBranch benchmark, this benchmark appends to the same metadata branch on each iteration, causing performance to degrade as the branch grows. This may not accurately reflect the typical performance of the operation.

Consider creating a fresh BenchRepo inside b.Loop() or using b.Run() with different checkpoint counts as sub-benchmarks to get more meaningful performance data.



The bench:compare EXIT trap only removes the temp directory but doesn't resto...

