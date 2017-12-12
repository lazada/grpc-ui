import { Type, Field, Enum } from 'protobufjs';
import { isNumberType } from '../../util/proto';

export interface State {
  isValid(): boolean;
  getError(): string;
  getValue(): any;
}

// TODO support specific types of numbers
export class NumberState implements State {
  private value: number;

  constructor(private repr: string) {
    // TODO check concrete number type
    this.value = parseFloat(repr);
  }

  getString() {
    return this.repr;
  }

  setString(repr: string) {
    return new NumberState(repr);
  }

  isValid(): boolean {
    return !isNaN(this.value);
  }
  getError(): string {
    return '';
  }
  getValue() {
    return this.value;
  }
}

export class StringState implements State {
  constructor(private value: string) {}

  setValue(value: string) {
    return new StringState(value);
  }

  isValid(): boolean {
    return true;
  }
  getError(): string {
    return '';
  }
  getValue() {
    return this.value;
  }
}

export class BooleanState implements State {
  constructor(private value: boolean) {}

  isValid(): boolean {
    return true;
  }
  getError(): string {
    return '';
  }
  getValue() {
    return this.value;
  }
}

export class EnumState implements State {
  constructor(private value: number) {}

  isValid() {
    return true;
  }
  getError() {
    return '';
  }
  getValue() {
    return this.value;
  }
}

function copyMap<A, B>(map: Map<A, B>): Map<A, B> {
  const newMap: Map<A, B> = new Map();
  const arr = Array.from(map); // TODO

  for (const [key, val] of arr) {
    newMap.set(key, val);
  }

  return newMap;
}

export class MessageState implements State {
  constructor(private fields: Map<string, State>) {}

  getFieldState(field: string) {
    const fieldState = this.fields.get(field);

    if (!fieldState) {
      throw new Error(`Invalid field: ${field}`);
    }

    return fieldState;
  }

  isValid() {
    const arr = Array.from(this.fields.values()); // TODO

    for (const field of arr) {
      if (!field.isValid()) {
        return false;
      }
    }

    return true;
  }

  getError() {
    return '';
  }

  getValue() {
    const obj: { [key: string]: any } = {};

    this.fields.forEach((field, key) => {
      obj[key] = field.getValue();
    });

    return obj;
  }

  setFieldState(field: string, state: State): MessageState {
    const fields = copyMap(this.fields);
    fields.set(field, state);
    return new MessageState(fields);
  }
}

export function createDefaultStateForMessage(msg: Type): MessageState {
  const fields:  Map<string, State> = new Map();

  for (const field of msg.fieldsArray) {
    fields.set(field.name, createDefaultStateForField(field));
  }

  return new MessageState(fields);
}

function createDefaultStateForField(field: Field): State {
  if (field.resolvedType instanceof Type) {
    return createDefaultStateForMessage(field.resolvedType);
  }

  if (field.resolvedType instanceof Enum) {
    return new EnumState(0);
  }

  if (field.type === 'string') {
    return new StringState('');
  }

  if (field.type === 'bool') {
    return new BooleanState(false);
  }

  if (isNumberType(field.type)) {
    return new NumberState('0');
  }

  // TODO special unsupported state?
  throw new Error(`Unsupported type: ${field.type}`);
}