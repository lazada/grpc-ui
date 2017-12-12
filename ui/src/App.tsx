import * as React from 'react';
import { Root } from 'protobufjs';

import './App.css';
import InitialState from './InitialState';
import MainState from './MainState';
import ServerAdapter from './adapter/ServerAdapter';


let adapter = new ServerAdapter('');


interface Props {}

interface State {
  address: string | void,
  reflection: Root | void,
}

class App extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      address: undefined,
      reflection: undefined,
    };
  }

  render() {
    const { reflection } = this.state;
    if (reflection) {
      return (
        <MainState reflection={reflection} adapter={adapter} addr={this.state.address as string}/>
      );
    }

    return (
      <InitialState adapter={adapter} onInit={data => {
        this.setState({
          address: data.address,
          reflection: data.reflection,
        });
      }}/>
    );
  }
}

export default App;
