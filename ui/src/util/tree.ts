import { Root, ReflectionObject, Namespace, Type, Enum, Service, Method } from 'protobufjs';

export type NodeType =
  | 'Root'
  | 'Package'
  | 'Message'
  | 'Service'
  | 'Method'
  | 'Enum'
  | 'Unknown'
;

export interface TreeNode {
  type: NodeType;
  obj: ReflectionObject;
  name: string;
  children: TreeNode[];
}

function convertObject(obj: ReflectionObject): TreeNode | null {
  const node: TreeNode = {
    type: 'Unknown',
    obj,
    name: obj.name,
    children: [],
  };

  if (obj instanceof Type) {
    node.type = 'Message';
    node.children = convertChildren(obj);
    return node;
  }

  if (obj instanceof Enum) {
    node.type = 'Enum';
    return node;
  }

  if (obj instanceof Root) {
    node.type = 'Root';
    node.children = convertChildren(obj);
    return node;
  }

  if (obj instanceof Service) {
    node.type = 'Service';
    node.children = obj.methodsArray.map(convertObject) as TreeNode[];
    return node;
  }

  if (obj instanceof Method) {
    node.type = 'Method';
    return node;
  }

  if (obj instanceof Namespace) {
    node.type = 'Package';
    node.children = convertChildren(obj);
    return node;
  }

  return null;
}

function convertChildren(obj: Namespace): TreeNode[] {
  const result: TreeNode[] = [];

  for (const child of obj.nestedArray) {
    const node = convertObject(child);

    if (node) {
      result.push(node);
    }

  }

  return result;
}

function flattenPackages(node: TreeNode): TreeNode {
  const { type, children, name, obj } = node;

  if (type === 'Package' && children.length === 1 && children[0].type === 'Package') {
    const flattenedChild = flattenPackages(children[0]);

    return {
      type,
      obj: flattenedChild.obj,
      name: name + '.' + flattenedChild.name,
      children: flattenedChild.children,
    };
  }

  return {
    type,
    name,
    obj,
    children: children.map(flattenPackages),
  };
}

export function buildTree(root: Root): TreeNode {
  const node = convertObject(root);  

  if (!node) {
    return {
      type: 'Root',
      obj: root,
      name: '',
      children: [],
    };
  }

  return flattenPackages(node);
}

export function filterTree(node: TreeNode, term: string): TreeNode | null {
  if (!term || node.name.toLowerCase().includes(term)) {
    return node;
  }

  term = term.toLowerCase();

  const {
    type,
    obj,
    name,
    children
  } = node;

  const filteredChildren: TreeNode[] = [];

  for (const child of children) {
    const filteredChild = filterTree(child, term);

    if (filteredChild) {
      filteredChildren.push(filteredChild);
    }
  }

  if (filteredChildren.length > 0) {
    return {
      type,
      obj,
      name,
      children: filteredChildren
    };
  }

  return null;
}

export function getDefaultNode(node: TreeNode): TreeNode | null {
  if (node.type === 'Service' && node.children.length > 0) {
    return node.children[0];
  }

  for (const child of node.children) {
    const defaultNode = getDefaultNode(child);

    if (defaultNode) {
      return defaultNode;
    }
  }

  return null;
}