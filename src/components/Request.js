import React, {Component} from 'react';


const Field = (props) => {
    switch(props.type_id) {
        case 8:
            console.log(props);

            return <input className="field__input" name={props.name} id={props.name} type="checkbox" checked={props.val === 'true'} onChange={(e) => props.onChange(e.target.checked ? 'true' : 'false')}/>
        default:
            return <input className="field__input" name={props.name} id={props.name} type="text" value={props.val} onChange={(e) => props.onChange(e.target.value)}/>

    }
}

const RepeatedField = (props) =>
    <div>
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
                    {this.props.fields.map((f, i) => {
                            return <div className="field">
                                <label className="field__label" htmlFor={f.name}>{f.name} {getTypeName(f.type_id)} {f.is_repeated ? '(repeated)': ''} = {f.number} </label>
                                {f.is_repeated ?
                                    <RepeatedField name={f.name}
                                                          number={f.number}
                                                          val={this.state.data[i]}
                                                          type_id={f.type_id}
                                                          onChange={(val) => this.handleChange(i, val)}/>
                                    :
                                    <Field  name={f.name}
                                                   number={f.number}
                                                   val={this.state.data[i]}
                                                   type_id={f.type_id}
                                                   onChange={(val) => this.handleChange(i, val)} />
                                }
                            </div>
                    })}
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
      case 1:  return "TYPE_DOUBLE";
      case 2:  return "TYPE_FLOAT";
      case 3:  return "TYPE_INT64";
      case 4:  return "TYPE_UINT64";
      case 5:  return "TYPE_INT32";
      case 6:  return "TYPE_FIXED64";
      case 7:  return "TYPE_FIXED32";
      case 8:  return "TYPE_BOOL";
      case 9:  return "TYPE_STRING";
      case 10: return "TYPE_GROUP";
      case 11: return "TYPE_MESSAGE";
      case 12: return "TYPE_BYTES";
      case 13: return "TYPE_UINT32";
      case 14: return "TYPE_ENUM";
      case 15: return "TYPE_SFIXED32";
      case 16: return "TYPE_SFIXED64";
      case 17: return "TYPE_SINT32";
      case 18: return "TYPE_SINT64";
      default:
          return '???';
  }
};