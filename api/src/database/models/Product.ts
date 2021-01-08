import {Entity, Column, ObjectIdColumn} from 'typeorm'
import Discount from "./Discount";

@Entity( "product" )
export default class Product {
    @ObjectIdColumn()
    _id: number
    @Column()
    price_in_cents: string
    @Column()
    title: string
    @Column()
    description: string
    @Column(type => Discount)
    discount: Discount
}