import * as React from 'react';
import TextInput from '../../components/TextInput';
import { StringState } from './state';

interface Props {
  state: StringState,
  onChange: (state: StringState) => void,
}

const StringInput = ({ state, onChange }: Props) =>
  <TextInput
    value={state.getValue()}
    onChange={val => onChange(state.setValue(val))}
  />


export default StringInput;
