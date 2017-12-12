import * as React from 'react';

interface Props {
  obj: protobuf.Type;
}

const MessageView = ({ obj }: Props) =>
  <div style={{ padding: 10 }}>
    <div>{obj.parent && obj.parent.fullName.slice(1)}</div>
    <h1>{obj.name}</h1>
  </div>
;

export default MessageView;