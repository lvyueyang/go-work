import PageTable from '@/components/PageTable';
import { ModalType, useFormModal } from '@/hooks/useFormModal';
import { PERMISSION_CODE, Permission } from '@/permission';
import { transformPagination } from '@/utils';
import { message } from '@/utils/notice';
import { ActionType, ProColumns } from '@ant-design/pro-components';
import { useRequest } from 'ahooks';
import { Button, Form, Input, Modal, Popconfirm, Space, Spin, Transfer } from 'antd';
import { useRef, useState } from 'react';
import { createApi, getCodeListApi, getListApi, removeApi, updateApi, updateCodeApi } from './module';
import { ApiAdminRoleUpdateReq, ModelAdminRole } from '@opd/api-interface';

type TableItem = ModelAdminRole;
type FormValues = ApiAdminRoleUpdateReq;

class CodeModalState {
  role?: number;
  open: boolean = false;
  loading: boolean = false;
  values: string[] = [];
}

export default function AdminRoleList() {
  const tableRef = useRef<ActionType>();
  const [codeModal, setCodeModal] = useState(new CodeModalState());

  const { data: codesList } = useRequest(() => {
    return getCodeListApi().then((res) => {
      const list = res.data.data.map((item) => {
        return {
          ...item,
        };
      });
      /** 根据 code 排序 */
      list.sort((a, b) => a.code!.localeCompare(b.code!));
      return list;
    });
  });

  const infoModal = useFormModal<FormValues & { id?: number }>({
    submit: (values, modal) => {
      if (modal.type === ModalType.UPDATE) {
        return updateApi({
          ...values,
          id: Number(values.id),
        }).then(() => {
          tableRef.current?.reload();
        });
      }
      return createApi(values).then(() => {
        tableRef.current?.reload();
      });
    },
  });

  const submitCodeHandler = (codes: string[]) => {
    if (!codeModal.role) return;
    setCodeModal((state) => ({ ...state, loading: true }));

    updateCodeApi({ codes, id: codeModal.role })
      .then(() => {
        message.success('权限修改成功');
        tableRef.current?.reload();
      })
      .finally(() => {
        setCodeModal((state) => ({ ...state, loading: false }));
      });
  };
  const columns: ProColumns<TableItem>[] = [
    {
      dataIndex: 'id',
      title: 'ID',
      width: 100,
    },
    {
      dataIndex: 'name',
      title: '角色名称',
      width: 140,
    },
    {
      dataIndex: 'code',
      title: '角色编码',
      width: 140,
    },
    {
      dataIndex: 'desc',
      title: '描述',
      width: 200,
      ellipsis: true,
    },
    {
      dataIndex: 'created_at',
      title: '创建时间',
      valueType: 'dateTime',
    },
    {
      dataIndex: 'updated_at',
      title: '修改时间',
      valueType: 'dateTime',
    },
    {
      dataIndex: 'operate',
      title: '操作',
      hideInSearch: true,
      width: 160,
      render: (_, row) => {
        return (
          <Space>
            <Permission code={PERMISSION_CODE.ADMIN_ROLE_UPDATE_CODE}>
              <a
                onClick={() => {
                  setCodeModal((state) => ({
                    ...state,
                    values: row.permission_codes || [],
                    open: true,
                    role: row.id,
                  }));
                }}
              >
                编辑权限
              </a>
            </Permission>
            <Permission code={PERMISSION_CODE.ADMIN_ROLE_UPDATE_INFO}>
              <a
                onClick={() => {
                  infoModal.form.setFieldsValue(row);
                  infoModal.formModalShow(ModalType.UPDATE);
                }}
              >
                编辑
              </a>
            </Permission>
            <Permission code={PERMISSION_CODE.ADMIN_ROLE_DELETE}>
              <Popconfirm
                title={`确定要删除角色 ${row.name} 吗？`}
                onConfirm={() => {
                  const close = message.loading('删除中...', 0);
                  removeApi(row.id)
                    .then(() => {
                      message.success('删除成功');
                      tableRef.current?.reload();
                    })
                    .finally(() => {
                      close();
                    });
                }}
              >
                <a>删除</a>
              </Popconfirm>
            </Permission>
          </Space>
        );
      },
    },
  ];

  return (
    <>
      <PageTable<TableItem>
        columns={columns}
        request={(params) => {
          return getListApi({ ...transformPagination(params) }).then(({ data }) => {
            return { data: data.data.list, total: data.data.total || 0 };
          });
        }}
        actionRef={tableRef}
        toolBarRender={() => [
          <Button
            type="primary"
            key="create"
            onClick={() => {
              infoModal.form.resetFields();
              infoModal.formModalShow();
            }}
          >
            新增角色
          </Button>,
        ]}
      />

      <Modal
        maskClosable={false}
        open={infoModal.formModal.open}
        title={`${infoModal.formModalTitle}`}
        onCancel={infoModal.formModalClose}
        onOk={infoModal.submitHandler}
        okButtonProps={{
          loading: infoModal.submitLoading,
        }}
        width={400}
      >
        <br />
        <Form form={infoModal.form} labelCol={{ flex: '80px' }} initialValues={{ is_stable: false }}>
          {infoModal.formModal.type === ModalType.UPDATE && (
            <Form.Item name="id" hidden>
              <Input />
            </Form.Item>
          )}
          <Form.Item name="name" label="角色名称" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="code" label="角色编码" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="desc" label="描述">
            <Input.TextArea />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        maskClosable={false}
        open={codeModal.open}
        title={`修改权限`}
        onCancel={() => {
          setCodeModal((state) => ({
            ...state,
            open: false,
          }));
        }}
        footer={false}
        okButtonProps={{
          loading: codeModal.loading,
        }}
        width={700}
        style={{ top: 5 }}
      >
        <br />
        <Spin spinning={codeModal.loading}>
          <Transfer
            listStyle={{
              width: 300,
              height: 600,
            }}
            showSearch
            rowKey={(row) => row.code!}
            dataSource={codesList}
            titles={['未拥有权限', '已拥有权限']}
            targetKeys={codeModal.values}
            onChange={(values) => {
              setCodeModal((state) => ({
                ...state,
                values,
              }));
              submitCodeHandler(values);
            }}
            render={(item) => item.name!}
          />
        </Spin>
      </Modal>
    </>
  );
}
