// frontend/src/Home.js
import React from 'react';
import { Link } from 'react-router-dom';
import './Home.css';

function Home() {
    return (
        <div className="home">
            <h1>Welcome to YouTube Dashboard</h1>
            <div className="buttons-container">
                <Link to="/likedvideos" className="nav-button">
                    Liked Videos
                </Link>
                <Link to="/subscriptions" className="nav-button">
                    Subscriptions
                </Link>
                <Link to="/searchvideos" className="nav-button">
                    Search Videos
                </Link>
                <Link to="/channeldetails" className="nav-button">
                    Channel Details
                </Link>
            </div>
        </div>
    );
}

export default Home;
