const express = require("express");
const app = express();

app.get("/", (req, res) => res.send("Welcome to the Node web server run by a Docker image"));

app.listen(8081, () => {
});