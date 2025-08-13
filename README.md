# Scrum Poker

A real-time **Scrum Poker** application for Agile teams to estimate tasks collaboratively.  
Built with **Golang** (backend), **Next.js** (frontend), and **Docker Compose** for a fully containerized setup.

---

## 📌 Features (Current Scope)

- **Guest User Join** – No sign-up required, auto-generate a temporary guest ID.
- **Create & Join Rooms** – Share unique room links with your team.
- **In-Room Chat** – Communicate instantly while estimating.
- **WebSocket Support** – Live room and chat updates without refreshing.

> This is the **MVP scope** — voting, reveal, and persistence will be added in later milestones.

---

## 🛠 Tech Stack

| Layer     | Technology     | Purpose                         |
| --------- | -------------- | ------------------------------- |
| Frontend  | Next.js        | UI & client-side logic          |
| Backend   | Golang         | REST API & WebSocket handling   |
| Container | Docker Compose | Local development orchestration |
| Realtime  | WebSockets     | Room state & chat updates       |

---
