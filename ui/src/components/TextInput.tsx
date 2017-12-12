import * as React from 'react';

interface Props {
  value: string;
  onChange: (value: string) => void;
  isValid?: boolean;
}

const TextInput = ({ value, onChange, isValid = true }: Props) =>
  <input
    className={'pt-input pt-fill' + (isValid ? '' : ' pt-intent-danger')}
    value={value}
    onChange={e => {
      onChange(e.target.value);
    }}
  />
;

export default TextInput;