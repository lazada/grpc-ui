import * as React from 'react';
import { parse } from 'protobufjs';
import MethodView from './views/MethodView';

const proto = `
  syntax = "proto3";

  message SearchRequest {
    string query = 1;
    int32 foo = 2;
    bool corpus = 3;
    Foo foooo = 4;
  }

  message Foo {
    string query = 1;
    int32 foo = 2;
  }

  message SearchResponse {
  }

  service SearchService {
    rpc Search (SearchRequest) returns (SearchResponse);
  }
`;

const root = parse(proto).root;
root.resolveAll();
const method = root.lookupService('.SearchService').methodsArray[0];


interface State {
}


class App extends React.Component<{}, State> {
  render() {
    return (
      <div style={{ padding: 10 }}>
        <MethodView obj={method} adapter={{
          runMethod() {
            return new Promise(r => setTimeout(() => { r({ foo: 123 })}, 1000))
          }
        } as any} addr=""/>
      </div>
    );
  }
}


export default App;