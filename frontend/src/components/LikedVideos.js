// public/frontend/src/components/LikedVideos.js
import React, { useEffect, useState } from 'react';
import './LikedVideos.css';

function LikedVideos() {
    const [videos, setVideos] = useState([]);

    useEffect(() => {
        fetch('/api/likedvideos')
            .then(response => response.json())
            .then(data => setVideos(data))
            .catch(error => console.error('Error fetching liked videos:', error));
    }, []);

    return (
        <div className="liked-videos">
            <h1>Liked Videos</h1>
            <div className="videos-container">
                {videos.map((video, index) => (
                    <div className="video-card" key={index}>
                        <img src={video.thumbnail} alt={video.title} />
                        <h2>{video.title}</h2>
                        <p>{video.description}</p>
                        <div className="video-stats">
                            <span>Views: {video.viewCount}</span>
                            <span> Likes: {video.likeCount}</span>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default LikedVideos;
