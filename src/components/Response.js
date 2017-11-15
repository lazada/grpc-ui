import React, {Component} from 'react';
import injectSheet from 'react-jss';
import {getTypeName, TYPE_MESSAGE} from './types';


const styles = {
    response: {
        'padding': '20px',
    },
    title: {
        'font-weight': 'body',
        'font-size': '20px',
        'margin': '0',
        'padding': '0',
        'padding-bottom': '20px',
    }
};

const Response = (props) => {
    const type = props.types[props.type_name];

    return type ?
        <div className={props.classes.response}>
            <h4 className={props.classes.title}>{props.type_name}</h4>

            <table className="message">
                {type.fields.map(f =>
                    <tr className="field">
                        <td className="message__cell message__cell--first">
                            <label className="field__label" htmlFor={f.name}><b>{f.name}</b></label>
                        </td>
                        <td className="message__cell message__cell--last">
                            {f.type_id === TYPE_MESSAGE ? f.type_name : getTypeName(f.type_id)} {f.is_repeated ? '(+)': ''}
                        </td>
                    </tr>
                )}
            </table>
        </div> :
        <div>Unknown type: {props.type_name}</div>
}


export default injectSheet(styles)(Response);