import React, { useEffect, useState } from 'react';
import './Subscriptions.css';

function Subscriptions() {
    const [subscriptions, setSubscriptions] = useState([]);

    useEffect(() => {
        fetch('/api/subscriptions')
            .then(response => response.json())
            .then(data => setSubscriptions(data))
            .catch(error => console.error('Error fetching subscriptions:', error));
    }, []);

    return (
        <div className="subscriptions-container">
            {subscriptions.map((subscription, index) => (
                <div className="subscription-box" key={index}>
                    <div className="subscription-thumbnail">
                        <img src={subscription.thumbnail} alt={subscription.title} />
                    </div>
                    <div className="subscription-title">{subscription.title}</div>
                    <a 
                        href={subscription.url} 
                        target="_blank" 
                        rel="noopener noreferrer" 
                        className="subscription-link"
                    >
                        Visit Channel
                    </a>
                </div>
            ))}
        </div>
    );
}

export default Subscriptions;
