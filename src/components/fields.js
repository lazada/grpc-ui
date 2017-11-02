import React from 'react';


export const Field = (props) => {
    let input = null;

    switch(props.type_id) {
        case 8:
            input =  <input className="field__input" name={props.name} id={props.name} type="checkbox" checked={props.val === 'true'} onChange={(e) => props.onChange(e.target.checked ? 'true' : 'false')}/>
            break;
        default:
            input = <input className="field__input field__input--text" name={props.name} id={props.name} type="text" value={props.val} onChange={(e) => props.onChange(e.target.value)}/>
            break;
    }

    return <div className="field__group">
        {input}
    </div>;
};

export const RepeatedField = (props) =>
    <div className="field__group">
        {props.val.map((v, i) => <Field
            name={props.name}
            number={props.number}
            type_id={props.type_id}
            val={v}
            onChange={(val) => {
                const newVal = props.val.slice();
                newVal[i] = val;
                props.onChange(newVal);
            }}/> )}
        <div className="field__controls">
            <button type="button" className="button button--small" onClick={() => {props.onChange(props.val.concat([''])); }}>+</button>
        </div>
    </div>;
