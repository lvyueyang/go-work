import Header from '@/components/Header';
import PageContainer from '@/components/PageContainer';
import UploadImage from '@/components/UploadImage';
import { BANNER_POSITION, BANNER_STATUS } from '@/constants';
import { BannerUpdateDto } from '@/interface/serverApi';
import { message } from '@/utils/notice';
import { useRequest } from 'ahooks';
import { Button, Card, Form, Input, Row, Select } from 'antd';
import { useEffect } from 'react';
import { history, useParams } from 'umi';
import { createApi, getDetailApi, updateApi } from './module';

type FormValues = BannerUpdateDto;

export default function BannerForm() {
  const { id } = useParams();
  const isUpdate = !!id;
  const [form] = Form.useForm<FormValues>();
  const { run: submitHandler, loading } = useRequest(
    async () => {
      const values = form.getFieldsValue();
      if (isUpdate) {
        await updateApi(id, {
          ...values,
        });
        message.success('更新成功');
      } else {
        await createApi(values);
        history.push('/banner');
        message.success('创建成功');
      }
    },
    { manual: true },
  );

  useEffect(() => {
    if (id) {
      getDetailApi(Number(id)).then((res) => {
        form.setFieldsValue(res.data.data);
      });
    }
  }, [id]);

  return (
    <>
      <Header title={isUpdate ? '修改广告位' : '新增广告位'} />
      <PageContainer>
        <Card style={{ maxWidth: 1000, margin: '0 auto' }}>
          <Form<FormValues>
            form={form}
            onFinish={submitHandler}
            labelCol={{ xs: 3 }}
            initialValues={{
              position: BANNER_POSITION.HOME.value,
              status: BANNER_STATUS.NORMAL.value,
            }}
          >
            <Form.Item label="广告位标题" name="title" rules={[{ required: true }]}>
              <Input />
            </Form.Item>
            <Form.Item
              label="广告位封面"
              name="cover"
              rules={[{ required: true, message: '请上传广告位封面' }]}
            >
              <UploadImage />
            </Form.Item>
            <Form.Item label="跳转地址" name="link">
              <Input />
            </Form.Item>
            <Form.Item label="广告位描述" name="desc">
              <Input.TextArea />
            </Form.Item>
            <Form.Item label="所处位置" name="position">
              <Select options={Object.values(BANNER_POSITION)} style={{ width: 300 }} />
            </Form.Item>
            <Form.Item label="广告状态" name="status">
              <Select options={Object.values(BANNER_STATUS)} style={{ width: 300 }} />
            </Form.Item>
            <Form.Item label=" ">
              <Row justify="center">
                <Button style={{ width: 160 }} type="primary" htmlType="submit" loading={loading}>
                  提交
                </Button>
              </Row>
            </Form.Item>
          </Form>
        </Card>
      </PageContainer>
    </>
  );
}
