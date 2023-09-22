import React from 'react';
import styles from './index.module.less';
import SideBar from '../SideBar';
import HeaderBar from '../HeaderBar';

interface IProps {
  children: React.ReactNode;
}

export default function Layout({ children }: IProps) {
  return (
    <>
      <HeaderBar />
      <div style={{ display: 'flex' }}>
        <SideBar />
        <div className={styles.content}>{children}</div>
      </div>
    </>
  );
}
