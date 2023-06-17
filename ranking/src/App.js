import React, { useState } from 'react';
import useWebSocket from 'react-use-websocket';

const App = () => {
    const { sendMessage, lastMessage, readyState } = useWebSocket('ws://ranking-api-owmg.onrender.com/ws');

    const handleSendMessage = () => {
        sendMessage('Olá, servidor!'); // Enviar mensagem para o servidor
    };

    return (
        <div>
            <div>Status da conexão: {readyState === 1 ? 'Conectado' : 'Desconectado'}</div>
            <div>Última mensagem recebida: {lastMessage && lastMessage.data}</div>
            <button onClick={handleSendMessage}>Enviar mensagem para o servidor</button>
        </div>
    );
};

export default App;