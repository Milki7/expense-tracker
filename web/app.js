const API_URL = 'http://localhost:8080/api/v1/expenses';
const form = document.getElementById('expense-form');
const expenseList = document.getElementById('expense-list');
const totalDisplay = document.getElementById('total-spending');
let myChart; // Stores the Chart.js instance

/**
 * --- DATA LOADING & DASHBOARD UPDATE ---
 */
async function loadDashboard(query = '') {
    try {
        // 1. Fetch Expenses (with optional search/sort query)
        const res = await fetch(`${API_URL}${query}`);
        const expenses = await res.json();
        
        // 2. Fetch Total Summary
        const sumRes = await fetch(`${API_URL}/summary`);
        const sumData = await sumRes.json();

        // 3. Update UI Components
        displayExpenses(expenses);
        totalDisplay.innerText = `$${sumData.total_spending.toLocaleString(undefined, {minimumFractionDigits: 2})}`;
        
        // 4. Refresh the Chart
        updateChart();
    } catch (err) {
        console.error("Dashboard load failed:", err);
    }
}

/**
 * --- TABLE RENDERING ---
 */
function displayExpenses(expenses) {
    expenseList.innerHTML = '';
    
    if (expenses.length === 0) {
        expenseList.innerHTML = '<tr><td colspan="5" style="text-align:center;">No expenses found.</td></tr>';
        return;
    }

    expenses.forEach(exp => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${new Date(exp.CreatedAt).toLocaleDateString()}</td>
            <td><strong>${exp.title}</strong></td>
            <td><span class="badge">${exp.category}</span></td>
            <td style="font-family: monospace; font-weight: bold;">$${exp.amount.toFixed(2)}</td>
            <td style="text-align: right;">
                <button class="btn-small" onclick="prepareEdit(${exp.ID}, '${exp.title}', ${exp.amount}, '${exp.category}')">Edit</button>
                <button class="btn-small btn-cancel" onclick="deleteExpense(${exp.ID})">Del</button>
            </td>
        `;
        expenseList.appendChild(row);
    });
}

/**
 * --- CREATE & UPDATE ---
 */
form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const id = document.getElementById('expense-id').value;
    const payload = {
        title: document.getElementById('title').value,
        amount: parseFloat(document.getElementById('amount').value),
        category: document.getElementById('category').value
    };

    const method = id ? 'PUT' : 'POST';
    const url = id ? `${API_URL}/${id}` : API_URL;

    try {
        const response = await fetch(url, {
            method: method,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        if (response.ok) {
            resetForm();
            loadDashboard();
        } else {
            const error = await response.json();
            alert("Error: " + error.error);
        }
    } catch (err) {
        alert("Server connection failed.");
    }
});

/**
 * --- DELETE ---
 */
async function deleteExpense(id) {
    if (confirm('Are you sure you want to delete this expense?')) {
        await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
        loadDashboard();
    }
}

/**
 * --- CHART.JS LOGIC ---
 */
async function updateChart() {
    try {
        const res = await fetch(`${API_URL}/categories`);
        const stats = await res.json();

        const labels = stats.map(item => item.category);
        const totals = stats.map(item => item.total);

        const ctx = document.getElementById('categoryChart').getContext('2d');

        if (myChart) { myChart.destroy(); }

        myChart = new Chart(ctx, {
            type: 'doughnut',
            data: {
                labels: labels,
                datasets: [{
                    data: totals,
                    backgroundColor: ['#3498db', '#e74c3c', '#f1c40f', '#2ecc71', '#9b59b6', '#34495e'],
                    hoverOffset: 10
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: { legend: { position: 'bottom' } }
            }
        });
    } catch (err) {
        console.warn("Chart data unavailable");
    }
}

/**
 * --- SEARCH & SORT & EXPORT ---
 */
document.getElementById('search-btn').addEventListener('click', () => {
    const title = document.getElementById('search-title').value;
    const sortBy = document.getElementById('sort-by').value;
    // Constructs query like: ?title=coffee&sort=amount&order=desc
    loadDashboard(`?title=${title}&sort=${sortBy}&order=desc`);
});

document.getElementById('export-btn').onclick = () => {
    window.location.href = `${API_URL}/export`;
};

/**
 * --- UI HELPERS ---
 */
function prepareEdit(id, title, amount, category) {
    document.getElementById('expense-id').value = id;
    document.getElementById('title').value = title;
    document.getElementById('amount').value = amount;
    document.getElementById('category').value = category;
    
    document.getElementById('form-title').innerText = "Edit Expense";
    document.getElementById('submit-btn').innerText = "Update";
    document.getElementById('cancel-btn').style.display = 'inline-block';
    window.scrollTo({ top: 0, behavior: 'smooth' });
}

function resetForm() {
    form.reset();
    document.getElementById('expense-id').value = '';
    document.getElementById('form-title').innerText = "Add Expense";
    document.getElementById('submit-btn').innerText = "Save";
    document.getElementById('cancel-btn').style.display = 'none';
}

document.getElementById('cancel-btn').onclick = resetForm;

// INITIAL LAUNCH
loadDashboard();