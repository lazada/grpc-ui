import React, {Component} from 'react';
import {Field, RepeatedField} from './fields';

export default class Request extends Component {
    constructor(props) {
        super(props);
        this.state = {
            data: props.fields.map((f) => {
                return f.is_repeated ? [] : f.default_value || getDefaultValue(f.type_id);
            })
        };
    }
    handleInvokeMethod(e) {
        e.preventDefault();
        const args = this.props.fields.map((f, i) => {
            return {
                number: f.number,
                val: this.state.data[i],
            };
        });
        this.props.onInvokeMethod(args);
    }

    handleChange(i, val) {
        let data = this.state.data.slice();
        data[i] = val;
        this.setState({
            data: data,
        });
    }

    render() {
        return (
            <div className="message message--in">
                <h4 className="message__title">Request</h4>
                <form onSubmit={this.handleInvokeMethod.bind(this)}>
                    <table className="message__table">
                        {this.props.fields.map((f, i) =>
                            <tr className="field">
                                <td style={{'width': '100px'}} className="message__cell">
                                    <label className="field__label" htmlFor={f.name}><b>{f.name}</b></label>
                                </td>

                                <td className="message__cell">
                                    {f.is_repeated ?
                                        <RepeatedField name={f.name}
                                                       number={f.number}
                                                       val={this.state.data[i]}
                                                       type_id={f.type_id}
                                                       onChange={(val) => this.handleChange(i, val)}/>
                                        :
                                        <Field name={f.name}
                                               number={f.number}
                                               val={this.state.data[i]}
                                               type_id={f.type_id}
                                               onChange={(val) => this.handleChange(i, val)} />
                                    }
                                </td>
                                <td style={{'text-align': 'right', 'width': '100px'}} className="message__cell">
                                    {getTypeName(f.type_id)} {f.is_repeated ? '(+)': ''}
                                </td>
                            </tr>
                        )}
                    </table>
                    <div className="message__controls">
                        <button type="submit" className="button">Invoke</button>
                    </div>
                </form>
            </div>
        );
    }
}

const getDefaultValue = (type_id) => {
    switch (type_id) {
        case 8:  return 'false';
        default:
            return '';
    }
};

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