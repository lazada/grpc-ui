import * as React from 'react';
import { Type } from 'protobufjs';
import { MessageState } from './state';
import FieldInput from './FieldInput';
import OneOfInput from './OneOfInput'

interface Props {
  message: Type;
  state: MessageState;
  onChange: (state: MessageState) => void;
  level: number;
}

type Field =
  | { t: 'Simple', field: protobuf.Field }
  | { t: 'OneOf', oneOf: protobuf.OneOf }
;

function extractFields(msg: protobuf.Type): Array<Field> {
  const result: Array<Field> = [];
  const oneOfsSeen: Set<protobuf.OneOf> = new Set();

  for (const field of msg.fieldsArray) {
    const oneOf = field.partOf;
    if(oneOf) {
      if (!oneOfsSeen.has(oneOf)) {
        oneOfsSeen.add(oneOf);
        result.push({ t: 'OneOf', oneOf });
      }

    } else {
      result.push({ t: 'Simple', field });
    }
  }

  return result;
}

const MessageInput = ({ message, state, level, onChange }: Props) =>
  <div style={level === 0 ? {} : { borderRadius: 3, boxShadow: 'inset 0 0 0 1px rgba(16, 22, 26, 0.2)', padding: 10 }}>
    {
      extractFields(message).map(field =>
        field.t === 'Simple' ?

        <FieldInput
          key={field.field.name}
          state={state.getFieldState(field.field.name)}
          field={field.field}
          level={level}
          onChange={fieldState => { onChange(state.setFieldState(field.field.name, fieldState)); }}
        />

        : <OneOfInput
            key={field.oneOf.name}
            oneOf={field.oneOf}
            state={state.getOneOfState(field.oneOf.name)}
            onChange={oneOfState => { onChange(state.setOneOfState(field.oneOf.name, oneOfState)); }}
          />
      )
    }
  </div>
;

export default MessageInput;