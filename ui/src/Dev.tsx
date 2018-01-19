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

  message Empty {
  }

  message OneOfRequest {
    string before = 1;

    oneof test_oneof {
      string option_a = 2;
      int32 option_b = 3;
    }

    string after = 4;
  }

  service SearchService {
    rpc Search (SearchRequest) returns (Empty);
    rpc TestOneOf(OneOfRequest) returns (Empty);
  }
`;

const root = parse(proto).root;
const adapter = new MockAdapter();

root.resolveAll();

const Dev = () =>
  <MainState reflection={root} adapter={adapter} addr={''}/>
;

export default Dev;