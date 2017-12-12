import * as React from 'react';

interface Props {
  field: protobuf.Field;
}

const style: { style: React.CSSProperties } = {
  style: {
    display: 'inline-block',
    backgroundColor: '#0074D9',
    color: '#fff',
    borderRadius: 3,
    padding: '0px 2px',
    fontSize: 10,
    marginLeft: 5,
    top: -5,
    position: 'relative',
    lineHeight: '14px',
  }
};

const TypeBadge = ({field}: Props) =>
  <span {...style}>{field.type}</span>
;

export default TypeBadge;