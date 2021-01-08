import express from 'express'
import routes from "./routes"
import * as http from 'http';

const app = express()
const server = http.createServer(app)
app.use(express.json())
app.use(routes)


server.listen(3334)
