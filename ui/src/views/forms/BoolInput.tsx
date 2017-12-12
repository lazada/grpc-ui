import * as React from 'react';
import { BooleanState } from './state';

interface Props {
  state: BooleanState,
  onChange: (state: BooleanState) => void,
}

const EnumInput = ({ state, onChange }: Props) =>
  <div style={{ marginBottom: 5 }}>
    <div className="pt-select">
      <select value={'' + state.getValue()} onChange={e => {
        onChange(new BooleanState(e.target.value === 'true'))
      }}>
         <option value="false">false</option>
         <option value="true">true</option>
      </select>
    </div>
  </div>
;


export default EnumInput;