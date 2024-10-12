package contentstore

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

var ServerEnv = `PORT = 4001
DB_URL = "mongodb://127.0.0.1:27017/projectDB" 
JWT_SECRET_KEY = KEY`

var GitIgn = `# Logs
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*
lerna-debug.log*

node_modules
dist
dist-ssr
*.local
.env
# Editor directories and files
.vscode/*
!.vscode/extensions.json
.idea
.DS_Store
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?
`

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
