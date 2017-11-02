import React, {Component} from 'react';

import Sidebar from './Sidebar';
import Method from './Method';

import './app.sass';
import axios from 'axios';


class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            packages: [],
            types: {},
            enums: {},
        }
    }
    componentDidMount() {
        axios.get('/api/info')
            .then(({data: {packages, types, enums}}) => {
               this.setState({
                   packages, types, enums,
               })
            });
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
                            <Sidebar packages={this.state.packages}/>
                        </div>
                        <div className="app__right">
                            <div className="packages-list">
                                {Object.keys(this.state.packages).map(package_name => {
                                    return this.state.packages[package_name].map((service) => {
                                        return <div className="package">
                                            <h3 className="package__title">{package_name + ' / ' + service.name}</h3>
                                            {service.methods.map((method) =>
                                                    <Method key={method.name}
                                                            {...method}
                                                            service_name={service.name}
                                                            package_name={package_name}
                                                            types={this.state.types}
                                                            enums={this.state.enums}
                                                />
                                            )}
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


export default App;