# Personal Website — Gin (Go)

A toy/personal website built by **Louis Lizzadro** to learn web development, Go/Gin, and Claude Code. Built **in parallel** with a Python/Flask twin repo, [PersonalWebsiteFlask](https://github.com/lllizzadro/PersonalWebsiteFlask) — every feature is built in Flask **first**, then mirrored here in Gin.

## About Louis
- Software developer learning web dev. GitHub: `lllizzadro`.
- **Hands-on learner — this matters.** Guide him to write the code himself: explain the concept and what's needed, let him implement it, then review what he wrote and explain any mistakes. Only write code directly when demonstrating a brand-new concept for the first time, or when he's stuck after trying.

## Stack & conventions
- **Gin + Go `html/template`.** Templates in `templates/`, loaded via `router.LoadHTMLGlob("templates/*")`. Uses **composition, not inheritance**: `components.html` defines `head`/`nav`/`footer`; each page wraps its content in `{{define "name"}}...{{end}}` and pulls components in with `{{template "nav" .}}`.
- **Tailwind CSS via CDN** — no build step.
- **SQLite** via `database/sql` + **`modernc.org/sqlite`** (pure Go, no CGO — important on Windows). Driver name is **`"sqlite"`**, NOT `"sqlite3"`. One shared `*sql.DB` pool (thread-safe) opened once in `initDB()` — unlike Flask, do NOT open per-request.
- Read rows with `rows.Next()` + `rows.Scan(&...)` into a struct. Template-accessed struct fields must be **exported** (capitalized).
- **Always use parameterized queries** (`?` placeholders).
- State-changing forms use **POST/Redirect/GET** (`c.Redirect(http.StatusSeeOther, "...")`).
- Static files served via `router.Static("/static", "./static")`; reference directly as `/static/js/dice.js` (no `url_for` helper like Flask has).

## Run
```
go run main.go       # http://localhost:8080
```

## Progress
- ✅ Home (`/`) — bio + projects list
- ✅ Guestbook (`/guestbook`) — SQLite-backed, form + POST/Redirect/GET
- ✅ Dice roller (`/dice`) — client-side JS in `static/js/dice.js` (identical to the Flask version)
- ⏭️ **Next: styling pass** — convert the guestbook `<table>` to mobile-friendly cards, learn Tailwind responsive breakpoints (`sm:`/`md:`/`lg:`), and style the form inputs.
