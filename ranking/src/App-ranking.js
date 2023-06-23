import styled from "styled-components";
import estrela from "./img/estrela.png";
import React, { useEffect, useState } from 'react';
import axios from "axios";


function Star() {
  return (
    <>
      <div className="start">
        <img className="estrela" src={estrela} alt="estrela"/>
      </div>
    </>
  )
}

function Stars({current = 0}) {
  let startsComp = []; 
  for (let i = 0; i < current; i++) {
    startsComp.push(<Star key={i}/>)
  }
  
  return startsComp
}

export default function AppRanking(){
  const api = process.env.REACT_APP_RANKING_API;
  const [users, setUsers] = useState([])

  useEffect(() => {
    const fetchData = async () => {
      try {
        const {data} = await axios.get(api); // Substitua pela URL da sua API
        setUsers(data.data.Users);
      } catch (error) {
        console.error(error);
      }
    };

    const intervalId = setInterval(fetchData, 5000); // Chama a função fetchData a cada 3 segundos

    return () => {
      clearInterval(intervalId); // Limpa o intervalo quando o componente for desmontado
    };
  }, []);

  return(<>
  <ContainerRanking>
    <h1>Ranking</h1>
    <Users>
      <ul>
        {
          users.length > 0 ? (
            users.map((user, index) => (
              <li className="user" key={index}>
                <span className="name">{user.name}</span>
                <div className="starts">
                  {Stars({ current: user.current })}
                </div>
              </li>
            ))
          ): <></>
        }
      </ul>
    </Users>
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
const Users = styled.div`
 .user{
  margin: 10px;
  padding: 10px;
  background-color: white;
  color: white;
  border: 4px solid black;
  margin-bottom: 10px;
  font-size: larger;
  font-family:'Press Start 2P', cursive;
  color: black;
  overflow-x: scroll;
  display: flex;
  align-items: center;
  .name {
    margin-top:3px;
  }
  .starts {
    display: flex;
  }
  .estrela{
    width: 30px;
    height: 30px;
    margin-left: 5px;
  }
 }


`