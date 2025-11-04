# expense-tracker

A lightweight command-line tool written in Go for managing personal expenses, tracking budgets, and generating summaries.

---

## Features

* Add, update, and delete expenses
* Categorize expenses
* View all expenses or filter by category
* Generate total or monthly summaries
* Set monthly budgets

---

## Installation

```bash
git clone https://github.com/mesh-dell/expense-tracker.git
cd expense-tracker
go build -o expense-tracker
```

## Usage

```
expense-tracker [command] [options]
```

### Commands

**add**

```
expense-tracker add --description "Lunch" --amount 20 --category "Food"
```

**update**

```
expense-tracker update --id 1 --description "Supper" --amount 25 --category "Food"
```

**delete**

```
expense-tracker delete --id 2
```

**list**

```
expense-tracker list --category "Luxuries"
```

**summary**

```
expense-tracker summary --month 8
```

**budget**

```
expense-tracker budget --month 12 --amount 250
```

**help**

```
expense-tracker help
```

---

## Example

```bash
expense-tracker add --description "Coffee" --amount 5 --category "Food"
expense-tracker list
expense-tracker summary --month 11
expense-tracker budget --month 11 --amount 300
```
