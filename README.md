# LeetCode Solutions

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![LeetCode](https://img.shields.io/badge/LeetCode-FFA116?style=for-the-badge&logo=leetcode&logoColor=black)
![Auto Sync](https://img.shields.io/badge/auto--sync-GitHub_Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)

My accepted LeetCode solutions, mostly in **Go**. This repository is kept up to date automatically: a scheduled GitHub Action pulls every new *Accepted* submission and commits it here.

## How it works

Solutions are synced via the [`joshcai/leetcode-sync`](https://github.com/joshcai/leetcode-sync) GitHub Action. On each run it fetches accepted submissions from my LeetCode account and creates one folder per problem (`<id>-<slug>/`) containing the solution and the problem statement.

- **Language:** primarily Go (algorithms practice for backend interviews)
- **Sync:** daily + on manual trigger — see [`.github/workflows/sync.yml`](.github/workflows/sync.yml)

## Structure

```
<problem-id>-<problem-slug>/
  ├── <slug>.go        # accepted solution
  └── README.md        # problem statement
```

## Why this repo

Part of my structured preparation for backend internships — a transparent, always-current log of the problems I solve. Each solution is code I wrote and got accepted, not copied.
