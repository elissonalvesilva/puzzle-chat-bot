import styled from "styled-components";
import estrela from "./img/estrela.png";
import React, { useEffect } from 'react';
import io from 'socket.io-client';

const WebSocketComponent = () => {
  useEffect(() => {
    const socket = io('wss://exemplo.com'); 

    const handleReceiveMessage = (message) => {
      console.log('Mensagem recebida:', message);
    };

    socket.on('connect', () => {
      setInterval(() => {
        socket.emit('message', 'Olá, servidor!');
      }, 5000);
    });

    socket.on('disconnect', () => {
      console.log('Conexão encerrada.');
    });
     socket.on('message', handleReceiveMessage);

    return () => {
      socket.disconnect();
    };
  }, []);

  return <div>WebSocket Component</div>;
};

export default function AppRanking(){
  return(<>
  <ContainerRanking>
    <h1>Ranking</h1>
    <Names>
      <ul>
        <li className="name">Nomsdsdsdes
           <img className="estrela" src={estrela} alt="estrela"/>
        </li>
         <li className="name">Nomsdsdsdes
           <img className="estrela" src={estrela} alt="estrela"/>
        </li>
      </ul>
    </Names>
  </ContainerRanking>
  
  </>);
}

const ContainerRanking = styled.div`
background-color: #3f55ad;
padding: 20px;
font-size: larger;
font-family:'Press Start 2P', cursive;
color: white;

`
const Names = styled.div`
 .name{
  margin: 10px;
  padding: 10px;
  background-color: white;
  color: white;
  border: 4px solid black;
  margin-bottom: 10px;
  font-size: larger;
  font-family:'Press Start 2P', cursive;
  color: black;
  .estrela{
    width: 30px;
    height: 30px;
    margin-left: 5px;
  }
 }


`
const Stars = styled.div`
img {
  width: 100%;
}
height: 25px;
width: 25px;

`