import React, { useState } from 'react';
import './ChannelDetails.css';

function ChannelDetails() {
    const [channelId, setChannelId] = useState('');
    const [channel, setChannel] = useState(null);

    const handleGetDetails = async () => {
        const response = await fetch(`/api/channeldetails?channelId=${channelId}`);
        const data = await response.json();
        setChannel(data);
    };

    return (
        <div className="channel-details">
            <input 
                type="text" 
                value={channelId} 
                onChange={(e) => setChannelId(e.target.value)} 
                placeholder="Enter channel ID" 
            />
            <button onClick={handleGetDetails}>Get Details</button>
            {channel && (
                <div className="channel-info">
                    <img src={channel.snippet.thumbnails.default.url} alt={channel.snippet.title} />
                    <p>{channel.snippet.title}</p>
                    <p>Subscribers: {channel.statistics.subscriberCount}</p>
                    <p>Views: {channel.statistics.viewCount}</p>
                </div>
            )}
        </div>
    );
}

export default ChannelDetails;
