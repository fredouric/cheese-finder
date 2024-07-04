import { GrpcTransport } from '@protobuf-ts/grpc-transport';
import { CheeseAPIClient } from '../../gen/cheese/v1/cheese.client';
import { ChannelCredentials } from '@grpc/grpc-js';

const transport = new GrpcTransport({
	host: 'localhost:3000',
	channelCredentials: ChannelCredentials.createInsecure()
});

export const cheeseClient = new CheeseAPIClient(transport);
