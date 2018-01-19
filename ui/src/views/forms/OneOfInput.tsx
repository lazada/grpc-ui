import * as React from 'react';
import { OneOfState } from './state';
import TypeBadge from './TypeBadge';
import { renderInput } from './FieldInput'; // TODO refactor

interface Props {
  oneOf: protobuf.OneOf,
  state: OneOfState,
  onChange: (state: OneOfState) => void,
}

const OneOfInput = ({ oneOf, state, onChange }: Props) => {
  const field = oneOf.fieldsArray.find(field => field.name === state.getCurrent())!;
  return (
    <div style={{ marginBottom: 10}}>
      <div style={{ marginBottom: 5 }}>
        {oneOf.name}
        <TypeBadge field={field}/>
        <div className="pt-select" style={{ marginLeft: 10 }}>
          <select
            value={state.getCurrent()}
            onChange={e => {
              onChange(state.setCurrent(e.target.value));
            }}
          >
            {
              oneOf.fieldsArray.map(field =>
                <option key={field.name} value={field.name}>{field.name}</option>
              )
            }
          </select>
        </div>

      </div>
      {renderInput({
        field,
        state: state.getCurrentFieldState(),
        onChange: (fieldState) => { onChange(state.setCurrentFieldState(fieldState)); },
        level: 1,
      })}
    </div>
  );
};

export default OneOfInput;