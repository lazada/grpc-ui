import * as React from 'react';
import Tree from './Tree';
import { Root } from 'protobufjs';
import { buildTree, TreeNode, getDefaultNode } from './util/tree';
import { renderViewForObject } from './views/main';
import Adapter from './adapter/Adapter';

interface Props {
  reflection: Root;
  adapter: Adapter;
  addr: string;
}

interface State {
  reflection: Root;
  tree: TreeNode;
  currentNode: TreeNode | null;
}

class MainState extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    const tree = buildTree(props.reflection);

    this.state = {
      reflection: props.reflection,
      tree,
      currentNode: getDefaultNode(tree),
    };
  }

  render() {
    return (
      <div>
      <nav className="pt-navbar .modifier">
        <div className="pt-navbar-group pt-align-left">
          <div className="pt-navbar-heading">GRPC-UI</div>
        </div>
      </nav>
      <div style={{ display: 'flex', marginTop: 20 }}>
        <div style={{ width: 300 }}>
          <Tree
            tree={this.state.tree}
            selected={this.state.currentNode}
            onSelect={node => {
              this.setState({ currentNode: node });
            }}
          />
        </div>
        <div style={{ flex: 1 }}>
          {this.state.currentNode ?
            <div key={this.state.currentNode.obj.fullName}>
              {renderViewForObject(this.state.currentNode.obj, this.props.adapter, this.props.addr)}
            </div> : null}
        </div>
      </div>
    </div>
    );
  }
}

export default MainState;
