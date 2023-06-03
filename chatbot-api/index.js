const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const QRCode = require('qrcode-svg');
const app = express();
const steps = require('./steps');

const client = new Client({
    authStrategy: new LocalAuth()
});

client.on('ready', () => {
  console.log('Cliente está pronto!');
});

client.on('message', (message) => {

  if(message.body === "iniciar") {
    client.sendMessage(steps[0]['question'])
  }
  if(message.body === 2) {
    client.sendMessage(steps[1]['question'])
  }
  if(message.body === 3) {
    client.sendMessage(steps[2]['question'])
  }
  if(message.body === 4) {
    client.sendMessage(steps[3]['question'])
  }
  if(message.body === 5) {
    client.sendMessage(steps[4]['question'])
  }
  if(message.body === 6) {
    client.sendMessage(steps[5]['question'])
  }
  if(message.body === 7) {
    client.sendMessage(steps[6]['question'])
  }
});

client.initialize();


app.get('/auth', (req, res) => {
  client.on('qr', (qr) => {
    const qrCode = new QRCode(qr);
    const qrCodeSvg = qrCode.svg();
    res.setHeader('Content-Type', 'image/svg+xml');
    res.send(qrCodeSvg);
  });
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