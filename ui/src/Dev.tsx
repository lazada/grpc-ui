import * as React from 'react';
import { parse } from 'protobufjs';

import MainState from './MainState';
import MockAdapter from './adapter/MockAdapter';

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
const adapter = new MockAdapter();

root.resolveAll();

const Dev = () =>
  <MainState reflection={root} adapter={adapter} addr={''}/>
;

export default Dev;