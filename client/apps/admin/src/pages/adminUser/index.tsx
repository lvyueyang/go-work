import { ConfirmPasswordItem } from '@/components/ConfirmPasswordItem';
import PageTable from '@/components/PageTable';
import { ADMIN_USER_STATUS } from '@/constants';
import { ModalType, useFormModal } from '@/hooks/useFormModal';
import { ApiAdminUserCreateReq, ModelAdminUser } from '@opd/api-interface';
import { customTheme } from '@/theme';
import { transConstValue, transformPagination, transformSort } from '@/utils';
import { message } from '@/utils/notice';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { useRequest } from 'ahooks';
import { Avatar, Button, Form, Input, Modal, Popconfirm, Space, Spin, Tag, Transfer } from 'antd';
import { useRef, useState } from 'react';
import { getListApi as getRoleListApi } from '../adminRole/module/services';
import {
  createUserApi,
  deleteUserApi,
  getUserListApi,
  resetPasswordApi,
  updateRoleApi,
  updateStatusApi,
  updateUserApi,
} from './module';

type TableItem = ModelAdminUser;

class RoleModalState {
  user?: number;
  open: boolean = false;
  loading: boolean = false;
  values: string[] = [];
}

export default function UserAdminList() {
  const tableRef = useRef<ActionType>();
  const [roleModal, setRoleModal] = useState(new RoleModalState());
  const [searchForm, setSearchForm] = useState({
    keyword: '',
  });
  const { data: roleList } = useRequest(() => {
    return getRoleListApi({ current: 1, page_size: 10000 }).then((res) => res.data.data.list);
  });

  const { submitLoading, form, formModal, formModalShow, formModalClose, submitHandler, formModalTitle } = useFormModal<
    ApiAdminUserCreateReq & { id?: number }
  >({
    submit: (values, modal) => {
      if (modal.type === ModalType.UPDATE && values.id) {
        return updateUserApi({
          id: values.id,
          name: values.name,
        }).then(() => {
          tableRef.current?.reload();
        });
      }

      if (modal.type === ModalType.OTHER && values.id) {
        return resetPasswordApi({
          id: values.id,
          password: values.password,
        }).then(() => {
          tableRef.current?.reload();
        });
      }
      return createUserApi({
        username: values.username,
        password: values.password,
        name: values.name,
        email: values.email,
      }).then(() => {
        tableRef.current?.reload();
      });
    },
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
        const { label, color } = Object.values(ADMIN_USER_STATUS).find((v) => v.value === row.status) || {};
        return <Tag color={color}>{label}</Tag>;
      },
    },
    {
      dataIndex: 'role',
      title: '角色',
      ellipsis: true,
      renderText: (_, row) => {
        if (row.is_root) {
          return '超级管理员';
        }
        return row.roles?.map((r) => r.name).join('，');
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
            <a
              onClick={() => {
                form.setFieldsValue(row);
                formModalShow(ModalType.UPDATE);
              }}
            >
              编辑
            </a>
            {!row.is_root && (
              <a
                onClick={() => {
                  setRoleModal((state) => ({
                    ...state,
                    values: row.roles?.map((item) => item.id!.toString()) || [],
                    open: true,
                    user: row.id,
                  }));
                }}
              >
                变更角色
              </a>
            )}
            <a
              onClick={() => {
                form.setFieldsValue(row);
                formModalShow(ModalType.OTHER);
              }}
            >
              重置密码
            </a>
            {!row.is_root && (
              <Popconfirm
                title={
                  <div>
                    确定要 <span style={{ color: operate.color }}>{operate.action}</span> 用户 <b>{row.name}</b> 吗 ？
                  </div>
                }
                onConfirm={() => {
                  updateStatusApi({ id: row.id!, status: operate.value }).then(() => {
                    message.success(operate.action + '完成');
                    tableRef.current?.reload();
                  });
                }}
              >
                <a style={{ color: operate.color }}>{operate.action}</a>
              </Popconfirm>
            )}
            <Popconfirm
              title={
                <div>
                  确定要删除管理员 <b>{row.name}</b> 吗 ？
                </div>
              }
              onConfirm={() => {
                deleteUserApi(row.id!).then(() => {
                  message.success('删除完成');
                  tableRef.current?.reload();
                });
              }}
            >
              <a style={{ color: customTheme.colorError }}>删除</a>
            </Popconfirm>
          </Space>
        );
      },
    },
  ];

  const submitRoleHandler = (roles: number[]) => {
    if (!roleModal.user) return;
    setRoleModal((state) => ({ ...state, loading: true }));

    updateRoleApi({ role_ids: roles, user_id: roleModal.user })
      .then(() => {
        message.success('角色修改成功');
        tableRef.current?.reload();
      })
      .finally(() => {
        setRoleModal((state) => ({ ...state, loading: false }));
      });
  };

  return (
    <>
      <PageTable<TableItem>
        columns={columns}
        request={(params, sort) => {
          return getUserListApi({
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
        toolBarRender={() => [
          <Button
            key="create"
            type="primary"
            onClick={() => {
              form.resetFields();
              formModalShow();
            }}
          >
            新增管理账户
          </Button>,
        ]}
      />
      <Modal
        maskClosable={false}
        open={formModal.open}
        title={formModal.type === ModalType.OTHER ? '重置密码' : `${formModalTitle}用户`}
        onCancel={formModalClose}
        onOk={submitHandler}
        okButtonProps={{
          loading: submitLoading,
        }}
      >
        <br />
        <Form form={form} labelCol={{ span: 4 }} initialValues={{ redundancy_count: 1 }}>
          {formModal.type !== ModalType.CREATE && (
            <Form.Item name="id" hidden>
              <Input />
            </Form.Item>
          )}
          <Form.Item name="username" label="用户名" rules={[{ required: true }]}>
            <Input disabled={formModal.type !== ModalType.CREATE} />
          </Form.Item>
          <Form.Item name="email" label="邮箱" rules={[{ required: true }]}>
            <Input type="email" disabled={formModal.type !== ModalType.CREATE} />
          </Form.Item>
          <Form.Item name="name" label="姓名" rules={[{ required: true }]}>
            <Input disabled={formModal.type === ModalType.OTHER} />
          </Form.Item>
          {(formModal.type === ModalType.CREATE || formModal.type === ModalType.OTHER) && (
            <>
              <Form.Item name="password" label="密码" rules={[{ required: true }]}>
                <Input.Password />
              </Form.Item>
              <ConfirmPasswordItem />
            </>
          )}
        </Form>
      </Modal>
      <Modal
        maskClosable={false}
        open={roleModal.open}
        title={`修改权限`}
        onCancel={() => {
          setRoleModal((state) => ({
            ...state,
            open: false,
          }));
        }}
        footer={false}
        okButtonProps={{
          loading: roleModal.loading,
        }}
        width={700}
      >
        <br />
        <Spin spinning={roleModal.loading}>
          <Transfer
            listStyle={{
              width: 300,
              height: 500,
            }}
            rowKey={(row) => row.id + ''}
            dataSource={roleList}
            titles={['未拥有', '已拥有']}
            targetKeys={roleModal.values}
            onChange={(values) => {
              setRoleModal((state) => ({
                ...state,
                values: values as string[],
              }));
              submitRoleHandler(values.map(Number));
            }}
            render={(item) => item.name}
          />
        </Spin>
      </Modal>
    </>
  );
}
