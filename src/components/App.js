import React, {Component} from 'react';
import {connect} from 'react-redux';

import Sidebar from './Sidebar';
import Method from './Method';
import {loadPackages, invokeMethod} from '../actions';

import './app.sass';

//
// <Package key={package_name} name={package_name}
//          service={service} types={this.props.types}
//          onInvokeMethod={(method_name, args) =>
//              this.props.invokeMethod(package_name, service.name, method_name, args)} />

class App extends Component {
    constructor(props) {
        super(props);
    }
    componentDidMount() {
        this.props.loadPackages();
    }
    render() {
        return (
            <div>
                <div className="navbar">
                    <div className="navbar__container">
                        <a href="" className="logo"/>
                    </div>
                </div>
                <div className="app">
                    <div className="app__container">
                        <div className="app__left">
                            <Sidebar packages={this.props.packages}/>
                        </div>
                        <div className="app__right">
                            <div className="packages-list">
                                {Object.keys(this.props.packages).map(package_name => {
                                    return this.props.packages[package_name].map((service) => {
                                        return <div className="package">
                                            <h3 className="package__title">{package_name + ' / ' + service.name}</h3>
                                            {service.methods.map((method) =>
                                                    <Method key={method.name}
                                                            {...method}
                                                            types={this.props.types}
                                                            onInvokeMethod={(args) => this.props.invokeMethod(package_name, service.name, method.name, args)}/>)}
                                        </div>

                                    });
                                })}
                            </div>
                        </div>

                    </div>
                </div>

            </div>
        );
    }
}


const mapStateToProps = state => {
    return {
        packages: state.packages,
        types: state.types,
    }
};

const mapDispatchToProps = (dispatch) => {
    return {
        loadPackages: () => {
            dispatch(loadPackages())
        },
        invokeMethod: (package_name, service_name, method_name, args) => {
            dispatch(invokeMethod(package_name, service_name, method_name, args))
        },
    };
};

App = connect(mapStateToProps, mapDispatchToProps)(App);

export default App;