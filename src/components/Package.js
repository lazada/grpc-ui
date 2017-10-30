import React from 'react';
import Method from './Method';

const Package = (props) => (
    <div className="package">
        <h3 className="package__title">{props.name + ' / ' + props.service.name}</h3>
        {props.service.methods.map((method) => <Method key={method.name} {...method} types={props.types}/>)}
    </div>
);

export default Package;