import {Router} from "express";
import Handlers from "./handlers";
import Client from "./discount/client/client"
import UsersRepository from "./database/repositories/usersRepository"
import ProductsRepository from "./database/repositories/productsRepository"

const routes = Router()
let grpcClient = new Client(3000)
let usersRepo = new UsersRepository()
let productsRepo = new ProductsRepository()
let handlers = new Handlers(grpcClient, usersRepo, productsRepo)
routes.get("/product",  handlers.GetProductsHandler.bind(handlers))

export default routes