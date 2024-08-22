// import React, { useState } from 'react';
// import './SearchVideos.css';

// function SearchVideos() {
//     const [query, setQuery] = useState('');
//     const [videos, setVideos] = useState([]);

//     const handleSearch = async () => {
//         const response = await fetch(`/api/searchvideos?q=${query}`);
//         const data = await response.json();
//         setVideos(data);
//     };

//     return (
//         <div className="search-videos">
//             <input 
//                 type="text" 
//                 value={query} 
//                 onChange={(e) => setQuery(e.target.value)} 
//                 placeholder="Search videos" 
//             />
//             <button onClick={handleSearch}>Search</button>
//             <div className="videos-list">
//                 {videos.map((video, index) => (
//                     <div key={index} className="video">
//                         <img src={video.snippet.thumbnails.default.url} alt={video.snippet.title} />
//                         <p>{video.snippet.title}</p>
//                     </div>
//                 ))}
//             </div>
//         </div>
//     );
// }

// export default SearchVideos;
import React, { useState } from 'react';
import './SearchVideos.css';

function SearchVideos() {
    const [query, setQuery] = useState('');
    const [videos, setVideos] = useState([]);

    const handleSearch = () => {
        fetch(`/api/search?q=${encodeURIComponent(query)}`)
            .then(response => response.json())
            .then(data => setVideos(data))
            .catch(error => console.error('Error fetching search results:', error));
    };

    return (
        <div className="search-container">
            <input
                type="text"
                placeholder="Search for videos..."
                value={query}
                onChange={(e) => setQuery(e.target.value)}
            />
            <button onClick={handleSearch}>Search</button>
            <div className="videos-container">
                {videos.map((video, index) => (
                    <div className="video-card" key={index}>
                        <img src={video.thumbnail} alt={video.title} />
                        <h2>{video.title}</h2>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default SearchVideos;

