# LeetCode Solutions

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![LeetCode](https://img.shields.io/badge/LeetCode-FFA116?style=for-the-badge&logo=leetcode&logoColor=black)
![Auto Sync](https://img.shields.io/badge/auto--sync-GitHub_Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white)

My accepted LeetCode solutions, mostly in **Go**. This repository is kept up to date automatically: a scheduled GitHub Action pulls every new *Accepted* submission and commits it here.

## Structure

```
<problem-id>-<problem-slug>/
  ├── <slug>.go        # accepted solution
  └── README.md        # problem statement
```

## Note — how the auto-sync is set up

Solutions are synced with the [`joshcai/leetcode-sync`](https://github.com/joshcai/leetcode-sync) GitHub Action (see [`.github/workflows/sync.yml`](.github/workflows/sync.yml)). If you want the same for your own repo, you need to give the Action your LeetCode session via two repository **secrets**.

**1. Grab the cookies.** Log in to [leetcode.com](https://leetcode.com), open DevTools (`F12`) → **Application** tab → **Cookies → `https://leetcode.com`**, and copy these two values:

| Cookie in the browser | GitHub secret name to create |
|---|---|
| `LEETCODE_SESSION` | `LEETCODE_SESSION` |
| `csrftoken` | `LEETCODE_CSRF_TOKEN` |

**2. Add the secrets.** Repo → **Settings → Secrets and variables → Actions → New repository secret** — create the two secrets above with the copied values.

**3. Run it.** **Actions** tab → **Sync LeetCode** → **Run workflow**. The first run backfills all past *Accepted* submissions; after that it runs daily on a schedule.

> **Security:** these cookies are account access — keep them in GitHub **Secrets** only, never commit them to the repo. `LEETCODE_SESSION` expires roughly every two weeks; if the Action starts failing, refresh that secret with a fresh cookie.

