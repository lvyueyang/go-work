import { Breadcrumb, BreadcrumbProps } from 'antd';
import pathToRegexp from 'path-to-regexp';
import React from 'react';
import { useAppData } from 'umi';
import styles from './index.module.less';

interface HeaderProps {
  items?: BreadcrumbProps['items'];
}

interface TreeNode {
  path?: string;
  title?: string;
  children?: TreeNode[];
}

function getParentNodesByKey(
  path?: string,
  tree: TreeNode[] = [],
  parentNodes: TreeNode[] = [],
): Omit<TreeNode, 'children'>[] | void {
  for (const node of tree) {
    const currentNode = { path: node.path, title: node.title };

    if (new RegExp(pathToRegexp(node.path!)).test(path!)) {
      return [...parentNodes, currentNode];
    }
    if (node.children) {
      const result = getParentNodesByKey(path, node.children, [...parentNodes, currentNode]);
      if (result) {
        return result;
      }
    }
  }
  return void 0;
}

export default function Header({ items, children }: React.PropsWithChildren<HeaderProps>) {
  const { clientRoutes } = useAppData();
  const menuRoutes = clientRoutes[0].children?.find((i: any) => i.meta?.isMenuRoot)?.children;
  const currentPathname = location.pathname;
  const list = getParentNodesByKey(currentPathname, menuRoutes);

  return (
    <div className={`${styles.headerContainer} header`}>
      <Breadcrumb
        items={
          items ||
          list?.map((item) => ({
            title: item.title,
          }))
        }
      />
      {children}
    </div>
  );
}

export function HeaderBreadcrumb() {
  const { clientRoutes } = useAppData();
  const menuRoutes = clientRoutes[0].children?.find((i: any) => i.meta?.isMenuRoot)?.children;
  const currentPathname = location.pathname;
  const list = getParentNodesByKey(currentPathname, menuRoutes);

  return (
    <Breadcrumb
      items={list?.map((item) => ({
        title: item.title,
      }))}
    />
  );
}
