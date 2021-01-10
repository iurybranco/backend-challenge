import {Entity, Column, ObjectIdColumn} from 'typeorm'

@Entity( "user" )
export default class User {
    @ObjectIdColumn()
    _id: number
    @Column()
    first_name: string
    @Column()
    last_name: string
    @Column()
    date_of_birth: Date
}