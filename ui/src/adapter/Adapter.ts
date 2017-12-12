import { Root } from 'protobufjs';

export default interface Adapter {
  fetchReflection(addr: string): Promise<Root>,
  runMethod(addr: string, method: protobuf.Method, data: any): Promise<{}>,
}