// Dark mode functionality
const toggleSwitch = document.querySelector('.theme-switch input[type="checkbox"]');

function switchTheme(e) {
    if (e.target.checked) {
        document.documentElement.setAttribute('data-theme', 'dark');
        localStorage.setItem('theme', 'dark');
    } else {
        document.documentElement.setAttribute('data-theme', 'light');
        localStorage.setItem('theme', 'light');
    }
}

toggleSwitch.addEventListener('change', switchTheme);

// Load saved theme preference
const currentTheme = localStorage.getItem('theme');
if (currentTheme) {
    document.documentElement.setAttribute('data-theme', currentTheme);
    if (currentTheme === 'dark') {
        toggleSwitch.checked = true;
    }
}

// Feature collapsing functionality
document.querySelectorAll('.feature-header').forEach(header => {
    header.addEventListener('click', () => {
        const content = header.nextElementSibling;
        const isActive = header.classList.contains('active');

        // Close all features
        document.querySelectorAll('.feature-header').forEach(h => h.classList.remove('active'));
        document.querySelectorAll('.feature-content').forEach(c => c.classList.remove('show'));

        // Open clicked feature if it wasn't active
        if (!isActive) {
            header.classList.add('active');
            content.classList.add('show');
        }
    });
});

// API calls functionality
const endpoints = {
    health: 'http://localhost:8086/health',
    version: 'http://localhost:8086/version',
    readfile: 'http://localhost:8086/readfile',
    readJson: 'http://localhost:8086/readjson',
    readexternal: 'https://jsonplaceholder.typicode.com/posts/1/comments'
};

async function callEndpoint(endpoint) {
    try {
        const response = await fetch(endpoints[endpoint], {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            mode: 'cors'
        });
        
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        
        const data = await response.json();
        displayResults(data);
    } catch (error) {
        if (error.message.includes('Failed to fetch')) {
            displayResults({ 
                error: 'Server connection failed. Please ensure that:\n' +
                      '1. The server is running\n' +
                      '2. Server is accessible at ' + endpoints[endpoint] + '\n' +
                      '3. CORS is properly configured on the server'
            });
        } else {
            displayResults({ error: error.message });
        }
    }
}

function displayResults(data) {
    const tbody = document.getElementById('results-body');
    tbody.innerHTML = '';

    if (Array.isArray(data)) {
        // Handle array response
        data.forEach((item, index) => {
            Object.entries(item).forEach(([key, value]) => {
                const row = tbody.insertRow();
                row.insertCell(0).textContent = `[${index}] ${key}`;
                row.insertCell(1).textContent = value;
            });
        });
    } else if (typeof data === 'object') {
        // Handle object response
        Object.entries(data).forEach(([key, value]) => {
            const row = tbody.insertRow();
            row.insertCell(0).textContent = key;
            row.insertCell(1).textContent = typeof value === 'object' ? JSON.stringify(value) : value;
        });
    } else {
        // Handle primitive response
        const row = tbody.insertRow();
        row.insertCell(0).textContent = 'Result';
        row.insertCell(1).textContent = data;
    }
}
