import React from 'react';

const Obj = ({ obj }) =>
  <div>
    <h1 style={{ marginBottom: 5 }}>{obj.name}</h1>
    <div>{obj.fullName}</div>
  </div>

export default Obj;