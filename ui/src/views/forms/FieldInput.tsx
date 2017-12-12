import * as React from 'react';
import { Field, Type, Enum } from 'protobufjs';
import FieldHeader from './FieldHeader';
import { State, MessageState, StringState, NumberState, EnumState, BooleanState } from './state';
import MessageInput from './MessageInput';
import StringInput from './StringInput';
import NumberInput from './NumberInput';
import EnumInput from './EnumInput';
import BoolInput from './BoolInput';
import { isNumberType } from '../../util/proto';

interface Props {
  field: Field;
  state: State;
  onChange: (state: State) => void;
  level: number;
}

function renderInput({ field, state, level, onChange }: Props): JSX.Element {
  if (field.resolvedType instanceof Type) {
    if (!(state instanceof MessageState)) {
      throw new Error('Internal error');
    }

    return (
      <MessageInput
        message={field.resolvedType}
        state={state}
        level={level + 1}
        onChange={onChange}
      />
    );
  }

  if (field.resolvedType instanceof Enum) {
    if (!(state instanceof EnumState)) {
      throw new Error('Internal error');
    }

    return (
      <EnumInput
        state={state}
        enum={field.resolvedType}
        onChange={onChange}
      />
    );
  }

  if (field.type === 'string') {
    if (!(state instanceof StringState)) {
      throw new Error('Internal error');
    }

    return (
      <StringInput
        state={state}
        onChange={onChange}
      />
    );
  }

  if (isNumberType(field.type)) {
    if (!(state instanceof NumberState)) {
      throw new Error('Internal error');
    }

    return (
      <NumberInput
        state={state}
        onChange={onChange}
      />
    );
  }

  if (field.type === 'bool') {
    if (!(state instanceof BooleanState)) {
      throw new Error('Internal error');
    }

    return (
      <BoolInput
        state={state}
        onChange={onChange}
      />
    );
  }

  throw new Error(`Unsupported type" ${field.type}`);
}

const FieldInput = (props: Props) => {
  return (
    <div style={{ marginBottom: 10 }}>
      <FieldHeader field={props.field}/>
      <div>
        {renderInput(props)}
      </div>
    </div>
  );
};

export default FieldInput;