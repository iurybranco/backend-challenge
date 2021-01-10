import Discount from "../../database/models/discount";

export interface IClient {
    calculate(product_id: number, user_id: number): Promise<Discount>;
}