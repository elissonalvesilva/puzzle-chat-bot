// import React, { useState } from 'react';
// import useWebSocket from 'react-use-websocket';

// //const App = () => {
//   const [message, setMessage] = useState('');
//   const { sendJsonMessage } = useWebSocket('ws://localhost:8001/ws');

//   const handleMessageChange = (event) => {
//     setMessage(event.target.value);
//   };

//   const handleSendMessage = () => {
//     sendJsonMessage({ message });
//   };

//   return (
//       <div>
//         <input type="text" value={message} onChange={handleMessageChange} />
//         <button onClick={handleSendMessage}>Enviar</button>
//       </div>
//   );
// };

// export default App;