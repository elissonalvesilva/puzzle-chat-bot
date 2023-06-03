require('dotenv').config();

const express = require('express');
const app = express();
app.use(express.json());

const token = process.env.WEBHOOK_TOKEN;


app.post('/webhook', (req, res) => {
  if (token === verificationToken) {
    res.sendStatus(200);
  } else {
    res.sendStatus(403);
  }
});

app.listen(3000, () => {
  console.log('Server running in http://localhost:3000');
});