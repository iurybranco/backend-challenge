import {Column} from 'typeorm'

export default class Discount {
    @Column({type: "float"})
    percentage: number
    @Column()
    value_in_cents: number
}