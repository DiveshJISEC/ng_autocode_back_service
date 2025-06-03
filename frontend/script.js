document.addEventListener('DOMContentLoaded', function() {
    // Form elements
    const agentListForm = document.getElementById('agentListForm');
    const orderBookForm = document.getElementById('orderBookForm');
    const healthCheck = document.getElementById('healthCheck');
    const refreshBtn = document.getElementById('refreshBtn');
    const backBtn = document.getElementById('backBtn');

    // Filter sections
    const agentListFilters = document.getElementById('agentListFilters');
    const orderBookFilters = document.getElementById('orderBookFilters');

    // Radio buttons
    document.querySelectorAll('input[name="agentListType"]').forEach(radio => {
        radio.addEventListener('change', (e) => {
            agentListFilters.style.display = e.target.value === 'filter' ? 'block' : 'none';
        });
    });

    document.querySelectorAll('input[name="orderBookType"]').forEach(radio => {
        radio.addEventListener('change', (e) => {
            orderBookFilters.style.display = e.target.value === 'filter' ? 'block' : 'none';
        });
    });

    // API endpoints
    const API_ENDPOINTS = {
        AGENT_LIST: 'http://localhost:8016/fd/v1/list/agent',
        ORDER_BOOK: 'http://localhost:8016/fd/v1/book/orderbook',
        HEALTH: 'http://localhost:8016/health'
    };

    // Store previous results for back button
    let resultHistory = [];
    let currentHistoryIndex = -1;

    // Helper function to make API calls
    async function makeApiCall(url, data = null) {
        try {
            const response = await fetch(url, {
                method: data ? 'POST' : 'GET',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: data ? JSON.stringify(data) : undefined
            });
            
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            
            const result = await response.json();
            return result;
        } catch (error) {
            console.error('API call failed:', error);
            throw error;
        }
    }

    // Function to create a table from JSON data
    function createTable(data) {
        if (!data || (!Array.isArray(data) && typeof data !== 'object')) {
            return '<div class="alert alert-info">No data available</div>';
        }

        const array = Array.isArray(data) ? data : [data];
        if (array.length === 0) {
            return '<div class="alert alert-info">No records found</div>';
        }

        const headers = Object.keys(array[0]);
        let table = '<table class="table table-striped table-hover">';
        
        // Table headers
        table += '<thead><tr>';
        headers.forEach(header => {
            table += `<th>${header}</th>`;
        });
        table += '</tr></thead>';

        // Table body
        table += '<tbody>';
        array.forEach(row => {
            table += '<tr>';
            headers.forEach(header => {
                table += `<td>${row[header] !== null ? row[header] : ''}</td>`;
            });
            table += '</tr>';
        });
        table += '</tbody></table>';

        return table;
    }

    // Function to update result display
    function updateResult(data) {
        const resultTable = document.getElementById('resultTable');
        resultTable.innerHTML = createTable(data);
        
        // Add to history
        resultHistory = resultHistory.slice(0, currentHistoryIndex + 1);
        resultHistory.push(data);
        currentHistoryIndex = resultHistory.length - 1;
        
        updateNavigationButtons();
    }

    // Update navigation buttons state
    function updateNavigationButtons() {
        backBtn.disabled = currentHistoryIndex <= 0;
        refreshBtn.disabled = !resultHistory.length;
    }

    // Event Handlers
    agentListForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const isFilter = document.getElementById('agentFilterList').checked;
        
        let data = {};
        if (isFilter) {
            data = {
                fromDate: document.getElementById('fromDate').value,
                toDate: document.getElementById('toDate').value,
                agentCode: document.getElementById('agentCode').value,
                agentBranch: document.getElementById('agentBranch').value,
                agentCategory: parseInt(document.getElementById('agentCategory').value) || 0,
                agentActive: document.getElementById('agentActive').value === 'true',
                agentEmpType: document.getElementById('agentEmpType').value
            };
        }

        try {
            const result = await makeApiCall(API_ENDPOINTS.AGENT_LIST, data);
            updateResult(result);
        } catch (error) {
            document.getElementById('resultTable').innerHTML = 
                `<div class="alert alert-danger">Error: ${error.message}</div>`;
        }
    });

    orderBookForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const isFilter = document.getElementById('orderBookFilter').checked;
        
        let data = {};
        if (isFilter) {
            data = {
                fromDate: document.getElementById('obFromDate').value,
                toDate: document.getElementById('obToDate').value,
                agentCode: document.getElementById('obAgentCode').value,
                matchAccount: parseInt(document.getElementById('matchAccount').value) || 0,
                fdSchemeName: document.getElementById('fdSchemeName').value,
                interestRate: parseFloat(document.getElementById('interestRate').value) || 0,
                maturityDays: parseInt(document.getElementById('maturityDays').value) || 0
            };
        }

        try {
            const result = await makeApiCall(API_ENDPOINTS.ORDER_BOOK, data);
            updateResult(result);
        } catch (error) {
            document.getElementById('resultTable').innerHTML = 
                `<div class="alert alert-danger">Error: ${error.message}</div>`;
        }
    });

    healthCheck.addEventListener('click', async () => {
        try {
            const result = await makeApiCall(API_ENDPOINTS.HEALTH);
            updateResult(result);
        } catch (error) {
            document.getElementById('resultTable').innerHTML = 
                `<div class="alert alert-danger">Error: ${error.message}</div>`;
        }
    });

    refreshBtn.addEventListener('click', async () => {
        if (currentHistoryIndex >= 0) {
            try {
                const lastResult = resultHistory[currentHistoryIndex];
                updateResult(lastResult);
            } catch (error) {
                document.getElementById('resultTable').innerHTML = 
                    `<div class="alert alert-danger">Error refreshing data: ${error.message}</div>`;
            }
        }
    });

    backBtn.addEventListener('click', () => {
        if (currentHistoryIndex > 0) {
            currentHistoryIndex--;
            const previousResult = resultHistory[currentHistoryIndex];
            document.getElementById('resultTable').innerHTML = createTable(previousResult);
            updateNavigationButtons();
        }
    });

    // Initialize button states
    updateNavigationButtons();
});
