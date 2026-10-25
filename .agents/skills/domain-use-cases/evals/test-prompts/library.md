# Eval: Library Management

**Prompt to run against the skill:**

> I need to gather requirements for a library management system. Start with what happens when a member borrows a book and work outward from there. I want domain use-cases and concept documents — no API design yet.

**What good output looks like:**

- Actors include at minimum: Member, Librarian
- Domain objects include: Book (or Copy — the distinction matters), Loan, Member account
- Domain events include: LoanCreated, BookReturned, LoanOverdue, ReservationPlaced
- Lifecycle states: `available`, `on loan`, `reserved`, `lost` for a copy; `active`, `overdue`, `returned` for a loan
- No HTTP or database language in main flow steps
- The Book/Copy distinction is surfaced: a Book is a title; a Copy is a physical item you can lend — these are two objects with different lifecycles
- Concept documents produced with OKF frontmatter and correct kind tags

**Known hard parts:**

- Book vs Copy: the skill should ask enough to discover that a library has one Book (title) but many Copies — the Copy is what gets loaned, not the Book
- Reservation vs Loan: a member can reserve a book that's currently on loan — that's a separate domain object
- Overdue: is overdue a state or an event? The interview should surface this
