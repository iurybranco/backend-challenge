import {getMongoRepository} from "typeorm";
import User from "../models/User";
import {Request, Response} from "express";
import jsend from "jsend";
import Product from "../models/Product";

export default {

    async getAll(req: Request, res: Response) {
        //TODO implement get discount rule
        // const productRepository = getMongoRepository(Product);
        // const products = await productRepository.find();
        // console.log("product encontrado = ", products)
        // return res.status(200).json(products)
        // const userRepository = getMongoRepository(User);
        // const user = await userRepository.findOne({_id: 2});
        // console.log("user encontrado = ", user)
        // return res.status(200).json(user)
        return res.status(200).json(jsend.success([]))
    }
}