import * as React from 'react';
import { Tree as BTree, ITreeNode as BlueprintITreeNode, InputGroup, Classes } from '@blueprintjs/core';
import { TreeNode, filterTree } from './util/tree';
import { hasView } from './views/main';

interface TreeNodeProps extends BlueprintITreeNode {
  treeNode: TreeNode | null; 
}

interface Props {
  tree: TreeNode;
  selected: TreeNode | null;
  onSelect: (node: TreeNode) => void;
}

interface State {
  collapsedNodes: Set<number | string>;
  selected: TreeNode | null;
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

  getContents(): TreeNodeProps[] {
    return [
      {
        id: '@services',
        label: 'Services',
        isExpanded: !this.state.collapsedNodes.has('@services'),
        childNodes: this.getServices(),
        treeNode: null,
      },
      {
        id: '@types',
        label: 'Types',
        isExpanded: !this.state.collapsedNodes.has('@types'),
        childNodes: this.getTypes(),
        treeNode: null,
      }
    ];

  }

  getServices(): TreeNodeProps[] {
    const term = this.state.filterTerm.toLowerCase();

    const filtered = filterTree(this.props.tree, n =>
      n.type === 'Service' &&  n.name.toLowerCase().includes(term)
    );

    if (!filtered) {
      return [];
    }

    return filtered.children.map(n => this.convertNode(n, '@services.'));
  }

  getTypes(): TreeNodeProps[] {
    const term = this.state.filterTerm.toLowerCase();

    const filtered = filterTree(this.props.tree, n =>
      (n.type === 'Message' || n.type === 'Enum') && n.name.toLowerCase().includes(term)
    );

    if (!filtered) {
      return [];
    }

    return filtered.children.map(n => this.convertNode(n, '@types.'));
  }

  convertNode(node: TreeNode, prefix: string): TreeNodeProps {
    const id = prefix + '.' + node.obj.fullName;

    return {
      id,
      label: <span>{renderBadge(node)}{node.name}</span>,
      isSelected: this.state.selected === node,
      className: hasView(node.obj) ? 'selectable-node' : undefined,
      isExpanded: !this.state.collapsedNodes.has(id),
      hasCaret: node.type === 'Package',
      childNodes: node.children.map(n => this.convertNode(n, prefix)),
      treeNode: node,
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
          collapsedNodes.delete(node.id);
          this.setState({ collapsedNodes });
        }}
        onNodeCollapse={node => {
          const collapsedNodes = this.state.collapsedNodes;
          collapsedNodes.add(node.id);
          this.setState({ collapsedNodes });
        }}
        onNodeClick={(node: TreeNodeProps) => {
          if (node.treeNode && hasView(node.treeNode.obj)) {
            this.setState({
              selected: node.treeNode,
            });
            this.props.onSelect(node.treeNode);
          }
        }}
      />
      </div>
    );
  }
}

export default Tree;