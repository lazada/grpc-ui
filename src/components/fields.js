import React from 'react';
import {TYPE_BOOL, TYPE_ENUM, TYPE_INT32, TYPE_MESSAGE, getTypeName, getDefaultValue} from './types';

export const Field = (props) => {
    let input = null;

    switch(props.type_id) {
        case TYPE_BOOL:
            input =  <input className="field__input" name={props.name} id={props.name} type="checkbox" checked={props.val === 'true'} onChange={(e) => props.onChange(e.target.checked ? 'true' : 'false')}/>
            break;
        case TYPE_INT32:
            input =  <input className="field__input" name={props.name} id={props.name} type="number" value={props.val} onChange={(e) => props.onChange(e.target.value)}/>
            break;
        case TYPE_MESSAGE:
            const type = props.types[props.type_name];
            if (!type) {
                return <div>?????</div>;
            }
            input = <Message fields={type.fields} types={props.types} val={props.val} enums={props.enums} onChange={props.onChange}/>
            break;
        case TYPE_ENUM:
            const enum_ = props.enums[props.type_name];
            if (!enum_) {
                return <div>?????</div>;
            }
            input = <select className="field__input" value={props.val} onChange={(e) => props.onChange(e.target.value)}>
                {Object.keys(enum_.values).map(k => <option value={k}>{enum_.values[k]}</option>)}
            </select>;
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
            type_name={props.type_name}
            types={props.types}
            enums={props.enums}
            val={v}
            onChange={(val) => {
                const newVal = props.val.slice();
                newVal[i] = val;
                props.onChange(newVal);
            }}/> )}
        <div className="field__controls">
            {props.val.length ? <button type="button" className="button button--small" onClick={() => {props.onChange(props.val.slice(0, props.val.length - 1)); }}>-</button>: null}
            <button type="button" className="button button--small" onClick={() => {props.onChange(props.val.concat([getDefaultValue(props.type_id, false, props.type_name, props.enums, props.types)])); }}>+</button>
        </div>
    </div>;

export const Message = (props) =>
    <table className="message">
        {props.fields.map((f, i) =>
            <tr className="field">
                <td className="message__cell message__cell--first">
                    <label className="field__label" htmlFor={f.name}><b>{f.name}</b></label>
                </td>
                <td className="message__cell">
                    {f.is_repeated ?
                        <RepeatedField name={f.name}
                                       number={f.number}
                                       val={props.val[i]}
                                       type_id={f.type_id}
                                       type_name={f.type_name}
                                       types={props.types}
                                       enums={props.enums}
                                       onChange={(val) => {
                                           const newArr = props.val.slice();
                                           newArr[i] = val;
                                           props.onChange(newArr);
                                       }}/>
                        :
                        <Field name={f.name}
                               number={f.number}
                               val={props.val[i]}
                               type_id={f.type_id}
                               type_name={f.type_name}
                               types={props.types}
                               enums={props.enums}
                               onChange={(val) => {
                                   const newArr = props.val.slice();
                                   newArr[i] = val;
                                   props.onChange(newArr);
                               }} />
                    }
                </td>
                <td className="message__cell message__cell--last">
                    {getTypeName(f.type_id)} {f.is_repeated ? '(+)': ''}
                </td>
            </tr>
        )}
    </table>;
