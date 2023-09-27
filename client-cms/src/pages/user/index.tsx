import Header from '@/components/Header';
import PageContainer from '@/components/PageContainer';
import PageTable from '@/components/PageTable';
import { ADMIN_USER_STATUS } from '@/constants';
import { transConstValue, transformPagination, transformSort } from '@/utils';
import { message } from '@/utils/notice';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { Avatar, Input, Popconfirm, Space, Tag } from 'antd';
import { ModelUser } from 'interface/serverApi';
import { useRef, useState } from 'react';
import { getUserList, updateStatus } from './module';

type TableItem = ModelUser;

export default function UserAdminList() {
  const tableRef = useRef<ActionType>();
  const [searchForm, setSearchForm] = useState({
    keyword: '',
  });

  const columns: ProColumns<TableItem>[] = [
    {
      dataIndex: 'username',
      title: '用户名',
      width: 200,
      render: (_, row) => {
        return (
          <Space>
            <Avatar size={44} shape="square" src={row.avatar || row.name}>
              {row.name?.substring(0, 1)}
            </Avatar>
            <div>
              <Space size={1}>
                <Tag color="blue">{row.id}</Tag>
                {row.name}
              </Space>
              <div>
                <small>{row.username}</small>
              </div>
            </div>
          </Space>
        );
      },
    },
    {
      dataIndex: 'status',
      title: '状态',
      width: 100,
      render: (_, row) => {
        const { label, color } =
          Object.values(ADMIN_USER_STATUS).find((v) => v.value === row.status) || {};
        return <Tag color={color}>{label}</Tag>;
      },
    },
    {
      dataIndex: 'email',
      title: '邮箱',
    },
    {
      dataIndex: 'created_at',
      title: '创建时间',
      valueType: 'dateTime',
      sorter: true,
    },
    {
      dataIndex: 'updated_at',
      title: '修改时间',
      valueType: 'dateTime',
      sorter: true,
    },
    {
      dataIndex: 'operate',
      title: '操作',
      hideInSearch: true,
      width: 250,
      render: (_, row) => {
        const operate = transConstValue(ADMIN_USER_STATUS)[row.status === 1 ? -1 : 1];
        return (
          <Space>
            <Popconfirm
              title={
                <div>
                  确定要 <span style={{ color: operate.color }}>{operate.action}</span> 用户{' '}
                  <b>{row.name}</b> 吗 ？
                </div>
              }
              onConfirm={() => {
                updateStatus(row.id!, operate.value).then(() => {
                  message.success(operate.action + '完成');
                  tableRef.current?.reload();
                });
              }}
            >
              <a style={{ color: operate.color }}>{operate.action}</a>
            </Popconfirm>
          </Space>
        );
      },
    },
  ];

  return (
    <>
      <Header />
      <PageContainer>
        <PageTable<TableItem>
          columns={columns}
          request={(params, sort) => {
            return getUserList({
              ...transformPagination(params),
              ...transformSort(sort),
              keyword: searchForm.keyword,
            }).then(({ data }) => {
              return { data: data.data.list, total: data.data.total || 0 };
            });
          }}
          headerTitle={
            <Input.Search
              value={searchForm.keyword}
              onChange={(e) => {
                setSearchForm((state) => ({
                  ...state,
                  keyword: e.target.value.trim(),
                }));
              }}
              style={{ width: 400 }}
              placeholder="请输入姓名/用户名搜索"
              enterButton={<>搜索</>}
              allowClear={true}
              onSearch={() => {
                tableRef.current?.setPageInfo?.({ current: 1 });
                tableRef.current?.reload();
              }}
            />
          }
          actionRef={tableRef}
        />
      </PageContainer>
    </>
  );
}
