async function startScraping() {
    const urls = document.getElementById('urlInput').value.split('\n').filter(u => u.trim() !== "");
    const workers = parseInt(document.getElementById('workers').value);
    const profile = document.getElementById('profile').value;
    const resultsDiv = document.getElementById('results');

    if (urls.length === 0) {
        alert("Please enter at least one URL!");
        return;
    }

    document.getElementById('btn').disabled = true;
    resultsDiv.innerHTML = `
        <div class="flex flex-col items-center">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-700 mb-4"></div>
            <p class="text-indigo-700 font-medium">Scraping in progress...</p>
        </div>
    `;

    resultsDiv.classList.remove('justify-center', 'items-center', 'text-center');

    try {
        const response = await fetch('/api/scrape', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ urls, workers, profile })
        });

        const result = await response.json();
        renderResults(result);
    } catch (error) {
        resultsDiv.innerHTML = `<p class="text-red-500">Error: Could not reach the server.</p>`;
        document.getElementById('btn').disabled = false;
    }
}

function renderResults(res) {
    document.getElementById('btn').disabled = false;
    document.getElementById('stats').classList.remove('hidden');
    document.getElementById('time').innerText = res.execution_time;
    document.getElementById('count').innerText = res.total_results;

    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = "";

    res.data.forEach(item => {
        const card = document.createElement('div');
        card.className = `p-4 rounded-lg shadow bg-white border-l-4 ${item.error ? 'border-red-500' : 'border-green-500'}`;
        card.innerHTML = `
                    <h3 class="font-bold text-sm text-gray-500">${item.url}</h3>
                    <p class="text-lg">${item.title || '<span class="text-red-400">Error: ' + item.error + '</span>'}</p>
                    ${item.description ? `<p class="text-sm text-gray-600 mt-2">${item.description}</p>` : ''}
                `;
        resultsDiv.appendChild(card);
    });
}

document.getElementById('workers').oninput = function() {
    document.getElementById('workerVal').innerText = this.value;
};