import * as React from 'react';
import TypeBadge from './TypeBadge';

interface Props {
  field: protobuf.Field,
}

const FieldHeader = ({ field }: Props) =>
  <div style={{ marginBottom: 5 }}>
    {field.name}<TypeBadge field={field}/>
  </div>
;


export default FieldHeader;