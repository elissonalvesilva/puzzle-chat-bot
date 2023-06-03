const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const QRCode = require('qrcode-svg');
const app = express();
const steps = require('./steps');
const qrcode = require('qrcode-terminal');
const { default: axios } = require('axios');


const solverServiceURL = "https://solver-service.onrender.com"
let qrCodeGenerated = ''
const client = new Client({
    authStrategy: new LocalAuth()
});

client.on('ready', () => {
  console.log('Cliente está pronto!');
});

client.on('qr', (qr) => {
  qrCodeGenerated = qr;
  // const qrCode = new QRCode(qr);
  // const qrCodeSvg = qrCode.svg();
  // res.setHeader('Content-Type', 'image/svg+xml');
  // res.send(qrCodeSvg);
  qrcode.generate(qr, { small: true });

});

client.on('message', (message) => {
  if(message.body === "iniciar") {
    client.sendMessage(message.from, steps[0]['question'])
  }

  if(message.body == "2") {
    client.sendMessage(message.from, steps[1]['question'])
  }
  if(message.body == "3") {
    client.sendMessage(message.from, steps[2]['question'])
  }
  if(message.body == "4") {
    client.sendMessage(message.from, steps[3]['question'])
  }
  if(message.body == "5") {
    client.sendMessage(message.from, steps[4]['question'])
  }
  if(message.body == "6") {
    client.sendMessage(message.from, steps[5]['question'])
  }
  if(message.body == "7") {
    client.sendMessage(message.from, steps[6]['question'])
  }


  const pattern = /1,r=(.*)/;
  const correspondence = message.body.match(pattern)
  if (correspondence) {
    const puzzleId = correspondence[0]
    const answer = correspondence[1];
    axios.post(`${solverServiceURL}/puzzle`, {
      puzzle_id: puzzleId,
      answer,
    }).then((resp) => {
      client.sendMessage(message.from, "Resposta correta")
      console.log(resp.data)
    }).catch((err) => {
      console.log(err);
      client.sendMessage(message.from, "Resposta incorreta")
    })

  }else {
    client.sendMessage(message.from, "Mensagem invalida ou vazia")
  }
});

client.initialize();


app.get('/auth', (req, res) => {
  if(qrCodeGenerated === '') {
    res.send({"message": "not generated"});
  }
  const qrCode = new QRCode(qrCodeGenerated);
  const qrCodeSvg = qrCode.svg();
  res.setHeader('Content-Type', 'image/svg+xml');
  res.send(qrCodeSvg);
});

app.get('/iniciar', (req, res) => {
  const qr = new QRCode({
    content: "https://api.whatsapp.com/send?phone=5592993806358&text=iniciar",
    padding: 4,
    width: 256,
    height: 256,
    color: '#000000',
    background: '#ffffff',
    ecl: 'M',
  });
  
  // Gere o código SVG do QR code
  const qrCodeSvg = qr.svg();
  res.setHeader('Content-Type', 'image/svg+xml');
  res.send(qrCodeSvg);
})

app.listen(3000, () => {
  console.log('Servidor iniciado na porta 3000');
});