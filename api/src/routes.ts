import {Router} from "express";
import Handlers from "./handlers";
import Client from "./discount/client/client"
import UsersRepository from "./database/repositories/usersRepository"
import ProductsRepository from "./database/repositories/productsRepository"

const routes = Router()
let grpcHost = process.env.GRPC_SERVER_HOST ? process.env.GRPC_SERVER_HOST : "localhost"
let grpcPort = process.env.GRPC_SERVER_PORT ? process.env.GRPC_SERVER_PORT : "3005"
let grpcClient = new Client(grpcHost, parseInt(grpcPort))
let usersRepo = new UsersRepository()
let productsRepo = new ProductsRepository()
let handlers = new Handlers(grpcClient, usersRepo, productsRepo)
routes.get("/product", handlers.GetProductsHandler.bind(handlers))

export default routes