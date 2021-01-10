// package: 
// file: proto/discount.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as proto_discount_pb from "../proto/discount_pb";

interface IDiscountService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    calculate: IDiscountService_ICalculate;
}

interface IDiscountService_ICalculate extends grpc.MethodDefinition<proto_discount_pb.Request, proto_discount_pb.Response> {
    path: "/Discount/Calculate";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_discount_pb.Request>;
    requestDeserialize: grpc.deserialize<proto_discount_pb.Request>;
    responseSerialize: grpc.serialize<proto_discount_pb.Response>;
    responseDeserialize: grpc.deserialize<proto_discount_pb.Response>;
}

export const DiscountService: IDiscountService;

export interface IDiscountServer extends grpc.UntypedServiceImplementation {
    calculate: grpc.handleUnaryCall<proto_discount_pb.Request, proto_discount_pb.Response>;
}

export interface IDiscountClient {
    calculate(request: proto_discount_pb.Request, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
    calculate(request: proto_discount_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
    calculate(request: proto_discount_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
}

export class DiscountClient extends grpc.Client implements IDiscountClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public calculate(request: proto_discount_pb.Request, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
    public calculate(request: proto_discount_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
    public calculate(request: proto_discount_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_discount_pb.Response) => void): grpc.ClientUnaryCall;
}
