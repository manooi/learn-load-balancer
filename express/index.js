const express = require('express');
const app = express();
const port = 3000;
const path = require('path');

// Serve static files from the 'public' directory
app.use(express.static(path.join(__dirname, 'public')));

app.get('/test', (req, res) => {
  console.log("");
  res.json({"status": "ok"});
});

app.listen(port, () => {
  console.log(`Express API listening at http://localhost:${port}`);
});
