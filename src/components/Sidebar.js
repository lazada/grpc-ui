import React from 'react';
import { Root, Namespace, Service, Type, Enum } from 'protobufjs';
import './Sidebar.sass';



const getObjectBadge = obj => {
  if (obj instanceof Type) {
    return { letter: 'M', color: '#FF851B' };
  }

  if (obj instanceof Service) {
    return { letter: 'S', color: '#0074D9' };
  }

  if (obj instanceof Enum) {
    return { letter: 'E', color: '#FFDC00' };
  }
}



const renderBadge = obj => {
  const badge = getObjectBadge(obj);

  if (!badge) {
    return (
      <div style={{
        display: 'inline-block',
        width: 6,
       }}></div>
    );
  }

  return (
    <div style={{
      display: 'inline-block',
      backgroundColor: badge.color,
      width: 12,
      height: 11,
      color: 'white',
      fontSize: 8,
      lineHeight: '12px',
      fontKerning: 'normal',
      textAlign: 'center',
      verticalAlign: 'middle',
      borderRadius: 6,
      paddingTop: 1,
      marginRight: 2,
      marginTop: -3,
     }}>{badge.letter}</div>
  );
}




const renderItem = (item, level, onClick) => {
  if (item instanceof Root) {
    return (
      <div>
        {item.nestedArray.map(item => renderItem(item, level, onClick))}
      </div>
    );
  }

  let nested = null;

  if (item instanceof Service) {
    nested = (
      <div>
        {item.methodsArray.map(item => renderItem(item, level + 1, onClick))}
      </div>
    );
  } else if (item instanceof Namespace) {
    nested = (
      <div>
        {item.nestedArray.map(item => renderItem(item, level + 1, onClick))}
      </div>
    );
  }

  return (
    <div key={item.fullName} style={{ paddingLeft: level * 10, marginTop: 7 }}>
      <div>{renderBadge(item)}<a href="#" className="sidebar__link" onClick={e => {
        e.preventDefault();
        onClick(item);
      }}>{item.name}</a></div>
      {nested}
    </div>
  );
};

const Sidebar = ({ reflection, onSelect }) =>
  <div>
    {reflection ? renderItem(reflection, 0, onSelect) : null}
  </div>
;

export default Sidebar;
