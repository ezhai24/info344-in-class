"use strict";

const mongodb = require("mongodb");
const MongoStore = require("./taskstore");
const express = require("express");
const app = express();

const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");

const mongoAddr = process.env.DBADDR || "192.168.99.100:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

//connect to mongo db
mongodb.MongoClient.connect(mongoURL)
  .then(db => {
    //initialize a new task store
    let taskStore = new MongoStore(db, "tasks");

    //parses posted JSON and makes
    //it available from req.body
    app.use(express.json());

    app.post("/v1/tasks", (req, res) => {
      //insert
      let task = {
        title: req.body.title,
        completed: false
      }
      taskStore.insert(task)
        .then(task => {
          res.json(task);
        })
        .catch(err => {
          throw err;
        });
    })

    app.get("/v1/tasks", (req, res) => {
      //get all incomplete
      taskStore.getAll(false)
      .then(tasks => {
        res.json(tasks)
      })
      .catch(err => {
        throw err;
      });
    })

    app.patch("/v1/tasks", (req, res) => {
      //update single task by ID
      let taskIDToFetch = req.params.taskID;
      let updates = {
        completed: req.body.completed
      }
      taskStore.update(taskIDToFetch, updates)
        .then(result => {
          res.json(result)
        })
    })


    app.listen(port, host, () => {
      console.log(`server is listening at http://${addr}...`)
    })
  })
  .catch(err => {
    throw err;
  });
