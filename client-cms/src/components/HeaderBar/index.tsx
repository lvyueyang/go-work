import LOGO from '@/assets/logo.png';
import { outLogin } from '@/services';
import { useUserinfoStore } from '@/store/userinfo';
import { DownOutlined, FullscreenExitOutlined, FullscreenOutlined } from '@ant-design/icons';
import { useFullscreen } from 'ahooks';
import { Avatar, Dropdown } from 'antd';
import { useRef } from 'react';
import { Link, history } from 'umi';
import styles from './index.module.less';

function PageFullscreenButton(props: React.HTMLAttributes<HTMLDivElement>) {
  const ref = useRef(document.body);
  const [isFullscreen, { toggleFullscreen }] = useFullscreen(ref);

  return (
    <div {...props} onClick={toggleFullscreen}>
      {!isFullscreen ? <FullscreenOutlined /> : <FullscreenExitOutlined />}
    </div>
  );
}

export default function HeaderBar() {
  const { data: userInfo } = useUserinfoStore();

  return (
    <div className={`${styles.headerContainer} header`}>
      <Link to="/" className={styles.logoTitle}>
        <span className={styles.title}>
          <img src={LOGO} alt="" />
        </span>
      </Link>
      <div className={styles.userContainer}>
        <div className={styles.item}>
          <PageFullscreenButton />
        </div>
        <Dropdown
          menu={{
            items: [
              {
                label: '个人资料',
                key: 'userinfo',
                onClick: () => {
                  history.push('/userinfo');
                },
              },
              {
                label: '退出',
                key: 'outlogin',
                onClick: () => {
                  outLogin();
                  history.push('/login');
                },
              },
            ],
          }}
        >
          <div className={styles.item}>
            <Avatar src="https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png">
              Avatar
            </Avatar>
            <span className={styles.username}>{userInfo?.name || userInfo?.username}</span>
            <DownOutlined style={{ fontSize: 12 }} />
          </div>
        </Dropdown>
      </div>
    </div>
  );
}
