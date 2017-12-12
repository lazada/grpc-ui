import * as React from 'react';
import { Type } from 'protobufjs';
import { MessageState } from './state';
import FieldInput from './FieldInput';

interface Props {
  message: Type;
  state: MessageState;
  onChange: (state: MessageState) => void;
  level: number;
}

const MessageInput = ({ message, state, level, onChange }: Props) =>
  <div style={level === 0 ? {} : { borderRadius: 3, boxShadow: 'inset 0 0 0 1px rgba(16, 22, 26, 0.2)', padding: 10 }}>
    {
      message.fieldsArray.map(field =>
        <FieldInput
          key={field.name}
          state={state.getFieldState(field.name)}
          field={field}
          level={level}
          onChange={fieldState => { onChange(state.setFieldState(field.name, fieldState)); }}
        />
      )
    }
  </div>
;

export default MessageInput;