#!/usr/bin/env node
"use strict";

const amqp = require("amqplib");
const qName = "testQ";
const mqAddr = process.env.MQADDR || "192.168.99.100:5672";
const mqURL = `amqp://${mqAddr}`;

//immediately executing anonymous async await function
(async function() {
  try {
    console.log("connecting to %s", mqURL);
  
    //await ensures line resolves before next line is run
    let connection = await amqp.connect(mqURL);
    let channel = await connection.createChannel();
    let qConf = await channel.assertQueue(qName, {durable: false});
  
    //send messages to queue
    console.log("starting to send messages...")
    setInterval(() => {
      let msg = "Message sent at " + new Date().toLocaleTimeString();
      channel.sendToQueue(qName, new Buffer(msg));
    }, 1000);
  } catch(err) {
    console.log(err.stack);
  }
})();
