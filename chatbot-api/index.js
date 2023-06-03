const express = require('express');
const { urlencoded } = require('body-parser');
const { Twilio } = require('twilio');
const qrcode = require('qrcode-terminal');

const app = express();
app.use(urlencoded({ extended: false }));

const client = new Twilio(process.env.TWILIO_ACCOUNT_SID, process.env.TWILIO_AUTH_TOKEN);


app.get('/qr', (req, res) => {
  const twilioUrl = `https://api.whatsapp.com/send?phone=${encodeURIComponent(
    '+14155238886'
  )}`;
  qrcode.generate(twilioUrl, { small: true });
  res.send(twilioUrl);
});


app.post('/webhook', (req, res) => {
  const twiml = new MessagingResponse();
  twiml.message('Hello from your Twilio chatbot!');
  res.set('Content-Type', 'text/xml');
  res.send(twiml.toString());
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
});