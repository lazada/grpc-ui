import * as React from 'react';
import TextInput from '../../components/TextInput';
import { NumberState } from './state';

interface Props {
  state: NumberState,
  onChange: (state: NumberState) => void,
}

const StringInput = ({ state, onChange }: Props) =>
  <TextInput
    value={state.getString()}
    onChange={val => onChange(state.setString(val))}
    isValid={state.isValid()}
  />


export default StringInput;
