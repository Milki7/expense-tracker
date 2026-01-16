const API_URL = 'http://localhost:8080/api/v1/expenses';
const form = document.getElementById('expense-form');
const expenseList = document.getElementById('expense-list');
const totalDisplay = document.getElementById('total-spending');

// --- READ: Fetch and display data ---
async function loadDashboard() {
    try {
        // 1. Fetch Expenses
        const res = await fetch(API_URL);
        const data = await res.json();
        
        // 2. Fetch Summary
        const sumRes = await fetch(`${API_URL}/summary`);
        const sumData = await sumRes.json();

        displayExpenses(data);
        totalDisplay.innerText = `Total Spending: $${sumData.total_spending.toFixed(2)}`;
    } catch (err) {
        console.error("Failed to load data:", err);
    }
}

function displayExpenses(expenses) {
    expenseList.innerHTML = '';
    expenses.forEach(exp => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${new Date(exp.CreatedAt).toLocaleDateString()}</td>
            <td>${exp.title}</td>
            <td>${exp.category}</td>
            <td>$${exp.amount.toFixed(2)}</td>
            <td class="actions">
                <button class="edit-btn" onclick="prepareEdit(${exp.ID}, '${exp.title}', ${exp.amount}, '${exp.category}')">Edit</button>
                <button class="delete-btn" onclick="deleteExpense(${exp.ID})">Delete</button>
            </td>
        `;
        expenseList.appendChild(row);
    });
}

// --- CREATE & UPDATE: Submit logic ---
form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const id = document.getElementById('expense-id').value;
    const payload = {
        title: document.getElementById('title').value,
        amount: parseFloat(document.getElementById('amount').value),
        category: document.getElementById('category').value
    };

    let url = API_URL;
    let method = 'POST';

    if (id) {
        url = `${API_URL}/${id}`;
        method = 'PUT';
    }

    await fetch(url, {
        method: method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    });

    resetForm();
    loadDashboard();
});

// --- DELETE: Remove item ---
async function deleteExpense(id) {
    if (confirm('Are you sure?')) {
        await fetch(`${API_URL}/${id}`, { method: 'DELETE' });
        loadDashboard();
    }
}

// --- Helper Functions ---
function prepareEdit(id, title, amount, category) {
    document.getElementById('expense-id').value = id;
    document.getElementById('title').value = title;
    document.getElementById('amount').value = amount;
    document.getElementById('category').value = category;
    
    document.getElementById('form-title').innerText = "Edit Expense";
    document.getElementById('cancel-btn').style.display = 'inline-block';
}

function resetForm() {
    form.reset();
    document.getElementById('expense-id').value = '';
    document.getElementById('form-title').innerText = "Add New Expense";
    document.getElementById('cancel-btn').style.display = 'none';
}

// --- EXPORT: Trigger CSV Download ---
document.getElementById('export-btn').addEventListener('click', () => {
    
    window.location.href = 'http://localhost:8080/api/v1/expenses/export';
});

document.getElementById('search-btn').addEventListener('click', () => {
    const title = document.getElementById('search-title').value;
    const sortBy = document.getElementById('sort-by').value;
    
    // Construct URL with query params
    const queryUrl = `${API_URL}?title=${title}&sort=${sortBy}&order=desc`;
    
    fetch(queryUrl)
        .then(res => res.json())
        .then(data => displayExpenses(data));
});


document.getElementById('cancel-btn').onclick = resetForm;

// Initial Load
loadDashboard();