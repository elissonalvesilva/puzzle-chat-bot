const express = require('express');
const { Client, LocalAuth, MessageMedia } = require('whatsapp-web.js');
const app = express();
const qrcode = require('qrcode');
const sqlite3 = require('sqlite3').verbose();


const puzzle = require('./puzzle');
const steps = require('./steps');
const qrcodeT = require('qrcode-terminal');
const { default: axios } = require('axios');

const sqlInsert = `INSERT INTO users (phone_number, current_puzzle, step) VALUES (?, ?, ?)`;
const sqlUpdate = `Update users SET current_puzzle = ? WHERE phone_number = ?`;
const sqlUpdateStep = `Update users SET step = ? WHERE phone_number = ?`;
const sqlGet = `SELECT * FROM users WHERE phone_number = ?`;
const sqlDelete = `DELETE FROM users WHERE phone_number = ?`;


const solverServiceURL = "https://solver-service.onrender.com";
const rankingURL = "https://ranking-api-mrda.onrender.com/api"
let qrCodeGenerated = ''
const client = new Client({
    authStrategy: new LocalAuth()
});

function createTable() {
  const db = new sqlite3.Database('database.db');
  const sql = `CREATE TABLE IF NOT EXISTS users (
    phone_number VARCHAR(255) PRIMARY KEY,
    current_puzzle INTEGER,
    step INTEGER
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

async function updateCurrentStepUserByPhoneNumber(phoneNumber, step) {
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

client.on('message', async (message) => {

  const allowedMessageNumberRegex = /^\d+@c\.us$/;
  if (message.body != "" && allowedMessageNumberRegex.test(message.from) ) {
    
    const exitsUser = await getUserFromDBByPhoneNumer(message.from);
    if(!exitsUser) {
      await createUserByPhoneNumber(message.from);
    }

    if(message.body.toLowerCase() === "iniciar") {
      client.sendMessage(message.from, steps[0]['info'])
      await updateCurrentStepUserByPhoneNumber(message.from, steps[0]['step']);
      return
    }

    const user = await getUserFromDBByPhoneNumer(message.from);

    if (user.step === 1) {
      try {
        await axios.post(`${rankingURL}/create`, {
          phone: message.from,
          name: message.body,
          current: 1,
        })
        client.sendMessage(message.from, steps[0]['message'])
        client.sendMessage(message.from, puzzle[0]['question'])
        await updateCurrentStepUserByPhoneNumber(message.from, steps[1]['step']);
        return
      } catch (error) {
        console.log(error.data)
        client.sendMessage(message.from, steps[0]['error'])
        return
      }
    }

    if(user.step === 2 && user.current == 1) {
      const media = MessageMedia.fromFilePath('./img1.jpeg');
      client.sendMessage(message.from, media);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[2]['step']);
    }

    if(user.step === 2) {
      client.sendMessage(message.from, puzzle[1]['question']);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[2]['step']);
    }

    if(user.step === 4) {
      const linkLorena = "https://api.whatsapp.com/send?phone=5548998082313&text=Oi%20Lorena%2C%20desvendei%20um%20enigma%20e%20encontrei%20seu%20numero.%20Qual%20o%20proximo%20desafio%20%3F"
      client.sendMessage(message.from, linkLorena);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[3]['step']);
    }

    if(user.step == 5) {
      const dicas = `
      Dica: evangelho eterno
      Dica: Mensagem de salvação      
      `
      const media = MessageMedia.fromFilePath('./qrcode.jpeg');
      client.sendMessage(message.from, media);
      client.sendMessage(message.from, dicas)
      await updateCurrentStepUserByPhoneNumber(message.from, steps[4]['step']);
    }

    if(user.step == 6) {
      const media = MessageMedia.fromFilePath('./img2.jpeg');
      client.sendMessage(message.from, media);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[5]['step']);
    }

    if(user.step == 7) {
      const code = `70YDsj2oBL5RtOc6tJCJzc0GqRFv1Jet5_H0gf6Y9zfLBJ9qolZDm2xfx8nTFLrODL63wbBRzycar_KBJzR_DiA5qYLfnGQfgsxaD9iyxI2vf7QCgB5gXqFVMwyy7WPeFzVbkRLbqvUUERKeNvVKcg`;
      client.sendMessage(message.from, puzzle[3]['question']);
      client.sendMessage(message.from, "Texto criptografado...");
      client.sendMessage(message.from, code);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[6]['step']);
    }

    if(user.step == 8) {
      client.sendMessage(message.from, puzzle[4]['question']);
      await updateCurrentStepUserByPhoneNumber(message.from, steps[6]['step']);
    }
    

    if(user) {
    client.sendMessage(message.from, "Processando a resposta.... :)");
      const puzzleId = user.current_puzzle;
      const answer = message.body;
      axios.post(`${solverServiceURL}/puzzle`, {
        puzzle_id: puzzleId,
        answer,
      }).then(async (resp) => {
        client.sendMessage(message.from, "Resposta correta")
        console.log(resp.data)
        await updateCurrentPuzzleUserByPhoneNumber(message.from, user.current_puzzle+1);
        client.sendMessage(message.from, resp.data.message.clue);
      }).catch((err) => {
        console.log(err);
        const percentage = (err.response.data?.percentage * 100).toFixed(2);
        const formatedPercentage = percentage.toLocaleString('pt-BR');
        client.sendMessage(message.from, `Resposta incorreta, aproximação da resposta em ${formatedPercentage}%`)
      })
    }
  }

});

client.initialize();

app.listen(3000, () => {
  console.log('Servidor iniciado na porta 3000');
});