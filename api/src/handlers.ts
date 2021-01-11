import {Request, Response} from "express";
import {IProductsRepository, IUsersRepository} from "./database/repositories/interfaces";
import HttpStatus from "http-status-codes"
import jsend from "jsend"
import {IClient} from "./discount/client/interfaces";
import Product from "./database/models/product";

export default class Handlers {
    private grpcClient: IClient;
    private usersRepository: IUsersRepository;
    private productsRepository: IProductsRepository;

    constructor(grpcClient: IClient, usersRepository: IUsersRepository, productsRepository: IProductsRepository) {
        this.grpcClient = grpcClient;
        this.usersRepository = usersRepository;
        this.productsRepository = productsRepository;
    }

    async GetProductsHandler(req: Request, res: Response) {
        let products = await this.productsRepository.getAll()
        let userId = req.query.userId
        if (userId) {
            // @ts-ignore
            let user = await this.usersRepository.get(parseInt(userId))
            if (user) {
                let updatedProducts: Product[] = await Promise.all(products.map(async (product: Product): Promise<Product> => {
                    await this.grpcClient.calculate(product._id, user._id).then(discount => {
                        product.discount = discount
                    }).catch(() => {
                    })
                    return product
                }))
                return res.status(HttpStatus.OK).json(jsend.success(updatedProducts))
            }
        }
        return res.status(HttpStatus.OK).json(jsend.success(products))
    }

}