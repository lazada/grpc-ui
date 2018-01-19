import { Root } from 'protobufjs';
import Adapter from './Adapter';

const sleep = (ms: number) => new Promise(r => setTimeout(r, ms));

export default class MockAdapter implements Adapter {
  async fetchReflection(addr: string): Promise<Root> {
    if (addr.toLowerCase() === 'invalid') {
      throw new Error('Failed to fetch reflection');
    }

    await sleep(1000);
    return this.getMockReflection();
  }

  getMockReflection() {
    const root = new Root();

    return root;
  }

  runMethod(addr: string, method: protobuf.Method, data: any): Promise<{}> {
    console.log(method.resolvedRequestType!.verify(data));
    console.log(data);
    return Promise.resolve(data);
  }
}