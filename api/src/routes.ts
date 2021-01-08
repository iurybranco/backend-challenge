import {Router} from "express";
import ProductsController from "./database/controllers/ProductsController";

const routes = Router()

routes.get("/product", ProductsController.getAll)

export default routes