// @ts-check
"use strict";

//load modules
const express = require("express");
const morgan = require("morgan"); //logging middleware

const addr = process.env.ADDR || ":80";
const [host, port] = addr.split(":");
const portNum = parseInt(port);

const handlers = require("./handlers.js");

const app = express();
app.use(morgan(process.env.LOG_FORMAT || "dev"));

app.get("/", (req, res) => {
    res.set("Content-Type", "text/plain")
    res.send("Hello, Node.js!")
})

app.get("/v1/users/me/hello", (req, res) => {
    let userJSON = req.get("X-User");
    if(!userJSON) {
        throw new Error("No X-User header provided");
    }
    let user = JSON.parse(userJSON);
    res.json({
        message: `Hello ${user.firstName} ${user.lastname}`
    })
})

app.use(handlers({}))

//invoked when any handlers throw an exception or error
app.use((err, req, res, next) => {
    console.error(err.stack);
    res.set("Content-Type", "text/plain");
    res.send(err.message);
})

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`)
})
