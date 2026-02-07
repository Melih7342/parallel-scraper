async function startScraping() {
    const urls = document.getElementById('urlInput').value.split('\n').filter(u => u.trim() !== "");
    const workers = parseInt(document.getElementById('workers').value);
    const profile = document.getElementById('profile').value;

    // UI Reset
    document.getElementById('results').innerHTML = "Wait for results...";
    document.getElementById('btn').disabled = true;

    const response = await fetch('/api/scrape', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ urls, workers, profile })
    });

    const result = await response.json();
    renderResults(result);
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