import {Request, Response} from "express";
import {IProductsRepository, IUsersRepository} from "./database/repositories/interfaces";
import HttpStatus from "http-status-codes"
import jsend from "jsend"
import {IClient} from "./discount/client/interfaces";
import {ServiceError} from "@grpc/grpc-js";

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
        let userId = req.query.userId
        if (!userId) {
            return res.status(HttpStatus.BAD_REQUEST).json(jsend.fail({"userId": "userId is a required query param"}))
        }
        let products = await this.productsRepository.getAll()
        let user = await this.usersRepository.get(1)
        this.grpcClient.calculate(1,1).then(discount =>{
            return res.status(HttpStatus.OK).json(jsend.success(discount))
        }).catch((err: ServiceError)=>{
            return res.status(HttpStatus.INTERNAL_SERVER_ERROR).json(jsend.error(err.message))
        })
    }

}