# 🐣 Pocket Oni – Terminal Tamagotchi

A simple, terminal-based virtual pet inspired by classic Tamagotchi games.  
Take care of your creature by feeding it, playing with it, and keeping its stats balanced.  

![oni](internal/characters/oni_happy.gif)


Watch your pet's **Happiness**, **Hunger**, **Energy**, and **Hygiene** as they change over time, and make sure it stays healthy!

---

## 🖥️ Features

- Terminal-based interface (no GUI required)  
- Manage your pet's stats in real-time  
- Perform actions like feed, play, sleep, and clean  
- Save and load your progress  

---

## 🛠️ Installation

### Requirements

- Go 1.20+  
- Terminal that supports ANSI characters  

### Clone & Run

```bash
git clone github.com/Darkpowercross/pocket-oni.git
cd pocket-oni
go run .


Build Executable

```bash
go build -o pocket-oni ./cmd
./pocket-oni