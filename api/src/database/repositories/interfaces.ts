import Product from "../models/product";
import User from "../models/user";

export interface IProductsRepository {
    getAll(): Promise<Product[]>;
}

export interface IUsersRepository {
    get(id: number): Promise<User | undefined>;
}