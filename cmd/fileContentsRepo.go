package cmd

var Database = `require("dotenv").config();
const mongoose = require("mongoose");

const dbConnect = () =>{
    mongoose.connect(process.env.DB_URL,{
        useNewurlParser:true,
        useUnifiedTopology:true,
    }).then(()=>{
        console.log("connected Succesfully ");
    }).catch((error) => {
        console.log("Recieved an error" ,error );
    })
}

module.exports = dbConnect;`

var Env = `PORT = 4001
DB_URL = "mongodb://127.0.0.1:27017/projectDB" 
JWT_SECRET_KEY = KEY`

var GitIgn = ``

var ServerData = `const express = require("express");
require("dotenv").config();
const dbConnect = require("./config/database"); 
const cors = require("cors");
const app = express();
const jwt = require("jsonwebtoken");
const mongoose = require("mongoose");

app.use(express.json());
app.use(cors());
 
dbConnect();

app.get("/",(req,res) => {
    res.send("Hello World!");
})
server.listen(process.env.PORT, () => {
    console.log("Server started at port",process.env.PORT);
})`
