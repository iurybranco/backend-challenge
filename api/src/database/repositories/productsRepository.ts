import {getMongoRepository} from "typeorm";
import Product from "../models/product";
import {IProductsRepository} from "./interfaces";

export default class ProductsRepository implements IProductsRepository {

    async getAll(): Promise<Product[]> {
        return await getMongoRepository(Product).find();
    }

}