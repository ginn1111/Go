# Pomodoro cli app

---

## Requirements

### Functional

1. Listing the pomodoro sections
2. Start a section with a specific minute (default: 30m)
3. Stop a section
4. Count done sections, minutes in a day

### Non-functional

1. Notification to system when the the section is done
2. Render a pomodoro clock, show a random quote

---

## Define commands

- root: po
  - start --minutes | -m (default: 30m)
  - stop
  - list
  - count
