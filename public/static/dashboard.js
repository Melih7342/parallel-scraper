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
        console.log("Full Backend Response:", result);
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
    const profile = document.getElementById('profile').value;

    const resultsDiv = document.getElementById('results');
    resultsDiv.innerHTML = "";

    res.data.forEach(item => {
        // 1. Preparation of logic for dead links section
        let deadLinksSection = "";

        if (item.dead_links && item.dead_links.length > 0) {
            deadLinksSection = `
                <div class="mt-4 pt-4 border-t border-red-100">
                    <div class="flex items-center justify-between mb-2">
                        <span class="flex items-center text-red-600 font-bold text-sm">
                            <span class="mr-2">⚠️</span> ${item.dead_links_count} Dead Links found
                        </span>
                    </div>
                    <div class="bg-red-50 rounded p-2 border border-red-100">
                        <ul class="text-xs font-mono text-red-500 space-y-1 max-h-32 overflow-y-auto custom-scrollbar">
                            ${item.dead_links.map(link => `<li class="break-all hover:bg-red-100 p-1 rounded transition-colors">• ${link}</li>`).join('')}
                        </ul>
                    </div>
                </div>
            `;
        } else if (profile === "dead-links" && item.dead_links_count === 0) {
            deadLinksSection = `
        <div class="mt-4 pt-3 border-t border-gray-100 flex items-center text-xs text-green-600">
            <span class="mr-2">✅</span> 
            <span>Checked ${item.total_links_found} links. All are healthy!</span>
        </div>
    `;
    }

        // 2. Create Card
        const card = document.createElement('div');
        card.className = `p-5 rounded-lg shadow-sm bg-white border-l-4 transition-all hover:shadow-md ${item.error ? 'border-red-500' : 'border-green-500'}`;

        card.innerHTML = `
            <div class="flex justify-between items-start mb-2">
                <h3 class="font-bold text-xs text-gray-400 truncate max-w-[80%]">${item.url}</h3>
                ${item.dead_links_count > 0 ? `<span class="bg-red-100 text-red-700 text-[10px] px-2 py-1 rounded-full font-black">DEAD: ${item.dead_links_count}</span>` : ''}
            </div>
            
            <p class="text-lg font-semibold text-gray-800 leading-tight">
                ${item.title || (item.error ? `<span class="text-red-500 italic">Failed to load</span>` : `<span class="text-gray-400 italic font-normal">No title found</span>`)}
            </p>
            
            ${item.description ? `<p class="text-sm text-gray-600 mt-2 line-clamp-2">${item.description}</p>` : ''}
            
            ${item.error ? `<p class="text-xs bg-red-50 text-red-600 p-2 mt-3 rounded border border-red-100 font-mono">${item.error}</p>` : ''}
            
            ${deadLinksSection}
        `;
        resultsDiv.appendChild(card);
    });
}

document.getElementById('workers').oninput = function() {
    document.getElementById('workerVal').innerText = this.value;
};