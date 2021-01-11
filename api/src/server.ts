import express from 'express'
import 'express-async-errors'
import './database/connection'
import routes from "./routes"
import * as http from 'http';
import errorHandler from "./errors/handler";
import "dotenv-safe"
import morgan from "morgan"

const app = express()
const server = http.createServer(app)
app.use(morgan('dev'))
app.use(express.json())
app.use(routes)
app.use(errorHandler)
let port = process.env.SERVER_PORT ? process.env.SERVER_PORT : 3004
console.log(`Listening on ${port}`);
server.listen(port)
export default server