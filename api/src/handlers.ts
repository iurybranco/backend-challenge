import {Request, Response} from "express";
import {IProductsRepository, IUsersRepository} from "./database/repositories/interfaces";
import HttpStatus from "http-status-codes"
import jsend from "jsend"
import {IClient} from "./discount/client/interfaces";

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
                await Promise.all(products.map(async product => {
                    await this.grpcClient.calculate(1, user._id).then(discount => {
                        product.discount = discount
                    }).catch(() => {
                    })
                }))
                return res.status(HttpStatus.OK).json(jsend.success(products))
            }
        }
        return res.status(HttpStatus.OK).json(jsend.success(products))
    }

}