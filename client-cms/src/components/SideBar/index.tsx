import { cls } from '@/utils';
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { Menu } from 'antd';
import { useEffect, useState } from 'react';
import { history } from 'umi';
import { getDefaultOpenKeys, getNavMenu } from './getNavMenu';
import styles from './index.module.less';

const menuItems = getNavMenu();

function useSelectMenu() {
  const [selectKeys, setSelectKeys] = useState<string[]>(() => {
    return [location.pathname];
  });
  const [openKeys, setOpenKeys] = useState<string[]>(() => {
    return getDefaultOpenKeys();
  });
  const [collapsed, setCollapsed] = useState(false);

  const toggleCollapsed = () => {
    setCollapsed(!collapsed);
  };
  useEffect(() => {
    const openKeys = getDefaultOpenKeys();
    setOpenKeys(openKeys);
    setSelectKeys(openKeys);
  }, [location.pathname]);

  return {
    collapsed,
    openKeys,
    selectKeys,
    setSelectKeys,
    toggleCollapsed,
    setOpenKeys,
  };
}

export default function SideBar() {
  const { collapsed, openKeys, selectKeys, setSelectKeys, toggleCollapsed, setOpenKeys } =
    useSelectMenu();
  return (
    <div className={cls(styles.sideBarContainer, collapsed && styles.collapsed)}>
      <Menu
        inlineCollapsed={collapsed}
        mode="inline"
        className={styles.menuList}
        items={menuItems}
        openKeys={openKeys}
        selectedKeys={selectKeys}
        onClick={(e) => {
          setSelectKeys([e.key]);
          history.push(e.key);
        }}
        onOpenChange={(e) => {
          setOpenKeys(e);
        }}
      />
      <div className={styles.operate}>
        <div className={styles.item} onClick={toggleCollapsed}>
          {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
        </div>
      </div>
    </div>
  );
}
