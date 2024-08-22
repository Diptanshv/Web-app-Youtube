document.addEventListener('DOMContentLoaded', function () {
    if (window.location.pathname.includes('subscriptions.html')) {
        fetchSubscriptions();
    } else if (window.location.pathname.includes('likedvideos.html')) {
        fetchLikedVideos();
    }
});

function fetchSubscriptions() {
    fetch('/api/subscriptions')
        .then(response => response.json())
        .then(data => {
            const container = document.getElementById('subscriptions');
            container.innerHTML = '';
            data.forEach(subscription => {
                const item = document.createElement('div');
                item.className = 'subscription-item';
                item.innerHTML = `
                    <img src="${subscription.thumbnail}" alt="${subscription.title}">
                    <div>
                        <h3>${subscription.title}</h3>
                        <a href="${subscription.url}" target="_blank">Visit Channel</a>
                    </div>
                `;
                container.appendChild(item);
            });
        });
}

function fetchLikedVideos() {
    fetch('/api/likedvideos')
        .then(response => response.json())
        .then(data => {
            const container = document.getElementById('liked-videos');
            container.innerHTML = '';
            data.forEach(video => {
                const item = document.createElement('div');
                item.className = 'video-item';
                item.innerHTML = `
                    <img src="${video.thumbnail}" alt="${video.title}">
                    <div>
                        <h3>${video.title}</h3>
                        <p>Views: ${video.viewCount}</p>
                        <p>Likes: ${video.likeCount}</p>
                    </div>
                `;
                container.appendChild(item);
            });
        });
}

function searchVideos() {
    const query = document.getElementById('search-query').value;
    fetch(`/api/search?query=${encodeURIComponent(query)}`)
        .then(response => response.json())
        .then(data => {
            const container = document.getElementById('search-results');
            container.innerHTML = '';
            data.forEach(video => {
                const item = document.createElement('div');
                item.className = 'video-item';
                item.innerHTML = `
                    <img src="${video.thumbnail}" alt="${video.title}">
                    <div>
                        <h3>${video.title}</h3>
                    </div>
                `;
                container.appendChild(item);
            });
        });
}
