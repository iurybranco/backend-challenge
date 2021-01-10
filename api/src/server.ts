import express from 'express'
import 'express-async-errors'
import './database/connection'
import routes from "./routes"
import * as http from 'http';
import errorHandler from "./errors/handler";

const app = express()
const server = http.createServer(app)
app.use(express.json())
app.use(routes)
app.use(errorHandler)

server.listen(3333)
