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

// API endpoints
const endpoints = {
    health: 'http://localhost:8086/health',
    version: 'http://localhost:8086/version',
    readfile: 'http://localhost:8086/readfile',
    readJson: 'http://localhost:8086/readjson',
    readexternal: 'https://jsonplaceholder.typicode.com/posts/1/comments',
    agentList: 'http://localhost:8086/fd/v1/list/agent'
};

// Toggle filter inputs visibility
function toggleFilters(radio) {
    const filterInputs = document.querySelectorAll('.filter-inputs');
    filterInputs.forEach(row => {
        row.style.display = radio.value === 'filter' ? 'table-row' : 'none';
    });
}

// Fetch agent list
async function fetchAgentList() {
    try {
        const queryType = document.querySelector('input[name="queryType"]:checked').value;
        const requestBody = {
            "FML_USER_ID": "SYSTEM",
            "FML_SSSN_ID": "SESSION001",
            "FML_FROM_DT": "2025-01-01",
            "FML_TO_DT": "2025-12-31"
        };

        if (queryType === 'filter') {
            requestBody.FML_AGENT_CD = document.getElementById('agentCode').value || "";
            requestBody.FML_AGENT_BRANCH = document.getElementById('agentBranch').value || "";
            requestBody.FML_AGENT_ACTIVE = document.getElementById('agentActive').value === 'true';
        } else {
            // Default values for full query
            requestBody.FML_AGENT_CD = "";
            requestBody.FML_AGENT_BRANCH = "";
            requestBody.FML_AGENT_ACTIVE = true;
        }

        requestBody.FML_AGENT_CATEGORY = 1;
        requestBody.FML_AGENT_EMP_TYPE = "ALL";

        const response = await fetch(endpoints.agentList, {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            mode: 'cors',
            body: JSON.stringify(requestBody)
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        displayAgentResults(data);
    } catch (error) {
        displayResults({ 
            error: `Failed to fetch agent list: ${error.message}\n` +
                  'Please ensure that:\n' +
                  '1. The server is running\n' +
                  '2. Server is accessible at ' + endpoints.agentList + '\n' +
                  '3. CORS is properly configured on the server'
        });
    }
}

// Display agent results in a table
function displayAgentResults(data) {
    const tbody = document.getElementById('results-body');
    tbody.innerHTML = '';
    
    if (data && data.data && Array.isArray(data.data)) {
        // Add refresh button if not exists
        let refreshButton = document.querySelector('.refresh-button');
        if (!refreshButton) {
            refreshButton = document.createElement('button');
            refreshButton.className = 'refresh-button';
            refreshButton.textContent = 'Refresh';
            refreshButton.onclick = fetchAgentList;
            document.getElementById('results-container').insertBefore(refreshButton, document.getElementById('results-table'));
        }

        // Create table headers
        const thead = document.querySelector('#results-table thead tr');
        thead.innerHTML = `
            <th>Agent Code</th>
            <th>Name</th>
            <th>Branch</th>
            <th>Category</th>
            <th>Status</th>
            <th>Last Update</th>
        `;

        // Add data rows
        data.data.forEach(agent => {
            const row = tbody.insertRow();
            row.insertCell(0).textContent = agent.agentCode;
            row.insertCell(1).textContent = `${agent.agentFirstName} ${agent.agentLastName}`;
            row.insertCell(2).textContent = agent.agentBranch;
            row.insertCell(3).textContent = agent.agentCategory;
            row.insertCell(4).textContent = agent.agentActive ? 'Active' : 'Inactive';
            row.insertCell(5).textContent = agent.dtUpdate;
        });
    } else {
        displayResults(data); // Use the default display function for error cases
    }
}

// Generic endpoint calls functionality
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
