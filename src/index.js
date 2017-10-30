import React, {Component} from 'react';
import ReactDOM from 'react-dom';


const Int32Field = (props) => <input name={props.name} id={props.name} type="number"/>;

const StringField = (props) => <input name={props.name} id={props.name} type="text"/>;

const BooleanField = (props) => <input name={props.name} id={props.name} type="checkbox"/>;

const UnknownField = (props) => <input name={props.name} id={props.name} type="text" placeholder={'type id: ' + props.type_id}/>;

class RepeatedField extends Component {
    constructor(props) {
        super(props);
        this.state = {
          count: 3,
        };
    }
    render() {
        const children = [];
        for(let i = 0; i < this.state.count; i++) {
            children.push(this.props.children);
        }
        return <div>
            {children}
        </div>
    }
}

const MessageField = (props) => {
    const type = props.types[props.type_name];
    return <Message {...type} types={props.types}/>
};

const Field = (props) => {
    if (props.label === 3) { //LABEL_REPEATED
        return <RepeatedField>
                    <Field {...props} label={1}/>
               </RepeatedField>
    }
    switch (props.type_id) {
        case 5: // int32:
            return <Int32Field {...props}/>;
        case 8: // boolean
            return <BooleanField {...props}/>;
        case 9: // int32:
            return <StringField {...props}/>;
        case 11: // message
            return <MessageField {...props}/>;
        default:
            return <UnknownField  {...props}/>;
    }
};

const Message = (props) =>
    <div>
        <h4>{props.name}</h4>
        <hr/>
        <table className="table">
            <tr>
                <th>Name</th>
                <th>Input</th>
            </tr>
            {props.fields.map((f) => <tr>
                <td>
                    <label htmlFor={f.name}>{f.name}</label>
                </td>
                <td>
                    <Field {...f} types={props.types} />
                </td>
            </tr>)}

        </table>
    </div>;

const Method = (props) =>
    <div className="panel panel-default">
        <div className="panel-heading">
            <h4>{props.name}</h4>
            <div className="well">
                <Message {...props.types[props.in]} types={props.types}/>
            </div>
            <div className="well">
                <Message {...props.types[props.out]} types={props.types}/>
            </div>
        </div>
        <div className="panel-body">

        </div>
    </div>;

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            packages: {},
            types: {},
        };
    }
    componentDidMount() {
        const addr = '127.0.0.1:3001';
        fetch('/api/info?addr=' + addr)
            .then(r => r.json())
            .then(({packages, types}) => {
                this.setState({
                    packages,
                    types,
                })
            });
    }
    render() {
        const packages = Object.keys(this.state.packages).map(package_name => {
            return this.state.packages[package_name].map((service) => {
                return <div>
                    <h3>{package_name + ' / ' + service.name}</h3>
                    {service.methods.map((method) => <Method {...method} types={this.state.types}/>)}
                </div>
            });
        });


        return (
            <div>
                <div className="container">
                    <h1>GRPC UI</h1>
                    <div className="col-md-2">

                    </div>
                    <div className="col-md-10">
                        {packages}
                    </div>
                </div>
            </div>
        );
    }
}

ReactDOM.render(<App />, document.getElementById('root'));