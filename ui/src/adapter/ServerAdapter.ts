import { Root } from 'protobufjs';
import Adapter from './Adapter';
import { addFromFileDescriptor, resolveAll } from 'proto-descriptor';

import { ReflectionResponse, InvokeRequest, InvokeResponse } from './proto';

export class ServerAdapter implements Adapter {
  constructor(private prefix: string) {
  }

  async fetchReflection(addr: string): Promise<Root> {
    const res = await fetch(`${this.prefix}/api/reflection?host=${addr}`);
    
    if (!res.ok) {
      throw new Error('Failed to fetch reflection data');
    }

    const buf = await res.arrayBuffer();
    const decoded = ReflectionResponse.decode(new Uint8Array(buf));

    if (decoded.error) {
      throw new Error(decoded.error.message || 'Failed to fetch reflection data');
    }

    if (!decoded.reflection) {
      throw new Error('Failed to fetch reflection data');
    }

    const descriptors = decoded.reflection.fileDescriptor;

    const root = new Root();

    if (!descriptors) {
      return root;
    }

    for (const desc of descriptors) {
      addFromFileDescriptor(root, desc);
    }

    // TODO mark implemented services

    resolveAll(root);

    return root;
  }

  async runMethod(addr: string, method: protobuf.Method, data: any): Promise<{}> {
    const requestType = method.resolvedRequestType as protobuf.Type;
    const responseType = method.resolvedResponseType as protobuf.Type;
    const payload = requestType.encode(data).finish();
    const service = method.parent as protobuf.Service;
    const serviceName= service.fullName.slice(1);
    const req = InvokeRequest.encode({
      method: `/${serviceName}/${method.name}`,
      payload,
    }).finish();



    const res = await fetch(`${this.prefix}/api/invoke?host=${addr}`, {
      method: 'POST',
      body: req,
    });

    const buf = await res.arrayBuffer();

    const decoded = InvokeResponse.decode(new Uint8Array(buf));
    console.log(decoded);

    if (decoded.error) {
      throw new Error(decoded.error.message || 'Failed to invoke method');
    }

    if (!decoded.payload) {
      throw new Error('Failed to invoke method');
    }

    const decodedResponse = responseType.decode(decoded.payload);
    return decodedResponse.toJSON();
  }

}

export default ServerAdapter;
