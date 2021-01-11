import * as grpc from '@grpc/grpc-js';
import {DiscountClient} from '../proto/discount_grpc_pb';
import {Request} from "../proto/discount_pb";
import {IClient} from "./interfaces";
import Discount from "../../database/models/discount";

export default class Client implements IClient {
    private client: DiscountClient

    constructor(host: string, port: number) {
        this.client = new DiscountClient(
            `${host}:${port}`,
            grpc.credentials.createInsecure()
        );
    }

    calculate(product_id: number, user_id: number): Promise<Discount> {
        return new Promise<Discount>((resolve, reject) => {
            let req = new Request().setProductId(product_id).setUserId(user_id)
            this.client.calculate(req, (error, resp) => {
                if (error) {
                    reject(error)
                    return
                }
                let discount = new Discount()
                discount.percentage = resp.getPercentage()
                discount.value_in_cents = resp.getValueInCents()
                resolve(discount)
            });
        })

    }
}
