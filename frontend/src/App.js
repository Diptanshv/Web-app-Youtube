// import React from 'react';
// import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
// import Subscriptions from './Subscriptions';
// import LikedVideos from './components/LikedVideos';
// import './App.css';

// function App() {
//   return (
//     <Router>
//       <div className="App">
//         <nav>
//           <ul>
//             <li>
//               <Link to="/subscriptions">Subscriptions</Link>
//             </li>
//             <li>
//               <Link to="/likedvideos">Liked Videos</Link>
//             </li>
//           </ul>
//         </nav>
//         <Routes>
//           <Route path="/subscriptions" element={<Subscriptions />} />
//           <Route path="/likedvideos" element={<LikedVideos />} />
//         </Routes>
//       </div>
//     </Router>
//   );
// }

// export default App;
// frontend/src/App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './Home';
import LikedVideos from './components/LikedVideos';
import Subscriptions from './Subscriptions';
import SearchVideos from './SearchVideos';
import ChannelDetails from './ChannelDetails';
import './App.css';

function App() {
    return (
        <Router>
            <div className="app">
                <Routes>
                    <Route path="/likedvideos" element={<LikedVideos />} />
                    <Route path="/subscriptions" element={<Subscriptions />} />
                    <Route path="/searchvideos" element={<SearchVideos />} />
                    <Route path="/channeldetails" element={<ChannelDetails />} />
                    <Route path="/" element={<Home />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;

