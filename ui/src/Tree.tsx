import * as React from 'react';
import { Tree as BTree, ITreeNode, InputGroup, Classes } from '@blueprintjs/core';
import { TreeNode, filterTree } from './util/tree';
import { hasView } from './views/main';

interface Props {
  tree: TreeNode;
  selected: string | null;
  onSelect: (node: string) => void;
}

interface State {
  collapsedNodes: Set<string>;
  selected: string | null;
  filterTerm: string;
}

const renderBadge = (node: TreeNode) => {
  const style: React.CSSProperties = {
    display: 'inline-block',
    width: 12,
    height: 12,
    borderRadius: 6,
    marginRight: 4,
    textAlign: 'center',
    verticalAlign: 'middle',
    lineHeight: '13px',
    fontSize: 8,
    color: 'white',
    marginTop: -2,
  };

  if (node.type === 'Message') {
    style.backgroundColor = '#FF851B';
    return <div style={style}>M</div>;
  }

  if (node.type === 'Enum') {
    style.backgroundColor = '#FFDC00';
    return <div style={style}>E</div>;
  }

  if (node.type === 'Service') {
    style.backgroundColor = '#0074D9';
    return <div style={style}>S</div>;
  }

  return null;
};

class Tree extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      collapsedNodes: new Set(),
      selected: props.selected,
      filterTerm: '',
    };
  }

  getContents(): ITreeNode[] {
    const filtered = filterTree(this.props.tree, this.state.filterTerm);

    if (!filtered) {
      return [];
    }

    return filtered.children.map(n => this.convertNode(n));
  }

  convertNode(node: TreeNode): ITreeNode {
    const id = node.obj.fullName;
    return {
      id,
      label: <span>{renderBadge(node)}{node.name}</span>,
      isSelected: this.state.selected === id,
      className: hasView(node.obj) ? 'selectable-node' : undefined,
      isExpanded: !this.state.collapsedNodes.has(id),
      hasCaret: node.type === 'Package',
      childNodes: node.children.map(n => this.convertNode(n)),
    };
  }

  render() {
    return (
      <div>
        <div style={{ padding: 10 }}>
          <InputGroup
            className={Classes.ROUND}
            placeholder="Search"
            leftIconName="search"
            onChange={(e: React.FormEvent<HTMLInputElement>) => {
              this.setState({ filterTerm: (e.target as HTMLInputElement).value });
            }}
          />
        </div>
      <BTree
        contents={this.getContents()}
        onNodeExpand={node => {
          const collapsedNodes = this.state.collapsedNodes;
          collapsedNodes.delete(node.id as string);
          this.setState({ collapsedNodes });
        }}
        onNodeCollapse={node => {
          const collapsedNodes = this.state.collapsedNodes;
          collapsedNodes.add(node.id as string);
          this.setState({ collapsedNodes });
        }}
        onNodeClick={node => {
          if (node.className === 'selectable-node') { // TODO WTF
            this.setState({
              selected: node.id as string,
            });
            this.props.onSelect(node.id as string);
          }
        }}
      />
      </div>
    );
  }
}

export default Tree;