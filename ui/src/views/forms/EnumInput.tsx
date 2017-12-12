import * as React from 'react';
import { EnumState } from './state';

interface Props {
  enum: protobuf.Enum,
  state: EnumState,
  onChange: (state: EnumState) => void,
}

const EnumInput = ({ enum: en, state, onChange }: Props) =>
  <div style={{ marginBottom: 5 }}>
    <div className="pt-select">
      <select value={state.getValue()} onChange={e => {
        onChange(new EnumState(Number(e.target.value)));
      }}>
        {
          Object.keys(en.values)
            .map(key =>
              <option value={en.values[key]} key={key}>{key}</option>
            )
        }
      </select>
    </div>
  </div>
;


export default EnumInput;