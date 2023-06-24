const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const app = express();
const qrcode = require('qrcode');
const sqlite3 = require('sqlite3').verbose();


const puzzle = require('./puzzle');
const qrcodeT = require('qrcode-terminal');
const { default: axios } = require('axios');

const sqlInsert = `INSERT INTO users (phone_number, current_puzzle, step) VALUES (?, ?, ?)`;
const sqlUpdate = `Update users SET current_puzzle = ? WHERE phone_number = ?`;
const sqlUpdateStep = `Update users SET step = ? WHERE phone_number = ?`;
const sqlGet = `SELECT * FROM users WHERE phone_number = ?`;
const sqlDelete = `DELETE FROM users WHERE phone_number = ?`;


const solverServiceURL = "https://solver-service.onrender.com"
let qrCodeGenerated = ''
const client = new Client({
    authStrategy: new LocalAuth()
});

function createTable() {
  const db = new sqlite3.Database('database.db');
  const sql = `CREATE TABLE IF NOT EXISTS users (
    phone_number VARCHAR(255) PRIMARY KEY,
    current_puzzle INTEGER
  )`;

  db.run(sql, (err) => {
    if (err) {
      console.error(err.message);
    } else {
      console.log('Tabela criada com sucesso!');
    }
  });

  db.close();
}

client.on('ready', () => {
  console.log('Cliente está pronto!');
  createTable();
});

client.on('qr', async (qr) => {
  qrCodeGenerated = qr;
  qrCodeGenerated = await qrcode.toDataURL(qr, { width: 100 });
  qrcodeT.generate(qr, { small: true });
});


async function getUserFromDBByPhoneNumer(phoneNumber) {
  const db = new sqlite3.Database('database.db');

  return new Promise((resolve, reject) => {
    db.get(sqlGet, [phoneNumber], (err, row) => {
      if(err) {
        console.log(err);
        reject(null)
      }else {
        resolve(row)
      }
    })
    db.close();

  })
}

async function createUserByPhoneNumber(phoneNumber) {
  const db = new sqlite3.Database('database.db');

  return new Promise((resolve, reject) => {
    db.run(sqlInsert, [phoneNumber, 1], (err) => {
      if(err) {
        reject(err)
      }else {
        resolve(this.changes)
      }
    })
    db.close();

  });
}

async function updateCurrentPuzzleUserByPhoneNumber(phoneNumber, puzzleId) {
  const db = new sqlite3.Database('database.db');

  return new Promise((resolve, reject) => {
    db.run(sqlUpdate, [puzzleId, phoneNumber], (err) => {
      if(err) {
        reject(err)
      }else {
        resolve(this.lastID)
      }
    })
    db.close();

  });
}

async function updateCurrentPuzzleUserByPhoneNumber(phoneNumber, step) {
  const db = new sqlite3.Database('database.db');

  return new Promise((resolve, reject) => {
    db.run(sqlUpdateStep, [step, phoneNumber], (err) => {
      if(err) {
        reject(err)
      }else {
        resolve(this.lastID)
      }
    })
    db.close();

  });
}

// client.on('message', async (message) => {

//   const allowedMessageNumberRegex = /^\d+@c\.us$/;
//   if (message.body != "" && allowedMessageNumberRegex.test(message.from) ) {
    
//     const user = await getUserFromDBByPhoneNumer(message.from);
//     if(!user) {
//       await createUserByPhoneNumber(message.from);
//     }
  
//     // if(message.body.toLowerCase() === "iniciar") {
//     //   client.sendMessage(message.from, puzzle[0]['question'])
//     //   await updateCurrentPuzzleUserByPhoneNumber(message.from, 1);
//     //   return
//     // }

//     // if(message.body == "2") {
//     //   client.sendMessage(message.from, puzzle[1]['question'])
//     //   return
//     // }
//     // if(message.body == "3") {
//     //   client.sendMessage(message.from, puzzle[2]['question'])
//     //   return
//     // }
//     // if(message.body == "4") {
//     //   client.sendMessage(message.from, puzzle[3]['question'])
//     //   return
//     // }
//     // if(message.body == "5") {
//     //   client.sendMessage(message.from, puzzle[4]['question'])
//     //   return
//     // }
//     // if(message.body == "6") {
//     //   client.sendMessage(message.from, puzzle[5]['question'])
//     //   return
//     // }
//     // if(message.body == "7") {
//     //   client.sendMessage(message.from, puzzle[6]['question'])
//     //   return
//     // }

//     // if(user) {
//     // client.sendMessage(message.from, "Processando a resposta.... :)");
//     //   const puzzleId = user.current_puzzle;
//     //   const answer = message.body;
//     //   axios.post(`${solverServiceURL}/puzzle`, {
//     //     puzzle_id: puzzleId,
//     //     answer,
//     //   }).then(async (resp) => {
//     //     client.sendMessage(message.from, "Resposta correta")
//     //     console.log(resp.data)
//     //     await updateCurrentPuzzleUserByPhoneNumber(message.from, user.current_puzzle+1);
//     //     client.sendMessage(message.from, resp.data.message.clue);
//     //   }).catch((err) => {
//     //     console.log(err);
//     //     const percentage = (err.response.data?.percentage * 100).toFixed(2);
//     //     const formatedPercentage = percentage.toLocaleString('pt-BR');
//     //     client.sendMessage(message.from, `Resposta incorreta, aproximação da resposta em ${formatedPercentage}%`)
//     //   })
//     // }
//   }

// });

client.initialize();

app.listen(3000, () => {
  console.log('Servidor iniciado na porta 3000');
});