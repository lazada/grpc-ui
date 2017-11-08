import React from 'react';
import {getDefaultValue} from './Request';


export const Field = (props) => {
    let input = null;

    switch(props.type_id) {
        case 8:
            input =  <input className="field__input" name={props.name} id={props.name} type="checkbox" checked={props.val === 'true'} onChange={(e) => props.onChange(e.target.checked ? 'true' : 'false')}/>
            break;
        case 11:
            const type = props.types[props.type_name];
            if (!type) {
                return <div>?????</div>;
            }
            input = <Message fields={type.fields} types={props.types} val={props.val} enums={props.enums} onChange={props.onChange}/>
            break;
        case 14:
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

const getTypeName = (type_id) => {
    switch (type_id) {
        case 1:  return "double";
        case 2:  return "float";
        case 3:  return "int64";
        case 4:  return "uint64";
        case 5:  return "int32";
        case 6:  return "fixed64";
        case 7:  return "fixed32";
        case 8:  return "bool";
        case 9:  return "string";
        case 10: return "group";
        case 11: return "message";
        case 12: return "bytes";
        case 13: return "uint32";
        case 14: return "enum";
        case 15: return "sfixed32";
        case 16: return "sfixed32";
        case 17: return "sint32";
        case 18: return "sint64";
        default:
            return '???';
    }
};