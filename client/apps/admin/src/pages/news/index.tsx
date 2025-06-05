import { transformPagination, transformSort } from '@/utils';
import { message } from '@/utils/notice';
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components';
import { Button, Input, Popconfirm, Space, Tag } from 'antd';
import { useRef, useState } from 'react';
import { Link } from 'umi';
import { createApi, getListApi, removeApi, updateApi } from './module';
import { ApiNewsInfoRes, ApiNewsUpdateReq } from '@opd/api-interface';
import { FormModal } from './modal';
import { ModalType, useFormModal } from '@/hooks/useFormModal';

type TableItem = ApiNewsInfoRes;
type FormValues = ApiNewsUpdateReq;

export default function NewsListPage() {
  const [searchForm, setSearchForm] = useState({
    keyword: '',
  });
  const tableRef = useRef<ActionType>();

  const formModal = useFormModal<FormValues>({
    submit: async (values) => {
      const isUpdate = !!values.id;
      if (isUpdate) {
        await updateApi({
          ...values,
          id: values.id,
        });
      } else {
        await createApi({ ...values });
        message.success('创建成功');
      }
    },
  });

  const columns: ProColumns<TableItem>[] = [
    {
      dataIndex: 'cover',
      title: '缩略图',
      valueType: 'image',
    },
    {
      dataIndex: 'title',
      title: '新闻名称',
      render: (_, row) => {
        return (
          <Space size={1}>
            <Tag>{row.id}</Tag>
            {row.title}
          </Space>
        );
      },
    },
    {
      dataIndex: 'recommend',
      title: '推荐等级',
      sorter: true,
    },
    {
      dataIndex: 'desc',
      title: '新闻描述',
      width: 300,
      ellipsis: true,
    },
    {
      dataIndex: 'push_date',
      title: '发布日期',
      valueType: 'dateTime',
      sorter: true,
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
      width: 160,
      render: (_, row) => {
        return (
          <Space>
            <Link to={`/news/update/${row.id}`}>编辑</Link>
            <Popconfirm
              title="确定要删除这个新闻吗？"
              onConfirm={() => {
                const close = message.loading('删除中...', 0);
                removeApi(row.id!)
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
          </Space>
        );
      },
    },
  ];

  return (
    <>
      <ProTable<TableItem>
        columns={columns}
        rowKey="id"
        bordered
        search={false}
        request={(params, sort) => {
          return getListApi({
            ...transformPagination(params),
            ...transformSort(sort),
            ...searchForm,
          }).then(({ data }) => {
            return { data: data.data.list, total: data.data.total || 0 };
          });
        }}
        actionRef={tableRef}
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
            placeholder="请输入新闻名称或 ID 搜索"
            enterButton={<>搜索</>}
            onSearch={() => {
              tableRef.current?.setPageInfo?.({ current: 1 });
              tableRef.current?.reload();
            }}
          />
        }
        toolBarRender={() => [
          <Button
            key="create"
            type="primary"
            onClick={() => {
              // history.push('/news/create');
              formModal.form.resetFields();
              formModal.formModalShow();
            }}
          >
            新增新闻
          </Button>,
        ]}
      />

      <FormModal formModal={formModal} />
    </>
  );
}
