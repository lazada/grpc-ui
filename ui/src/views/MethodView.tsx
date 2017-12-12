import * as React from 'react';
import { Method } from 'protobufjs';
import MessageInput from './forms/MessageInput';
import { MessageState, createDefaultStateForMessage } from './forms/state';
import JSONTree from 'react-json-tree';
import { Button, Intent } from '@blueprintjs/core';
import Adapter from '../adapter/Adapter';

interface Props {
  obj: Method;
  adapter: Adapter;
  addr: string;
}

interface State {
  state: MessageState;
  response: {} | null;
  running: boolean;
}

const theme = {
  scheme: 'monokai',
  author: 'wimer hazenberg (http://www.monokai.nl)',
  base00: '#272822',
  base01: '#383830',
  base02: '#49483e',
  base03: '#75715e',
  base04: '#a59f85',
  base05: '#f8f8f2',
  base06: '#f5f4f1',
  base07: '#f9f8f5',
  base08: '#f92672',
  base09: '#fd971f',
  base0A: '#f4bf75',
  base0B: '#a6e22e',
  base0C: '#a1efe4',
  base0D: '#66d9ef',
  base0E: '#ae81ff',
  base0F: '#cc6633'
};

class MethodView extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      state: createDefaultStateForMessage(props.obj.resolvedRequestType as protobuf.Type),
      response: null,
      running: false,
    };
  }

  run = async () => {
    this.setState({ running: true });

    const response = await this.props.adapter.runMethod(this.props.addr, this.props.obj, this.state.state.getValue());
    // TODO catch error
    this.setState({
      running: false,
      response,
    });
  }

  render() {
    const obj = this.props.obj;
    const { response } = this.state;
    return (
      <div style={{ padding: 10 }}>
        <div>{obj.parent && obj.parent.fullName.slice(1)}</div>
        <h1>{obj.name}</h1>
        <div style={{ display: 'flex' }}>
          <div style={{ flex: 1 }}>
            <MessageInput
              message={obj.resolvedRequestType as protobuf.Type}
              level={0}
              state={this.state.state}
              onChange={state => {
                this.setState({ state });
              }}
            />
           <div style={{ marginTop: 10 }}>
            <Button
              text="Send request"
              intent={Intent.PRIMARY}
              disabled={!this.state.state.isValid()}
              loading={this.state.running}
              onClick={this.run}
            />
            </div>
          </div>
          <div style={{ flex: 1, paddingTop: 10, paddingLeft: 10 }}>
            {
              response ?
                <JSONTree
                  data={response}
                  hideRoot={true}
                  theme={theme}
                /> : null
            }
          </div>
        </div>
      </div>
    );
  }
}  

export default MethodView;