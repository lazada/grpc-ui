import * as React from 'react';
import { Method, ReflectionObject, Type } from 'protobufjs';
import MethodView from './MethodView';
import MessageView from './MessageView';
import Adapter from '../adapter/Adapter';

export function renderViewForObject(obj: ReflectionObject, adapter: Adapter, addr: string) {
  if (obj instanceof Method) {
    return <MethodView obj={obj} adapter={adapter} addr={addr}/>;
  }

  if (obj instanceof Type) {
    return <MessageView obj={obj}/>;
  }

  return null;
}

export function hasView(obj: ReflectionObject) {
  if (obj instanceof Method) {
    return true;
  }

  if (obj instanceof Type) {
    return true;
  }

  return false;
}