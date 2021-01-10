// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_discount_pb = require('../proto/discount_pb.js');

function serialize_Request(arg) {
  if (!(arg instanceof proto_discount_pb.Request)) {
    throw new Error('Expected argument of type Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Request(buffer_arg) {
  return proto_discount_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_Response(arg) {
  if (!(arg instanceof proto_discount_pb.Response)) {
    throw new Error('Expected argument of type Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Response(buffer_arg) {
  return proto_discount_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var DiscountService = exports.DiscountService = {
  calculate: {
    path: '/Discount/Calculate',
    requestStream: false,
    responseStream: false,
    requestType: proto_discount_pb.Request,
    responseType: proto_discount_pb.Response,
    requestSerialize: serialize_Request,
    requestDeserialize: deserialize_Request,
    responseSerialize: serialize_Response,
    responseDeserialize: deserialize_Response,
  },
};

exports.DiscountClient = grpc.makeGenericClientConstructor(DiscountService);
