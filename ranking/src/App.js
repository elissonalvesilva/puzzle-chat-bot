// import React, { useState } from 'react';
// import useWebSocket from 'react-use-websocket';

// //const App = () => {
//   const [message, setMessage] = useState('');
//   const { sendJsonMessage } = useWebSocket('ws://localhost:8001/ws');

<<<<<<< Updated upstream
//   const handleMessageChange = (event) => {
//     setMessage(event.target.value);
//   };

//   const handleSendMessage = () => {
//     sendJsonMessage({ message });
//   };
=======
    const handleSendMessage = () => {
        sendMessage('refresh'); // Enviar mensagem para o servidor
    };

    console.log(lastMessage?.data)

    return (
        <div>
            <div>Status da conexão: {readyState === 1 ? 'Conectado' : 'Desconectado'}</div>
            <div>Última mensagem recebida: {lastMessage && lastMessage.data}</div>
            <button onClick={handleSendMessage}>Enviar mensagem para o servidor</button>
        </div>
    );
};
>>>>>>> Stashed changes

//   return (
//       <div>
//         <input type="text" value={message} onChange={handleMessageChange} />
//         <button onClick={handleSendMessage}>Enviar</button>
//       </div>
//   );
// };

// export default App;