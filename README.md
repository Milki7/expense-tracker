# 💸 Go Expense Tracker Pro

A full-stack, lightweight personal finance management tool. This application allows you to track spending, categorize expenses, and visualize your financial habits through a real-time dashboard.

## 🚀 Features

### Backend (Golang & Gin)
* **RESTful API**: Clean CRUD operations for expense management.
* **Smart Search**: Filter by title or category via URL query parameters.
* **Dynamic Sorting**: Securely sort data by date, amount, or name.
* **Data Export**: Built-in CSV generator for financial reporting.
* **SQLite Persistence**: Uses GORM as an ORM for reliable data storage.

### Frontend (Vanilla JS & Chart.js)
* **Live Dashboard**: View total spending and recent transactions at a glance.
* **Visual Analytics**: Interactive doughnut charts showing spending distribution.
* **Responsive Design**: Modern UI built with Water.css that works on desktop and mobile.
* **AJAX Operations**: Add, edit, and delete expenses without page refreshes.

---

## 🛠️ Project Structure

```text
expense-tracker/
├── cmd/
│   └── api/main.go          # Application entry point
├── config/
│   └── database.go          # SQLite & GORM configuration
├── internals/
│   ├── handlers/            # HTTP request controllers
│   ├── models/              # Database structs (Schema)
│   ├── repositories/        # SQL queries and DB logic
│   └── services/            # Business logic and validation
├── web/
│   ├── index.html           # Dashboard UI
│   └── app.js               # Frontend logic & API calls
└── go.mod                   # Dependency management

🚦 Getting Started
1. Prerequisites

    Go (1.20 or higher)

    A browser (Chrome/Firefox recommended)

2. Installation
Bash

# Clone the repository
git clone <your-repo-url>

# Navigate to project
cd expense-tracker

# Install Go dependencies
go mod tidy

3. Run the App
Bash

go run cmd/api/main.go

The server will start at http://localhost:8080.
4. Access the UI

Open your browser and navigate to: http://localhost:8080/ui/index.html
📍 API Reference
Endpoint	Method	Description
/api/v1/expenses	GET	List/Search/Sort expenses
/api/v1/expenses	POST	Create a new expense
/api/v1/expenses/:id	PUT	Update an existing expense
/api/v1/expenses/:id	DELETE	Remove an expense
/api/v1/expenses/summary	GET	Get total spending sum
/api/v1/expenses/categories	GET	Get category breakdown
/api/v1/expenses/export	GET	Download CSV file
🧪 Testing with Query Params

You can test the search and sort logic directly in the browser or Postman:

    Sort by Price (High to Low): ?sort=amount&order=desc

    Search for Food: ?category=Food

    Combined: ?title=coffee&sort=amount&order=asc

🛡️ License

Distributed under the MIT License.


---

### How to use this file:
1.  In your project's root folder, create a new file named `README.md`.
2.  Paste the content above into that file.
3.  **Commit the change:**
    ```bash
    git add README.md
    git commit -m "docs: add professional README with project overview"
    ```