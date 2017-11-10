import axios from 'axios';
import proto from './proto.js';
import { Root } from 'protobufjs';
import { convertFileDescriptor, addFromFileDescriptor, resolveAll } from 'proto-descriptor';

class Adapter {
  fetchReflection() {
    return axios.get('/api/reflection', { responseType: 'arraybuffer' })
      .then(res => {
        const msg = proto.ReflectionResponse.decode(new Uint8Array(res.data));

        if (msg.error) {
          const error = msg.error;
          throw new Error(error.message);
        }

        const reflection = msg.reflection;
        const root = new Root();

        for (const fd of reflection.fileDescriptor) {
          addFromFileDescriptor(root, fd);
        }

        resolveAll(root);
        return root;
      });
  }

}

export default Adapter;