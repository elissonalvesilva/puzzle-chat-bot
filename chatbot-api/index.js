const express = require('express');
const { Client, LocalAuth } = require('whatsapp-web.js');
const QRCode = require('qrcode-svg');
const app = express();

const client = new Client({
    authStrategy: new LocalAuth()
});

client.on('ready', () => {
  console.log('Cliente está pronto!');
});

client.on('message', (message) => {
  console.log('Mensagem recebida:', message.body);
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

app.listen(3000, () => {
  console.log('Servidor iniciado na porta 3000');
});