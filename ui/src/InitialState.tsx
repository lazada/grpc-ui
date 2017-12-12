import * as React from 'react';
import { Root } from 'protobufjs';
import Adapter from './adapter/Adapter'
import { NonIdealState, InputGroup, Button, Tooltip, Classes, Intent, Spinner, Position, Toaster } from '@blueprintjs/core';


interface InitData {
  reflection: Root,
  address: string,
}

interface Props {
  adapter: Adapter,
  onInit: (data: InitData) => void,
}

interface State {
  address: string | void,
  connecting: boolean,
}

const OurToaster = Toaster.create({
  position: Position.TOP,
});

class InitialState extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      address: undefined,
      connecting: false,
    };
  }

  focusInput = (el: HTMLInputElement | void) => {
    if (el) {
      el.focus();
    }
  }

  async connect() {
    const address = this.state.address;

    if (!address) {
      return;
    }

    this.setState({ connecting: true });

    try {
      const reflection = await this.props.adapter.fetchReflection(address);

      this.props.onInit({
        reflection,
        address,
      });

    } catch (e) {
      this.setState({ connecting: false });
      OurToaster.show({
        message: 'Failed to connect to server',
        intent: Intent.DANGER,
      });
    }
  }

  render() {
    return (
      <div style={{ position: 'absolute', width: '100%', height: '100%' }}>
        <NonIdealState
          title="Start"
          visual="folder-open"
          description={this.renderDescription()}
        />
      </div>
    );
  }

  renderDescription() {
    const connecting = this.state.connecting;


    let rightElement;

    if (connecting) {
      rightElement = <Spinner className={Classes.SMALL}/>;
    } else {
      const disabled = !this.state.address;
      rightElement = (
        <Tooltip content="Connect" isDisabled={disabled}>
            <Button
              className={Classes.MINIMAL}
              intent={Intent.WARNING}
              iconName="arrow-right"
              disabled={disabled}
              onClick={() => this.connect()}
            />
        </Tooltip>
      );
    }

    return (
      <div>
        <form onSubmit={e => {
          e.preventDefault();
          this.connect()
        }}>
          <InputGroup
            inputRef={this.focusInput}
            placeholder="GRPC server address"
            disabled={connecting}
            rightElement={rightElement}
            onChange={(e: React.FormEvent<HTMLInputElement>) => {
              this.setState({ address: (e.target as any).value })
            }}
            value={this.state.address || ''}
          />
        </form>
      </div>
    );
  }

}

export default InitialState;
