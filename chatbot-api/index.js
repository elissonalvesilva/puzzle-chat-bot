const { Client } = require('whatsapp-web.js');
const qrcode = require('qrcode-terminal');
const express = require('express');

const client = new Client();

const app = express();

app.get('/', (req, res) => {
  res.send('Chatbot do WhatsApp');
});

app.get('/qr', (req, res) => {
  client.on('qr', (qr) => {
    qrcode.generate(qr, { small: true });
    res.send(qr);
  });
});

app.listen(3000, () => {
  console.log('Servidor rodando em http://localhost:3000');
});

client.on('message', (message) => {
  if (message.body.toLowerCase() === 'oi') {
    message.reply('Olá! Como posso ajudar?');
  }
});

client.on('ready', () => {
  console.log('Aplicação pronta!');
});

client.on('auth_failure', (error) => {
  console.error('Erro na autenticação:', error);
});

client.initialize();